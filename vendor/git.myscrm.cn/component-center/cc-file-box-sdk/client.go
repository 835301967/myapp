package file_box

import (
	"context"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/module/client_conn"
	"git.myscrm.cn/component-center/cc-file-box-sdk/proto/cc-file-box/personalization_policy_4sdk"
	"git.myscrm.cn/component-center/cc-file-box-sdk/proto/cc-file-box/store_scene_4sdk"
	_ "github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sync"
	"sync/atomic"
	"unsafe"
)

var (
	ccFileBoxConnOnce  sync.Once
	ccFileBoxConn      *grpc.ClientConn
	callInfoPool       sync.Pool
	grpcConnInitStatus int32
)

func init() {
	callInfoPool.New = func() interface{} {
		return &CallInfo{}
	}
}

func newCcFileBoxConn() *grpc.ClientConn {
	// 并发的协程只有一个会强到锁 去初始化grpc client
	onceNewCcFileBoxConn()
	// 如果抢到锁的协程初始化失败， 则强锁失败协程会在这里得到一次补偿，重新获取grpc client
	if atomic.LoadInt32(&grpcConnInitStatus) == 0 {
		atomic.StoreUint32((*uint32)(unsafe.Pointer(&ccFileBoxConnOnce)), 0)
		onceNewCcFileBoxConn()
	}
	return ccFileBoxConn
}

func onceNewCcFileBoxConn() {
	ccFileBoxConnOnce.Do(func() {
		// 初始化不成功则painc
		var (
			conn *client_conn.Conn
			err  error
		)
		// 增加重试， 防止grpc连接抖动造成文件盒子client不可用
		for i := 0; i < 3; i++ {
			conn, err = client_conn.NewConn("cc-file-box")
			if err != nil {
				continue
			}
			ccFileBoxConn, err = conn.GetAPMConn(context.Background())
			if err != nil {
				continue
			}
			break
		}
		if err == nil {
			// 证明grpc client 初始化成功
			atomic.StoreInt32(&grpcConnInitStatus, 1)
		} else {
			panic(err)
		}
	})
}

type Client struct {
	*Config
	hooks []Hook
	*ossManager
	ccFileBoxConn *grpc.ClientConn
}

func NewClient(c Config) *Client {
	if c.Scene == "" {
		panic("scene can not empty")
	}
	if c.ReadWriteTimeout == 0 {
		c.ReadWriteTimeout = defaultReadWriteTimeout
	}
	if c.ConnectOSSTimeout == 0 {
		c.ConnectOSSTimeout = defaultConnectOSSTimeout
	}
	if c.SignExpiredInSec == 0 {
		c.SignExpiredInSec = defaultSignExpiredInSec
	}
	if c.DownloadOnlineExp == 0 {
		c.DownloadOnlineExp = defaultDownloadOnlineExp
	}
	ccFileBoxConn := newCcFileBoxConn()
	// 初始化oss manager失败则panic
	ossManager, err := newOssManager(c, store_scene_4sdk.NewStoreScene4SdkServiceClient(ccFileBoxConn), personalization_policy_4sdk.NewPersonalizationPolicy4SdkServiceClient(ccFileBoxConn))
	if err != nil {
		panic(err)
	}
	return &Client{
		Config:        &c,
		ossManager:    ossManager,
		ccFileBoxConn: ccFileBoxConn,
		hooks:         []Hook{newMonitorHook()},
	}
}

func (c *Client) AddHook(hooks []Hook) {
	c.hooks = append(c.hooks, hooks...)
}

func (c *Client) genDefaultOp(ctx context.Context) *Op {
	var orgCode string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if len(md["orgcode"]) > 0 {
			orgCode = md["orgcode"][0]
		}
	}
	return &Op{
		c:       c.Config,
		orgCode: orgCode,
		scene:   c.Scene,
		appCode: c.defaultStoreScene.AppCode,
	}
}

func (c *Client) applyOp(ctx context.Context, ops ...OpOption) *Op {
	op := c.genDefaultOp(ctx)
	for _, option := range ops {
		option(op)
	}
	return op
}

func (c *Client) before(ctx context.Context, info *CallInfo) context.Context {
	for _, hook := range c.hooks {
		ctx = hook.Before(ctx, info)
	}
	return ctx
}

func (c *Client) after(ctx context.Context, info *CallInfo) {
	for _, hook := range c.hooks {
		hook.After(ctx, info)
	}
	callInfoPool.Put(info)
}

func (c *Client) newCallInfo(methodName MethodName, orgcode string) *CallInfo {
	callInfo := callInfoPool.Get().(*CallInfo)
	storeScene := c.getRuntimeStoreScene(orgcode)
	callInfo.AppCode = storeScene.AppCode
	callInfo.Method = methodName
	callInfo.CloudFactory = storeScene.CloudFactory
	callInfo.Bucket = storeScene.Bucket
	callInfo.Err = nil
	return callInfo
}

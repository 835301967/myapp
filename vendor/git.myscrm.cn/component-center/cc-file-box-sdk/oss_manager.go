package file_box

import (
	"context"
	"errors"
	"git.myscrm.cn/component-center/cc-file-box-sdk/iface"
	"git.myscrm.cn/component-center/cc-file-box-sdk/oss_impl/aliyun"
	"git.myscrm.cn/component-center/cc-file-box-sdk/oss_impl/common"
	"git.myscrm.cn/component-center/cc-file-box-sdk/proto/cc-file-box/personalization_policy_4sdk"
	"git.myscrm.cn/component-center/cc-file-box-sdk/proto/cc-file-box/store_scene_4sdk"
	"sync"
	"time"
)

const (
	ossExpire = 5 // min
)

var (
	cloudFactoryCannotMatch = errors.New("云厂商类型不匹配")
)

func newOSS(storeScene *storeScene, config Config) (oss iface.OSS, err error) {
	switch storeScene.CloudFactory {
	case AliYun:
		return aliyun.NewAliYunOss(common.AliYunSetting{
			CommonSetting: common.CommonSetting{
				AccessKeyId:     storeScene.AccessKeyId,
				AccessKeySecret: storeScene.AccessKeySecret,
				Bucket:          storeScene.Bucket.name,
				Location:        storeScene.Bucket.location,
				IsHttpScheme:    config.IsHttpScheme,
				CustomHost:      storeScene.CustomHost,
				IsInternal:      config.IsInternal,
			},
			ConnectTimeout:   config.ConnectOSSTimeout,
			ReadWriteTimeout: config.ReadWriteTimeout,
		})
	}
	err = cloudFactoryCannotMatch
	return
}

type ossManager struct {
	Config
	mu                                     sync.Mutex
	defaultOss                             iface.OSS
	personalizationPolicyOss               map[string]iface.OSS
	expireTime                             time.Time // 用于定期刷新ram账号
	storeScene4SdkServiceClient            store_scene_4sdk.StoreScene4SdkServiceClient
	personalizationPolicy4SdkServiceClient personalization_policy_4sdk.PersonalizationPolicy4SdkServiceClient
	defaultStoreScene                      *storeScene
	personalizationPolicyStoreScene        map[string]*storeScene
	breaker                                breaker
}

func newOssManager(config Config, storeScene4SdkServiceClient store_scene_4sdk.StoreScene4SdkServiceClient,
	personalizationPolicy4SdkServiceClient personalization_policy_4sdk.PersonalizationPolicy4SdkServiceClient) (*ossManager, error) {
	ossManager := &ossManager{
		Config:                                 config,
		storeScene4SdkServiceClient:            storeScene4SdkServiceClient,
		breaker:                                newOssManagerBreaker(),
		personalizationPolicy4SdkServiceClient: personalizationPolicy4SdkServiceClient,
	}
	return ossManager, ossManager.initOss(context.Background())
}

func (o *ossManager) getOSS(ctx context.Context, orgcode string) (iface.OSS, error) {
	oss := o.getOssFromCache(orgcode)
	if oss != nil {
		return oss, nil
	}
	err := o.breaker.Call(ctx, func(ctx context.Context) error {
		o.mu.Lock()
		defer o.mu.Unlock()
		oss := o.getOssFromCache(orgcode)
		if oss != nil {
			return nil // 说明其他抢到锁的协程已经刷新oss
		}
		return o.initOss(ctx)
	})

	if err != nil {
		if o.defaultOss != nil && o.personalizationPolicyOss != nil {
			// 不返回错误， 保证不影响使用
			return o.defaultOss, nil
		}
		// 说明是第一次初始化oss就失败
		return nil, err
	}
	return o.getOssFromCache(orgcode), nil
}

func (o *ossManager) initOss(ctx context.Context) error {
	var oss iface.OSS
	storeScene, err := o.getDefaultStoreScene(ctx, o.Scene)
	if err != nil {
		return err
	}
	oss, err = newOSS(storeScene, o.Config)
	if err != nil {
		return err
	}
	personalizationPolicyStoreScene, err := o.getPersonalizationPolicyStoreScene(ctx, o.Scene)
	if err != nil {
		return err
	}
	personalizationPolicyOss := make(map[string]iface.OSS)
	for orgcode, storeScene := range personalizationPolicyStoreScene {
		oss, err := newOSS(storeScene, o.Config)
		if err != nil {
			return err
		}
		personalizationPolicyOss[orgcode] = oss
	}
	// 成功则更改Oss相关信息
	o.defaultStoreScene = storeScene
	o.personalizationPolicyStoreScene = personalizationPolicyStoreScene
	o.defaultOss = oss
	o.personalizationPolicyOss = personalizationPolicyOss
	o.expireTime = time.Now().Add(ossExpire * time.Minute)
	return nil
}

func (o *ossManager) getOssFromCache(orgcode string) iface.OSS {
	if time.Now().Before(o.expireTime) {
		if oss, ok := o.personalizationPolicyOss[orgcode]; ok {
			return oss
		}
		return o.defaultOss
	}
	return nil
}

func (o *ossManager) getRuntimeStoreScene(orgcode string) *storeScene {
	if storeScene, ok := o.personalizationPolicyStoreScene[orgcode]; ok {
		return storeScene
	}
	return o.defaultStoreScene
}

func (o *ossManager) getDefaultStoreScene(ctx context.Context, scene string) (*storeScene, error) {
	resp, err := o.storeScene4SdkServiceClient.Get(ctx, &store_scene_4sdk.GetRequest{SceneCode: scene})
	if err != nil {
		return nil, err
	}
	return &storeScene{
		CloudFactory:    CloudFactoryType(resp.CloudFactory),
		AccessKeyId:     resp.AccessKeyId,
		AccessKeySecret: resp.AccessKeySecret,
		Arn:             resp.ARN,
		AppCode:         resp.AppCode,
		Bucket:          bucket{name: resp.BucketName, location: resp.BucketLocation},
		CustomHost:      resp.CustomHost,
		OverwritePolicy: OverwritePolicyType(resp.OverwritePolicy),
	}, nil
}

func (o *ossManager) getPersonalizationPolicyStoreScene(ctx context.Context, scene string) (map[string]*storeScene, error) {
	res := make(map[string]*storeScene)
	resp, err := o.personalizationPolicy4SdkServiceClient.Get(ctx, &personalization_policy_4sdk.GetRequest{SceneCode: scene})
	if err != nil {
		return nil, err
	}
	for orgcode, policy := range resp.Policy {
		res[orgcode] = &storeScene{
			CloudFactory:    CloudFactoryType(policy.CloudFactory),
			AccessKeyId:     policy.AccessKeyId,
			AccessKeySecret: policy.AccessKeySecret,
			AppCode:         policy.AppCode,
			Bucket:          bucket{name: policy.BucketName, location: policy.BucketLocation},
			CustomHost:      policy.CustomHost,
		}
	}
	return res, nil
}

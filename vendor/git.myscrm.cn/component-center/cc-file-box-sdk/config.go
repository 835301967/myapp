package file_box

const (
	defaultReadWriteTimeout  = 60
	defaultConnectOSSTimeout = 10
	defaultSignExpiredInSec  = 60 * 30 // 30 min
	defaultDownloadOnlineExp = 10      // 10 s
)

type Config struct {
	// 默认操作场景
	// required
	Scene string
	// 签名url是否是http， 默认是https
	IsHttpScheme bool
	// oss 读写超时
	ReadWriteTimeout int64
	// 连接oss超时
	ConnectOSSTimeout int64
	// 签名过期时间
	SignExpiredInSec int64
	// 下载远程文件上传文件中的下载超时时间
	DownloadOnlineExp int64
	// 是否走内网， 默认false
	IsInternal bool
}

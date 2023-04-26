package file_box

import (
	"git.myscrm.cn/component-center/cc-file-box-sdk/oss_impl/common"
	"strings"
)

type OpOption func(*Op)

type Op struct {
	c *Config
	// 应用code
	appCode string
	// scene code
	scene string
	// 租户code
	orgCode string
	// 文件有效期 单位s
	expiredInSec int64
	// 下载在线文件超时时间
	downloadOnlineExp int64

	optionKv         common.OptionKv
	uriDecode        bool
	downloadFileName string
}

// WithOrgCode path携带orgcode
func WithOrgCode(orgCode string) OpOption {
	return func(op *Op) {
		op.orgCode = orgCode
	}
}

// WithoutOrgCode path 去掉orgcode
func WithoutOrgCode() OpOption {
	return func(op *Op) {
		op.orgCode = ""
	}
}

// WithExpiredInSec 签名过期时间
func WithExpiredInSec(expiredInSec /*单位s*/ int64) OpOption {
	return func(op *Op) {
		op.expiredInSec = expiredInSec
	}
}

// WithDownloadOnlineExp 下载在线文件超时时间
func WithDownloadOnlineExp(downloadOnlineExp /*单位s*/ int64) OpOption {
	return func(op *Op) {
		op.downloadOnlineExp = downloadOnlineExp
	}
}

// WithUriDecode 下载在线文件超时时间
func WithUriDecode() OpOption {
	return func(op *Op) {
		op.uriDecode = true
	}
}

// WithSetDownloadFileName 指定下载的文件名
func WithSetDownloadFileName(name string) OpOption {
	return func(op *Op) {
		op.downloadFileName = name
		op.optionKv.SetHttpParams(common.ResponseContentDisposition, name)
	}
}

// WithHttpHeader 携带http头
func WithHttpHeader(key string, value interface{}) OpOption {
	return func(op *Op) {
		op.optionKv.SetHeader(key, value)
	}
}

// WithHttpParams 携带http参数
func WithHttpParams(key string, value interface{}) OpOption {
	return func(op *Op) {
		op.optionKv.SetHttpParams(key, value)
	}
}

// withForbidOverwrite 是否禁止覆盖文件
// 暂时不开放 通过文件盒子后台走配置
func withForbidOverwrite(forbidWrite bool) OpOption {
	return func(op *Op) {
		op.optionKv.SetHeader(common.OssForbidWrite, forbidWrite)
	}
}

// WithProcess 图片处理
func WithProcess(value string) OpOption {
	return func(op *Op) {
		op.optionKv.SetHttpParams(common.Process, value)
	}
}

func (o *Op) getExpiredInSec() int64 {
	if o.expiredInSec == 0 {
		return o.c.SignExpiredInSec
	}
	return o.expiredInSec
}

func (o *Op) getFilePathPrefix() string {
	b := strings.Builder{}
	b.WriteString(o.appCode)
	b.WriteString("/")
	b.WriteString(o.scene)
	if o.orgCode != "" {
		b.WriteString("/")
		b.WriteString(o.orgCode)
	}
	return b.String()
}

func (o *Op) getDownloadOnlineExp() int64 {
	if o.downloadOnlineExp == 0 {
		return o.c.DownloadOnlineExp
	}
	return o.downloadOnlineExp
}

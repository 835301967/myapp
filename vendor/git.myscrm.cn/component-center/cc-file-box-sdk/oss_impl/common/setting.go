package common

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type CommonSetting struct {
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
	Location        string
	CustomHost      string
	IsHttpScheme    bool
	IsInternal      bool
}

func (c CommonSetting) GetEndPoint() string {
	scheme := "https"
	if c.IsHttpScheme {
		scheme = "http"
	}
	if c.CustomHost == "" {
		if c.IsInternal {
			return fmt.Sprintf("%s://%s-internal.aliyuncs.com", scheme, c.Location)
		}
		return fmt.Sprintf("%s://%s.aliyuncs.com", scheme, c.Location)
	}
	return fmt.Sprintf("%s://%s", scheme, c.CustomHost)
}

type AliYunSetting struct {
	CommonSetting
	ARN              string
	ConnectTimeout   int64
	ReadWriteTimeout int64
}

func (a AliYunSetting) GenClientOption() []oss.ClientOption {
	options := make([]oss.ClientOption, 0)
	options = append(options, oss.Timeout(a.ConnectTimeout, a.ReadWriteTimeout))
	if a.CustomHost != "" {
		options = append(options, oss.UseCname(true))
	}
	return options
}

package aliyun

import (
	"git.myscrm.cn/component-center/cc-file-box-sdk/oss_impl/common"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	aliyunSts "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type sts struct {
	setting     common.AliYunSetting
	credentials *aliyunSts.Credentials
}

func newSts(s common.AliYunSetting) *sts {
	return &sts{setting: s}
}

// https://help.aliyun.com/document_detail/184381.htm?spm=a2c4g.11186623.2.19.58324bcehLCzK1#concept-1955433
func (s *sts) genStsCredentials() error {
	client, err := aliyunSts.NewClientWithAccessKey("cn-hangzhou", s.setting.AccessKeyId, s.setting.AccessKeySecret)
	if err != nil {
		return err
	}
	//构建请求对象。
	request := aliyunSts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.DurationSeconds = requests.NewInteger(900)
	request.RoleArn = s.setting.ARN
	request.RoleSessionName = s.setting.Bucket
	response, err := client.AssumeRole(request)
	if err != nil {
		return err
	}
	s.credentials = &response.Credentials
	return nil
}

func (s *sts) GetStsOssClient(endpoint string, connectTimeout int64, readWriteTimeout int64) (*oss.Client, error) {
	if s.credentials == nil {
		err := s.genStsCredentials()
		if err != nil {
			return nil, err
		}
	}
	// todo 判断时间 刷新Credentials
	client, err := oss.New(endpoint, s.credentials.AccessKeyId, s.credentials.AccessKeySecret,
		oss.SecurityToken(s.credentials.SecurityToken), oss.Timeout(connectTimeout, readWriteTimeout))
	if err != nil {
		return nil, err
	}
	return client, nil
}

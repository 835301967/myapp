package aliyun

import (
	"context"
	"git.myscrm.cn/component-center/cc-file-box-sdk/iface"
	"git.myscrm.cn/component-center/cc-file-box-sdk/oss_impl/common"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type Oss struct {
	setting common.AliYunSetting
	client  *oss.Client
	bucket  *oss.Bucket
}

func NewAliYunOss(setting common.AliYunSetting) (*Oss, error) {
	// New client
	client, err := oss.New(setting.GetEndPoint(), setting.AccessKeyId, setting.AccessKeySecret, setting.GenClientOption()...)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(setting.Bucket)
	if err != nil {
		return nil, err
	}
	return &Oss{client: client, bucket: bucket, setting: setting}, nil
}

func (o *Oss) Upload(ctx context.Context, filePath string, fileContent io.Reader, opKv common.OptionKv) (string, error) {
	err := o.bucket.PutObject(filePath, fileContent, opKv.ToAliYunOptions()...)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (o *Oss) Delete(ctx context.Context, filePath string, opKv common.OptionKv) error {
	return o.bucket.DeleteObject(filePath, opKv.ToAliYunOptions()...)
}

func (o *Oss) GetFileToLocal(ctx context.Context, filePath, localFile string, opKv common.OptionKv) error {
	return o.bucket.GetObjectToFile(filePath, localFile, opKv.ToAliYunOptions()...)
}

func (o *Oss) MultipartUpload(ctx context.Context, filePath string, opKv common.OptionKv) (string, iface.MultipartUploadEr, error) {
	imur, err := o.bucket.InitiateMultipartUpload(filePath, opKv.ToAliYunOptions()...)
	if err != nil {
		return "", nil, err
	}
	return filePath, &MultipartUpload{imur: imur, bucket: o.bucket}, err
}

func (o *Oss) GetURLWithSign(ctx context.Context, filePath string, expiredInSec /*文件的过期时间*/ int64, opKv common.OptionKv) (string, error) {
	return o.bucket.SignURL(filePath, oss.HTTPGet, expiredInSec, opKv.ToAliYunOptions()...)
}

func (o *Oss) IsObjectExist(ctx context.Context, filePath string, opKv common.OptionKv) (bool, error) {
	return o.bucket.IsObjectExist(filePath)
}

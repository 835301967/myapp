package iface

import (
	"context"
	"git.myscrm.cn/component-center/cc-file-box-sdk/oss_impl/common"
	"io"
)

type OSS interface {
	MultipartUpload(ctx context.Context, filePath string, opKv common.OptionKv) (string, MultipartUploadEr, error)
	Upload(ctx context.Context, filePath string, fileContent io.Reader, opKv common.OptionKv) (string, error)
	Delete(ctx context.Context, filePath string, opKv common.OptionKv) error
	GetFileToLocal(ctx context.Context, filePath, localFile string, opKv common.OptionKv) error
	GetURLWithSign(ctx context.Context, filePath string, expiredInSec /*文件的过期时间 单位s*/ int64, opKv common.OptionKv) (string, error)
	IsObjectExist(ctx context.Context, filePath string, opKv common.OptionKv) (bool, error)
}

// MultipartUploadEr 分片上传接口
type MultipartUploadEr interface {
	// Abort 取消分片上传
	Abort(ctx context.Context) error
	UploadPart(ctx context.Context, reader io.Reader, partSize int64, partNumber int) error
	UploadBySplitFileAndComplete(ctx context.Context, localFile string, chunkNum int) error
	SplitFileByPartNum(localFile string, chunkNum int) ([]FileChunk, error) // Abort 取消分片上传
	// Complete 完成分片上传
	Complete(ctx context.Context) error
}

type FileChunk struct {
	Number int   // Chunk number
	Offset int64 // Chunk offset
	Size   int64 // Chunk size.
}

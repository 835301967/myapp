package aliyun

import (
	"context"
	"git.myscrm.cn/component-center/cc-file-box-sdk/iface"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
)

type MultipartUpload struct {
	imur   oss.InitiateMultipartUploadResult
	parts  []oss.UploadPart
	bucket *oss.Bucket
}

func (m *MultipartUpload) Abort(ctx context.Context) error {
	return m.bucket.AbortMultipartUpload(m.imur)
}

func (m *MultipartUpload) UploadPart(ctx context.Context, reader io.Reader, partSize int64, partNumber int) error {
	part, err := m.bucket.UploadPart(m.imur, reader, partSize, partNumber)
	if err != nil {
		return err
	}
	m.parts = append(m.parts, part)
	return nil
}

func (m *MultipartUpload) UploadBySplitFileAndComplete(ctx context.Context, localFile string, chunkNum int) error {
	chunks, err := m.SplitFileByPartNum(localFile, chunkNum)
	if err != nil {
		return err
	}
	fd, err := os.Open(localFile)
	defer fd.Close()
	for _, chunk := range chunks {
		_, err = fd.Seek(chunk.Offset, os.SEEK_SET)
		if err != nil {
			return err
		}
		part, err := m.bucket.UploadPart(m.imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			return err
		}
		m.parts = append(m.parts, part)
	}
	return m.Complete(ctx)
}

func (m *MultipartUpload) Complete(ctx context.Context) error {
	_, err := m.bucket.CompleteMultipartUpload(m.imur, m.parts)
	return err
}

func (m *MultipartUpload) SplitFileByPartNum(localFile string, chunkNum int) ([]iface.FileChunk, error) {
	chunks, err := oss.SplitFileByPartNum(localFile, chunkNum)
	if err != nil {
		return nil, err
	}
	fileChunks := make([]iface.FileChunk, 0, len(chunks))
	for _, chunk := range chunks {
		fileChunks = append(fileChunks, iface.FileChunk{
			Size:   chunk.Size,
			Offset: chunk.Offset,
			Number: chunk.Number,
		})
	}
	return fileChunks, nil
}

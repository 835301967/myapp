package file_box

import (
	"context"
	"git.myscrm.cn/component-center/cc-file-box-sdk/iface"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
	"time"
)

const (
	upload              MethodName = "Upload"
	UploadOnlineFile    MethodName = "UploadOnlineFile"
	UploadLocalFile     MethodName = "UploadLocalFile"
	GetObjectToFile     MethodName = "GetObjectToFile"
	GetObject           MethodName = "GetObject"
	Delete              MethodName = "Delete"
	InitMultipartUpload MethodName = "InitMultipartUpload"
	IsObjectExist       MethodName = "IsObjectExist"
)

func (c *Client) upload(ctx context.Context, fileName string, reader io.Reader, op *Op) (path string, err error) {
	var (
		oss iface.OSS
	)
	oss, err = c.getOSS(ctx, op.orgCode)
	if err != nil {
		return
	}
	// 上传文件
	path, err = oss.Upload(ctx, c.genFilePath(op.getFilePathPrefix(), fileName), reader, op.optionKv)
	return
}

func (c *Client) Upload(ctx context.Context, fileName string, reader io.Reader, ops ...OpOption) (path string, err error) {
	op := c.applyOp(ctx, append([]OpOption{withForbidOverwrite(c.defaultStoreScene.OverwritePolicy == ForbidOverwrite)}, ops...)...)
	callInfo := c.newCallInfo(upload, op.orgCode)
	ctx = c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()
	path, err = c.upload(ctx, fileName, reader, op)
	return
}

func (c *Client) UploadOnlineFile(ctx context.Context, fileName, onlineFileUrl string, ops ...OpOption) (path string, err error) {
	op := c.applyOp(ctx, append([]OpOption{withForbidOverwrite(c.defaultStoreScene.OverwritePolicy == ForbidOverwrite)}, ops...)...)
	callInfo := c.newCallInfo(UploadOnlineFile, op.orgCode)
	ctx = c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()

	// http 请求获取文件内容
	hc := &http.Client{
		Timeout: time.Duration(op.getDownloadOnlineExp()) * time.Second,
	}
	resp, err := hc.Get(onlineFileUrl)
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	// 上传文件
	return c.upload(ctx, fileName, resp.Body, op)
}

func (c *Client) UploadLocalFile(ctx context.Context, fileName, sourceFilePath string, ops ...OpOption) (path string, err error) {
	op := c.applyOp(ctx, append([]OpOption{withForbidOverwrite(c.defaultStoreScene.OverwritePolicy == ForbidOverwrite)}, ops...)...)
	callInfo := c.newCallInfo(UploadLocalFile, op.orgCode)
	ctx = c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()
	// 读取文件
	fd, err := os.Open(sourceFilePath)
	if err != nil {
		return
	}
	defer func() {
		_ = fd.Close()
	}()
	// 上传文件
	path, err = c.upload(ctx, fileName, fd, op)
	return
}

func (c *Client) GetObjectUrlByUpload(ctx context.Context, uploadFunc func(ctx context.Context) (string, error), ops ...OpOption) (url string, path string, err error) {
	path, err = uploadFunc(ctx)
	if err != nil {
		return
	}
	url, err = c.GetObjectUrl(ctx, path, ops...)
	return
}

func (c *Client) GetObjectToFile(ctx context.Context, path string, localFile string, ops ...OpOption) (err error) {
	op := c.applyOp(ctx, ops...)
	callInfo := c.newCallInfo(GetObjectToFile, op.orgCode)
	ctx = c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()
	oss, err := c.getOSS(ctx, op.orgCode)
	if err != nil {
		return err
	}
	err = oss.GetFileToLocal(ctx, path, localFile, op.optionKv)
	return
}

func (c *Client) IsObjectExist(ctx context.Context, path string, ops ...OpOption) (ok bool, err error) {
	op := c.applyOp(ctx, ops...)
	callInfo := c.newCallInfo(IsObjectExist, op.orgCode)
	ctx = c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()
	oss, err := c.getOSS(ctx, op.orgCode)
	if err != nil {
		return false, err
	}
	return oss.IsObjectExist(ctx, path, op.optionKv)
}

func (c *Client) GetObjectUrl(ctx context.Context, path string, ops ...OpOption) (url string, err error) {
	op := c.applyOp(ctx, ops...)
	callInfo := c.newCallInfo(GetObject, op.orgCode)
	c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()
	oss, err := c.getOSS(ctx, op.orgCode)
	if err != nil {
		return
	}
	url, err = oss.GetURLWithSign(ctx, path, op.getExpiredInSec(), op.optionKv)
	if err != nil {
		return
	}
	if op.uriDecode {
		urlArr := strings.Split(url, "?")
		strBuilder := strings.Builder{}
		uriDecodeStr, err := url2.QueryUnescape(urlArr[0])
		if err != nil {
			return "", nil
		}
		strBuilder.WriteString(uriDecodeStr)
		strBuilder.WriteString("?")
		for i := 1; i < len(urlArr); i++ {
			strBuilder.WriteString(urlArr[i])
		}
		return strBuilder.String(), nil
	}
	return url, nil
}

func (c *Client) Delete(ctx context.Context, path string, ops ...OpOption) (err error) {
	op := c.applyOp(ctx, ops...)
	callInfo := c.newCallInfo(Delete, op.orgCode)
	ctx = c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()
	oss, err := c.getOSS(ctx, op.orgCode)
	if err != nil {
		return
	}
	err = oss.Delete(ctx, path, op.optionKv)
	return
}

func (c *Client) InitMultipartUpload(ctx context.Context, fileName string, ops ...OpOption) (path string, mu iface.MultipartUploadEr, err error) {
	op := c.applyOp(ctx, ops...)
	callInfo := c.newCallInfo(InitMultipartUpload, op.orgCode)
	ctx = c.before(ctx, callInfo)
	defer func() {
		callInfo.Err = err
		c.after(ctx, callInfo)
	}()
	oss, err := c.getOSS(ctx, op.orgCode)
	return oss.MultipartUpload(ctx, c.genFilePath(op.getFilePathPrefix(), fileName), op.optionKv)
}

func (c *Client) genFilePath(prefix string, fileName string) string {
	b := strings.Builder{}
	b.WriteString(prefix)
	b.WriteString("/")
	b.WriteString(fileName)
	return b.String()
}

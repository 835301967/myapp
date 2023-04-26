# 文件盒子sdk【组件中心】

## 简介

当前各业务部门均依赖于阿里OSS服务存储非结构化数据，数据量庞大且缺乏规范化的管理。由于历史原因，存在多个应用共用同一个存储桶、不同租户的数据难以区分等问题。考虑到**多应用接入**、**数据安全性**和数据私有化等场景。我们需要提供一个产品，提升各部门接入和使用第三方对象存储服务的效率，通过合理的使用方法，来更好的满足各类业务场景需求。

**使用sdk之前，需要在产品运营平台文件盒子配置好存储策略。使用存储==策略编码==即可使用。**

## 应用场景

### 文件读写

* 上传文件
* 下载文件
* 获取文件地址（带签名）
* 删除文件

## 快速上手

### 创建client

```go
// 通过存储策略code创建client
client = NewClient(Config{
   Scene: "xxx",
})
```

### 文件读写

* 上传文件

  ```go
  upLoadPath, err = client.Upload(context.Background(), "test.txt", strings.NewReader("Upload"))
  ```

  > path 的构成是 {appCode}/{sceneCode}/{orgcode}/filename 
  >
  > 如果orgcode为空， 则为{appCode}/{sceneCode}/filename 

* 上传在线文件

  ```go
  onlinePath, err = client.UploadOnlineFile(context.Background(), "test.txt", "http://xxx/text.txt")
  ```

* 上传本地文件

  ```go
  localPath, err = client.UploadLocalFile(context.Background(), "test.txt", "./client.go")
  ```

* 获取文件到本地

  ```go
  err := client.GetObjectToFile(context.Background(), path/*上传文件返回的path*/, "./test.txt")
  ```

* 获取文件地址（带签名）

  ```go
  url, err := client.GetObjectUrl(context.Background(), path/*上传文件返回的path*/)
  ```

  >Url的构成是 http://{bucketName}.oss-cn-shenzhen.aliyuncs.com/{path}?sign=xxx

* 删除文件

  ```go
  err := client.Delete(context.Background(), path/*上传文件返回的path*/)
  ```

## 额外功能

### 不定项参数

* path携带orgcode

  ```go
  upLoadPath, err = client.Upload(context.Background(), "test.txt", strings.NewReader("Upload"), file_box.WithOrgCode("fangzhi_test"))
  ```

  **默认情况，如果ctx，携带orgcode， 则会自动获取orgcode**，如果`ctx`没有`orgcode`， 如果需要`path`需要携带`orgcode`， 可以使用``WithOrgCode("xxx")``携带。

* path 去掉orgcode

  ```go
  upLoadPath, err = client.Upload(context.Background(), "test.txt", strings.NewReader("Upload"), file_box.WithOutOrgCode())
  ```

  如果`ctx`，携带`orgcode`，但不想`path`有`orgcode`属性，则可以使用``WithOutOrgCode()``去掉。

* 设置签名过期时间

  ```go
  // 指定过期时间1小时
  url, err = client.GetObjectUrl(ctx, path, file_box.WithExpiredInSec(60*60))
  ```

  创建client的时候，可以设置默认的签名过期时间，在获取文件地址的时候， 也可指定过期时间，单位为秒。

* 下载在线文件超时时间

  ```go
  // 指定远程文件下载超时时间
  client.UploadOnlineFile(ctx, "test.txt", "http://xxx.txt", file_box.WithDownloadOnlineExp(60))
  ```

  上传在线文件时，需要先下载在线文件，可以通过``WithDownloadOnlineExp(xx/*单位s*/)``设置下载超时时间。（创建client时，可以设置默认下载超时时间）。

* 是否禁止覆盖文件

  ```go
  // 遇到重名文件上传失败
  path, err = client.Upload(ctx, "test.txt", strings.NewReader("test"), file_box.WithForbidOverWrite(true))
  ```

  正常情况，上传重名文件不会报错，可以指定`file_box.WithForbidOverWrite(true)`，这样遇到重名文件则会上传失败。

* 携带http header头或携带http参数

  阿里云等文件存储的参数都是通过http header头或者http params参数设置的， 比如上面的`WithForbidOverWrite`， 如果有什么特殊需求，可以通过`WithHttpHeader`和`WithHttpParams`进行设置。

### 钩子

client可以增加自定义钩子，在调用具体功能前后做一些事情。业务组自定义参数可以通过ctx进行传递。

```go
type Hook interface {
	Before(context.Context, *CallInfo) context.Context
	After(context.Context, *CallInfo)
}

type CallInfo struct {
	Method       MethodName
	CloudFactory CloudFactoryType
	Bucket       bucket
	AppCode      string
	Err          error
}
```

client增加hook

```go
type MyHook struct{}

func (m *MyHook) Before(ctx context.Context, call *CallInfo) context.Context {
	return ctx
}

func (m *MyHook) After(ctx context.Context, call *CallInfo) {
}

client.AddHook(&MyHook{})
```

`sdk`默认添加了监控的`hook`，方便业务组和组件中心对sdk进行监控，方便监控告警和排除问题

```go
		prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "file_box_sdk_request_time",
			Help: "file_box_sdk_request_time",
		}, []string{"cloud_factory", "method", "bucket"})
	 prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "file_box_sdk_request_total",
			Help: "file_box_sdk_request_total",
		}, []string{"cloud_factory", "method", "bucket"})
		 prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "file_box_sdk_request_err_total",
			Help: "file_box_sdk_request_err_total",
		}, []string{"cloud_factory", "method", "bucket"})
```

业务组可以通过`matrics`进行监控告警
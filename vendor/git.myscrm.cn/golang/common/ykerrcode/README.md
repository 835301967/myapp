## ykerrcode模块
#### 说明
目前ykerrcode模块封装了通用错误码定义,错误消息相关接口,gRPC相关接口

#### 通用错误码定义说明
| 常量 | 错误码 | 说明 |
| - | :-: | - | 
| FAIL | 10000000 | 未知错误 |
| INVALID_PARAMS | 10000001 | 非法请求参数 |
| UNAUTH | 10000002 | 无访问权限 |
| NOT_FOUND_RESOURCE | 10000003 | 找不到资源 |
| TIMEOUT | 10000004| 请求超时 |
| DB_ERR | 10000005 | 数据库出错 |
| CACHE_ERR | 10000006 | 缓存出错 |
| CREATE_FILE_FAIL | 10000007 | 创建文件失败 |
| SIGN_ERROR | 10000008 | 签名验证失败 |

#### 代码示例-项目自定义错误
建议在工程下建立一个errcode的文件夹，文件名用模块缩写.go命名  
例如：bitbucket.org/teroy/go/errcode/ydxs.go
```
// 错误码-移动销售模块代号001
package errcode

import (
    "git.myscrm.cn/golang/common/ykerrcode"
)

const (
    YDXS_NOT_FOUND_OPP = 20010001
    YDXS_NOT_FOUND_CST = 20010002
)

func init() {
	dict := map[int]string{
		YDXS_NOT_FOUND_OPP: "找不到机会",
		YDXS_NOT_FOUND_CST: "找不到客户",
	}
	ykerrcode.RegisterErrMsgDict(dict)
}
```

#### 代码示例-项目获取错误信息
```
// 获取通用模块错误信息
errMsg := ykerrcode.GetErrMsg(ykerrcode.DB_ERR)

// 获取自定义模块错误信息
errMsg := ykerrcode.GetErrMsg(errcode.YDXS_NOT_FOUND_OPP)
```

#### 代码示例-gRPC相关接口
```
// 根据错误code获取gRPC自定义返回error
err := ykerrcode.TogRPCError(ykerrcode.DB_ERR)

// 根据错误code获取HTTP状态码
statusCode := ykerrcode.ToHttpStatusCode(ykerrcode.DB_ERR)
```
// 通用错误码组件-通用模块代号000
package ykerrcode

const (
	SUCCESS                = 0
	FAIL                   = 10000000
	INVALID_PARAMS         = 10000001
	UNAUTH                 = 10000002
	NOT_FOUND              = 10000003
	DB_ERR                 = 10000004
	CACHE_ERR              = 10000005
	CREATE_FILE_FAIL       = 10000006
	SIGN_ERROR             = 10000007
	GRPC_SYS_ERR           = 10000008
	CONFIG_ERR             = 10000009
	UNKNOWN                = 10000010
	DEADLINE_EXCEEDED      = 10000011
	ACCESS_DENIED          = 10000012
	LIMIT_EXCEED           = 10000013
	METHOD_NOT_ALLOWED     = 10000014
	GET_TENANT_DB_ERR      = 10000015
	SERVICE_UNAVAILABLE    = 10000016
	PARAMS_TELEPHONE_ERROR = 20000000
)

func init() {
	dict := map[interface{}]string{
		SUCCESS:                "成功",
		FAIL:                   "内部错误",
		INVALID_PARAMS:         "非法请求参数",
		UNAUTH:                 "无访问权限",
		NOT_FOUND:              "找不到资源",
		DB_ERR:                 "数据库出错",
		CACHE_ERR:              "缓存出错",
		CREATE_FILE_FAIL:       "创建文件失败",
		SIGN_ERROR:             "签名验证失败",
		GRPC_SYS_ERR:           "系统错误",
		CONFIG_ERR:             "配置错误",
		UNKNOWN:                "未知错误",
		DEADLINE_EXCEEDED:      "操作超时",
		ACCESS_DENIED:          "拒绝访问",
		LIMIT_EXCEED:           "请求过多，请稍后重试",
		METHOD_NOT_ALLOWED:     "方法不被允许",
		GET_TENANT_DB_ERR:      "获取租户数据库出错",
		SERVICE_UNAVAILABLE:    "服务暂不可用，请稍后重试",
		PARAMS_TELEPHONE_ERROR: "手机号码格式有误",
	}
	RegisterErrMsg(dict)
}

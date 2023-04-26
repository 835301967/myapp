package common

import aliyunOss "github.com/aliyun/aliyun-oss-go-sdk/oss"

const (
	OssForbidWrite             = "X-Oss-Forbid-Overwrite"
	Process                    = "x-oss-process"
	ResponseContentDisposition = "response-content-disposition"
)

type OptionKv struct {
	httpHeader map[string]interface{}
	httpParams map[string]interface{}
}

func (o *OptionKv) ToAliYunOptions() []aliyunOss.Option {
	res := make([]aliyunOss.Option, 0)
	if forbidWrite, ok := o.httpHeader[OssForbidWrite]; ok {
		res = append(res, aliyunOss.ForbidOverWrite(forbidWrite.(bool)))
		delete(o.httpHeader, OssForbidWrite)
	}
	if value, ok := o.httpParams[Process]; ok {
		res = append(res, aliyunOss.Process(value.(string)))
		delete(o.httpParams, Process)
	}
	if value, ok := o.httpParams[ResponseContentDisposition]; ok {
		res = append(res, aliyunOss.ResponseContentDisposition("attachment;filename="+value.(string)))
		delete(o.httpParams, ResponseContentDisposition)
	}
	for k, v := range o.httpHeader {
		res = append(res, aliyunOss.SetHeader(k, v))
	}

	for k, v := range o.httpParams {
		res = append(res, aliyunOss.AddParam(k, v))
	}
	return res
}

func (o *OptionKv) SetHeader(k string, v interface{}) {
	if o.httpHeader == nil {
		o.httpHeader = make(map[string]interface{})
	}
	o.httpHeader[k] = v
}

func (o *OptionKv) SetHttpParams(k string, v interface{}) {
	if o.httpParams == nil {
		o.httpParams = make(map[string]interface{})
	}
	o.httpParams[k] = v
}

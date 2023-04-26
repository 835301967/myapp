package file_box

type CloudFactoryType int64
type OverwritePolicyType int

const (
	AliYun CloudFactoryType = 1

	AllowOverwrite OverwritePolicyType = 1

	ForbidOverwrite OverwritePolicyType = 2
)

type bucket struct {
	name     string
	location string
}

type storeScene struct {
	CloudFactory    CloudFactoryType
	AccessKeyId     string
	AccessKeySecret string
	Arn             string
	AppCode         string
	Bucket          bucket
	CustomHost      string
	OverwritePolicy OverwritePolicyType
}

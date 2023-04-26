// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.5.1
// source: cc-file-box/store_scene_4sdk/store_scene.proto

package store_scene_4sdk

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OverwritePolicy int32

const (
	// 允许
	OverwritePolicy_Allow OverwritePolicy = 0
	// 禁止
	OverwritePolicy_Forbid OverwritePolicy = 1
)

// Enum value maps for OverwritePolicy.
var (
	OverwritePolicy_name = map[int32]string{
		0: "Allow",
		1: "Forbid",
	}
	OverwritePolicy_value = map[string]int32{
		"Allow":  0,
		"Forbid": 1,
	}
)

func (x OverwritePolicy) Enum() *OverwritePolicy {
	p := new(OverwritePolicy)
	*p = x
	return p
}

func (x OverwritePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OverwritePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_cc_file_box_store_scene_4sdk_store_scene_proto_enumTypes[0].Descriptor()
}

func (OverwritePolicy) Type() protoreflect.EnumType {
	return &file_cc_file_box_store_scene_4sdk_store_scene_proto_enumTypes[0]
}

func (x OverwritePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OverwritePolicy.Descriptor instead.
func (OverwritePolicy) EnumDescriptor() ([]byte, []int) {
	return file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescGZIP(), []int{0}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @required
	// 场景code
	// @inject_tag: validate:"required" comment:"场景code"
	SceneCode string `protobuf:"bytes,1,opt,name=scene_code,json=sceneCode,proto3" json:"scene_code,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetSceneCode() string {
	if x != nil {
		return x.SceneCode
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudFactory    int64           `protobuf:"varint,1,opt,name=cloud_factory,json=cloudFactory,proto3" json:"cloud_factory,omitempty"`
	AccessKeyId     string          `protobuf:"bytes,2,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	AccessKeySecret string          `protobuf:"bytes,3,opt,name=access_key_secret,json=accessKeySecret,proto3" json:"access_key_secret,omitempty"`
	ARN             string          `protobuf:"bytes,4,opt,name=ARN,proto3" json:"ARN,omitempty"`
	AppCode         string          `protobuf:"bytes,5,opt,name=appCode,proto3" json:"appCode,omitempty"`
	BucketName      string          `protobuf:"bytes,6,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	BucketLocation  string          `protobuf:"bytes,7,opt,name=bucket_location,json=bucketLocation,proto3" json:"bucket_location,omitempty"`
	CustomHost      string          `protobuf:"bytes,8,opt,name=custom_host,json=customHost,proto3" json:"custom_host,omitempty"`
	OverwritePolicy OverwritePolicy `protobuf:"varint,9,opt,name=overwrite_policy,json=overwritePolicy,proto3,enum=store_scene_4sdk.OverwritePolicy" json:"overwrite_policy,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetCloudFactory() int64 {
	if x != nil {
		return x.CloudFactory
	}
	return 0
}

func (x *GetResponse) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *GetResponse) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *GetResponse) GetARN() string {
	if x != nil {
		return x.ARN
	}
	return ""
}

func (x *GetResponse) GetAppCode() string {
	if x != nil {
		return x.AppCode
	}
	return ""
}

func (x *GetResponse) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *GetResponse) GetBucketLocation() string {
	if x != nil {
		return x.BucketLocation
	}
	return ""
}

func (x *GetResponse) GetCustomHost() string {
	if x != nil {
		return x.CustomHost
	}
	return ""
}

func (x *GetResponse) GetOverwritePolicy() OverwritePolicy {
	if x != nil {
		return x.OverwritePolicy
	}
	return OverwritePolicy_Allow
}

var File_cc_file_box_store_scene_4sdk_store_scene_proto protoreflect.FileDescriptor

var file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x63, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x62, 0x6f, 0x78, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x34, 0x73, 0x64, 0x6b, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x34, 0x73,
	0x64, 0x6b, 0x22, 0x2b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0xe7, 0x02, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x52, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x41, 0x52, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x10, 0x6f,
	0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x34, 0x73, 0x64, 0x6b, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2a, 0x28, 0x0a, 0x0f, 0x4f, 0x76, 0x65,
	0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x62, 0x69,
	0x64, 0x10, 0x01, 0x32, 0x5d, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x34, 0x53, 0x64, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x5f, 0x34, 0x73, 0x64, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f,
	0x34, 0x73, 0x64, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescOnce sync.Once
	file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescData = file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDesc
)

func file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescGZIP() []byte {
	file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescOnce.Do(func() {
		file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescData = protoimpl.X.CompressGZIP(file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescData)
	})
	return file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDescData
}

var file_cc_file_box_store_scene_4sdk_store_scene_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cc_file_box_store_scene_4sdk_store_scene_proto_goTypes = []interface{}{
	(OverwritePolicy)(0), // 0: store_scene_4sdk.OverwritePolicy
	(*GetRequest)(nil),   // 1: store_scene_4sdk.GetRequest
	(*GetResponse)(nil),  // 2: store_scene_4sdk.GetResponse
}
var file_cc_file_box_store_scene_4sdk_store_scene_proto_depIdxs = []int32{
	0, // 0: store_scene_4sdk.GetResponse.overwrite_policy:type_name -> store_scene_4sdk.OverwritePolicy
	1, // 1: store_scene_4sdk.StoreScene4SdkService.Get:input_type -> store_scene_4sdk.GetRequest
	2, // 2: store_scene_4sdk.StoreScene4SdkService.Get:output_type -> store_scene_4sdk.GetResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cc_file_box_store_scene_4sdk_store_scene_proto_init() }
func file_cc_file_box_store_scene_4sdk_store_scene_proto_init() {
	if File_cc_file_box_store_scene_4sdk_store_scene_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cc_file_box_store_scene_4sdk_store_scene_proto_goTypes,
		DependencyIndexes: file_cc_file_box_store_scene_4sdk_store_scene_proto_depIdxs,
		EnumInfos:         file_cc_file_box_store_scene_4sdk_store_scene_proto_enumTypes,
		MessageInfos:      file_cc_file_box_store_scene_4sdk_store_scene_proto_msgTypes,
	}.Build()
	File_cc_file_box_store_scene_4sdk_store_scene_proto = out.File
	file_cc_file_box_store_scene_4sdk_store_scene_proto_rawDesc = nil
	file_cc_file_box_store_scene_4sdk_store_scene_proto_goTypes = nil
	file_cc_file_box_store_scene_4sdk_store_scene_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreScene4SdkServiceClient is the client API for StoreScene4SdkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreScene4SdkServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type storeScene4SdkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreScene4SdkServiceClient(cc grpc.ClientConnInterface) StoreScene4SdkServiceClient {
	return &storeScene4SdkServiceClient{cc}
}

func (c *storeScene4SdkServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/store_scene_4sdk.StoreScene4SdkService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreScene4SdkServiceServer is the server API for StoreScene4SdkService service.
type StoreScene4SdkServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedStoreScene4SdkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreScene4SdkServiceServer struct {
}

func (*UnimplementedStoreScene4SdkServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterStoreScene4SdkServiceServer(s *grpc.Server, srv StoreScene4SdkServiceServer) {
	s.RegisterService(&_StoreScene4SdkService_serviceDesc, srv)
}

func _StoreScene4SdkService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreScene4SdkServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store_scene_4sdk.StoreScene4SdkService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreScene4SdkServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreScene4SdkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "store_scene_4sdk.StoreScene4SdkService",
	HandlerType: (*StoreScene4SdkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _StoreScene4SdkService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cc-file-box/store_scene_4sdk/store_scene.proto",
}

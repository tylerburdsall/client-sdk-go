// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: protos/cacheclient.proto

package client_sdk_go

import (
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

type ECacheResult int32

const (
	ECacheResult_Invalid ECacheResult = 0
	ECacheResult_Ok      ECacheResult = 1
	ECacheResult_Hit     ECacheResult = 2
	ECacheResult_Miss    ECacheResult = 3
)

// Enum value maps for ECacheResult.
var (
	ECacheResult_name = map[int32]string{
		0: "Invalid",
		1: "Ok",
		2: "Hit",
		3: "Miss",
	}
	ECacheResult_value = map[string]int32{
		"Invalid": 0,
		"Ok":      1,
		"Hit":     2,
		"Miss":    3,
	}
)

func (x ECacheResult) Enum() *ECacheResult {
	p := new(ECacheResult)
	*p = x
	return p
}

func (x ECacheResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECacheResult) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_cacheclient_proto_enumTypes[0].Descriptor()
}

func (ECacheResult) Type() protoreflect.EnumType {
	return &file_protos_cacheclient_proto_enumTypes[0]
}

func (x ECacheResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ECacheResult.Descriptor instead.
func (ECacheResult) EnumDescriptor() ([]byte, []int) {
	return file_protos_cacheclient_proto_rawDescGZIP(), []int{0}
}

type XGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CacheKey []byte `protobuf:"bytes,1,opt,name=cache_key,json=cacheKey,proto3" json:"cache_key,omitempty"`
}

func (x *XGetRequest) Reset() {
	*x = XGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cacheclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XGetRequest) ProtoMessage() {}

func (x *XGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cacheclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XGetRequest.ProtoReflect.Descriptor instead.
func (*XGetRequest) Descriptor() ([]byte, []int) {
	return file_protos_cacheclient_proto_rawDescGZIP(), []int{0}
}

func (x *XGetRequest) GetCacheKey() []byte {
	if x != nil {
		return x.CacheKey
	}
	return nil
}

type XGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    ECacheResult `protobuf:"varint,1,opt,name=result,proto3,enum=cache_client.ECacheResult" json:"result,omitempty"`
	CacheBody []byte       `protobuf:"bytes,2,opt,name=cache_body,json=cacheBody,proto3" json:"cache_body,omitempty"`
	Message   string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *XGetResponse) Reset() {
	*x = XGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cacheclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XGetResponse) ProtoMessage() {}

func (x *XGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cacheclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XGetResponse.ProtoReflect.Descriptor instead.
func (*XGetResponse) Descriptor() ([]byte, []int) {
	return file_protos_cacheclient_proto_rawDescGZIP(), []int{1}
}

func (x *XGetResponse) GetResult() ECacheResult {
	if x != nil {
		return x.Result
	}
	return ECacheResult_Invalid
}

func (x *XGetResponse) GetCacheBody() []byte {
	if x != nil {
		return x.CacheBody
	}
	return nil
}

func (x *XGetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type XSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CacheKey        []byte `protobuf:"bytes,1,opt,name=cache_key,json=cacheKey,proto3" json:"cache_key,omitempty"`
	CacheBody       []byte `protobuf:"bytes,2,opt,name=cache_body,json=cacheBody,proto3" json:"cache_body,omitempty"`
	TtlMilliseconds uint32 `protobuf:"varint,3,opt,name=ttl_milliseconds,json=ttlMilliseconds,proto3" json:"ttl_milliseconds,omitempty"`
}

func (x *XSetRequest) Reset() {
	*x = XSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cacheclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XSetRequest) ProtoMessage() {}

func (x *XSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cacheclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XSetRequest.ProtoReflect.Descriptor instead.
func (*XSetRequest) Descriptor() ([]byte, []int) {
	return file_protos_cacheclient_proto_rawDescGZIP(), []int{2}
}

func (x *XSetRequest) GetCacheKey() []byte {
	if x != nil {
		return x.CacheKey
	}
	return nil
}

func (x *XSetRequest) GetCacheBody() []byte {
	if x != nil {
		return x.CacheBody
	}
	return nil
}

func (x *XSetRequest) GetTtlMilliseconds() uint32 {
	if x != nil {
		return x.TtlMilliseconds
	}
	return 0
}

type XSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  ECacheResult `protobuf:"varint,1,opt,name=result,proto3,enum=cache_client.ECacheResult" json:"result,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *XSetResponse) Reset() {
	*x = XSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cacheclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XSetResponse) ProtoMessage() {}

func (x *XSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cacheclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XSetResponse.ProtoReflect.Descriptor instead.
func (*XSetResponse) Descriptor() ([]byte, []int) {
	return file_protos_cacheclient_proto_rawDescGZIP(), []int{3}
}

func (x *XSetResponse) GetResult() ECacheResult {
	if x != nil {
		return x.Result
	}
	return ECacheResult_Invalid
}

func (x *XSetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_cacheclient_proto protoreflect.FileDescriptor

var file_protos_cacheclient_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x0b, 0x5f, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x4b, 0x65, 0x79, 0x22, 0x7b, 0x0a, 0x0c, 0x5f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x74, 0x0a, 0x0b, 0x5f, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x29, 0x0a, 0x10,
	0x74, 0x74, 0x6c, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x74, 0x6c, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x0c, 0x5f, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x3c, 0x0a, 0x0c, 0x45, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x69,
	0x74, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x69, 0x73, 0x73, 0x10, 0x03, 0x22, 0x04, 0x08,
	0x04, 0x10, 0x06, 0x32, 0x85, 0x01, 0x0a, 0x03, 0x53, 0x63, 0x73, 0x12, 0x3e, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x5f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x5f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x03, 0x53,
	0x65, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x5f, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x5f, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x0a, 0x11, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x68, 0x71, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d,
	0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x64,
	0x6b, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_cacheclient_proto_rawDescOnce sync.Once
	file_protos_cacheclient_proto_rawDescData = file_protos_cacheclient_proto_rawDesc
)

func file_protos_cacheclient_proto_rawDescGZIP() []byte {
	file_protos_cacheclient_proto_rawDescOnce.Do(func() {
		file_protos_cacheclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_cacheclient_proto_rawDescData)
	})
	return file_protos_cacheclient_proto_rawDescData
}

var file_protos_cacheclient_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_cacheclient_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_cacheclient_proto_goTypes = []interface{}{
	(ECacheResult)(0),    // 0: cache_client.ECacheResult
	(*XGetRequest)(nil),  // 1: cache_client._GetRequest
	(*XGetResponse)(nil), // 2: cache_client._GetResponse
	(*XSetRequest)(nil),  // 3: cache_client._SetRequest
	(*XSetResponse)(nil), // 4: cache_client._SetResponse
}
var file_protos_cacheclient_proto_depIdxs = []int32{
	0, // 0: cache_client._GetResponse.result:type_name -> cache_client.ECacheResult
	0, // 1: cache_client._SetResponse.result:type_name -> cache_client.ECacheResult
	1, // 2: cache_client.Scs.Get:input_type -> cache_client._GetRequest
	3, // 3: cache_client.Scs.Set:input_type -> cache_client._SetRequest
	2, // 4: cache_client.Scs.Get:output_type -> cache_client._GetResponse
	4, // 5: cache_client.Scs.Set:output_type -> cache_client._SetResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_cacheclient_proto_init() }
func file_protos_cacheclient_proto_init() {
	if File_protos_cacheclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_cacheclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XGetRequest); i {
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
		file_protos_cacheclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XGetResponse); i {
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
		file_protos_cacheclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XSetRequest); i {
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
		file_protos_cacheclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XSetResponse); i {
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
			RawDescriptor: file_protos_cacheclient_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_cacheclient_proto_goTypes,
		DependencyIndexes: file_protos_cacheclient_proto_depIdxs,
		EnumInfos:         file_protos_cacheclient_proto_enumTypes,
		MessageInfos:      file_protos_cacheclient_proto_msgTypes,
	}.Build()
	File_protos_cacheclient_proto = out.File
	file_protos_cacheclient_proto_rawDesc = nil
	file_protos_cacheclient_proto_goTypes = nil
	file_protos_cacheclient_proto_depIdxs = nil
}

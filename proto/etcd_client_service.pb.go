// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/etcd_client_service.proto

package proto

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

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_etcd_client_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_proto_etcd_client_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_proto_etcd_client_service_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_etcd_client_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_etcd_client_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_proto_etcd_client_service_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type EtcdClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Union:
	//
	//	*EtcdClientRequest_Key
	//	*EtcdClientRequest_KeyValue
	Union isEtcdClientRequest_Union `protobuf_oneof:"union"`
}

func (x *EtcdClientRequest) Reset() {
	*x = EtcdClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_etcd_client_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdClientRequest) ProtoMessage() {}

func (x *EtcdClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_etcd_client_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdClientRequest.ProtoReflect.Descriptor instead.
func (*EtcdClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_etcd_client_service_proto_rawDescGZIP(), []int{2}
}

func (m *EtcdClientRequest) GetUnion() isEtcdClientRequest_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (x *EtcdClientRequest) GetKey() *Key {
	if x, ok := x.GetUnion().(*EtcdClientRequest_Key); ok {
		return x.Key
	}
	return nil
}

func (x *EtcdClientRequest) GetKeyValue() *KeyValue {
	if x, ok := x.GetUnion().(*EtcdClientRequest_KeyValue); ok {
		return x.KeyValue
	}
	return nil
}

type isEtcdClientRequest_Union interface {
	isEtcdClientRequest_Union()
}

type EtcdClientRequest_Key struct {
	Key *Key `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

type EtcdClientRequest_KeyValue struct {
	KeyValue *KeyValue `protobuf:"bytes,2,opt,name=keyValue,proto3,oneof"`
}

func (*EtcdClientRequest_Key) isEtcdClientRequest_Union() {}

func (*EtcdClientRequest_KeyValue) isEtcdClientRequest_Union() {}

type EtcdClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue *KeyValue `protobuf:"bytes,1,opt,name=keyValue,proto3,oneof" json:"keyValue,omitempty"`
	Status   Status    `protobuf:"varint,2,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Error    string    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EtcdClientResponse) Reset() {
	*x = EtcdClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_etcd_client_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdClientResponse) ProtoMessage() {}

func (x *EtcdClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_etcd_client_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdClientResponse.ProtoReflect.Descriptor instead.
func (*EtcdClientResponse) Descriptor() ([]byte, []int) {
	return file_proto_etcd_client_service_proto_rawDescGZIP(), []int{3}
}

func (x *EtcdClientResponse) GetKeyValue() *KeyValue {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *EtcdClientResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *EtcdClientResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_etcd_client_service_proto protoreflect.FileDescriptor

var file_proto_etcd_client_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6b, 0x0a, 0x11, 0x45, 0x74, 0x63,
	0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x48, 0x00, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x45, 0x74, 0x63, 0x64, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x48, 0x00, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xca, 0x01, 0x0a, 0x11, 0x45, 0x74,
	0x63, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x63, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74,
	0x63, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x50, 0x75,
	0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5d, 0x0a, 0x12, 0x73, 0x75, 0x2e, 0x73, 0x76, 0x6e,
	0x2e, 0x65, 0x74, 0x63, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x13, 0x45, 0x74,
	0x63, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x73, 0x6b, 0x75, 0x72, 0x69, 0x6b, 0x68, 0x69, 0x6e,
	0x2f, 0x65, 0x74, 0x63, 0x64, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_etcd_client_service_proto_rawDescOnce sync.Once
	file_proto_etcd_client_service_proto_rawDescData = file_proto_etcd_client_service_proto_rawDesc
)

func file_proto_etcd_client_service_proto_rawDescGZIP() []byte {
	file_proto_etcd_client_service_proto_rawDescOnce.Do(func() {
		file_proto_etcd_client_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_etcd_client_service_proto_rawDescData)
	})
	return file_proto_etcd_client_service_proto_rawDescData
}

var file_proto_etcd_client_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_etcd_client_service_proto_goTypes = []any{
	(*Key)(nil),                // 0: proto.Key
	(*KeyValue)(nil),           // 1: proto.KeyValue
	(*EtcdClientRequest)(nil),  // 2: proto.EtcdClientRequest
	(*EtcdClientResponse)(nil), // 3: proto.EtcdClientResponse
	(Status)(0),                // 4: proto.Status
}
var file_proto_etcd_client_service_proto_depIdxs = []int32{
	0, // 0: proto.EtcdClientRequest.key:type_name -> proto.Key
	1, // 1: proto.EtcdClientRequest.keyValue:type_name -> proto.KeyValue
	1, // 2: proto.EtcdClientResponse.keyValue:type_name -> proto.KeyValue
	4, // 3: proto.EtcdClientResponse.status:type_name -> proto.Status
	2, // 4: proto.EtcdClientService.Delete:input_type -> proto.EtcdClientRequest
	2, // 5: proto.EtcdClientService.Get:input_type -> proto.EtcdClientRequest
	2, // 6: proto.EtcdClientService.Put:input_type -> proto.EtcdClientRequest
	3, // 7: proto.EtcdClientService.Delete:output_type -> proto.EtcdClientResponse
	3, // 8: proto.EtcdClientService.Get:output_type -> proto.EtcdClientResponse
	3, // 9: proto.EtcdClientService.Put:output_type -> proto.EtcdClientResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_etcd_client_service_proto_init() }
func file_proto_etcd_client_service_proto_init() {
	if File_proto_etcd_client_service_proto != nil {
		return
	}
	file_proto_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_etcd_client_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Key); i {
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
		file_proto_etcd_client_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*KeyValue); i {
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
		file_proto_etcd_client_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EtcdClientRequest); i {
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
		file_proto_etcd_client_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EtcdClientResponse); i {
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
	file_proto_etcd_client_service_proto_msgTypes[2].OneofWrappers = []any{
		(*EtcdClientRequest_Key)(nil),
		(*EtcdClientRequest_KeyValue)(nil),
	}
	file_proto_etcd_client_service_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_etcd_client_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_etcd_client_service_proto_goTypes,
		DependencyIndexes: file_proto_etcd_client_service_proto_depIdxs,
		MessageInfos:      file_proto_etcd_client_service_proto_msgTypes,
	}.Build()
	File_proto_etcd_client_service_proto = out.File
	file_proto_etcd_client_service_proto_rawDesc = nil
	file_proto_etcd_client_service_proto_goTypes = nil
	file_proto_etcd_client_service_proto_depIdxs = nil
}

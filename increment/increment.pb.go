// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: increment/increment.proto

package increment

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

type IncrementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IncrementRequest) Reset() {
	*x = IncrementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_increment_increment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementRequest) ProtoMessage() {}

func (x *IncrementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_increment_increment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementRequest.ProtoReflect.Descriptor instead.
func (*IncrementRequest) Descriptor() ([]byte, []int) {
	return file_increment_increment_proto_rawDescGZIP(), []int{0}
}

func (x *IncrementRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type IncrementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Before  int64 `protobuf:"varint,2,opt,name=before,proto3" json:"before,omitempty"`
}

func (x *IncrementResponse) Reset() {
	*x = IncrementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_increment_increment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementResponse) ProtoMessage() {}

func (x *IncrementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_increment_increment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementResponse.ProtoReflect.Descriptor instead.
func (*IncrementResponse) Descriptor() ([]byte, []int) {
	return file_increment_increment_proto_rawDescGZIP(), []int{1}
}

func (x *IncrementResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *IncrementResponse) GetBefore() int64 {
	if x != nil {
		return x.Before
	}
	return 0
}

type UptimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UptimeRequest) Reset() {
	*x = UptimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_increment_increment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UptimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UptimeRequest) ProtoMessage() {}

func (x *UptimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_increment_increment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UptimeRequest.ProtoReflect.Descriptor instead.
func (*UptimeRequest) Descriptor() ([]byte, []int) {
	return file_increment_increment_proto_rawDescGZIP(), []int{2}
}

type UptimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime int64 `protobuf:"varint,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
}

func (x *UptimeResponse) Reset() {
	*x = UptimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_increment_increment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UptimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UptimeResponse) ProtoMessage() {}

func (x *UptimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_increment_increment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UptimeResponse.ProtoReflect.Descriptor instead.
func (*UptimeResponse) Descriptor() ([]byte, []int) {
	return file_increment_increment_proto_rawDescGZIP(), []int{3}
}

func (x *UptimeResponse) GetUptime() int64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

var File_increment_increment_proto protoreflect.FileDescriptor

var file_increment_increment_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x45, 0x0a, 0x11, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x55, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x32, 0x94, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x12, 0x46, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_increment_increment_proto_rawDescOnce sync.Once
	file_increment_increment_proto_rawDescData = file_increment_increment_proto_rawDesc
)

func file_increment_increment_proto_rawDescGZIP() []byte {
	file_increment_increment_proto_rawDescOnce.Do(func() {
		file_increment_increment_proto_rawDescData = protoimpl.X.CompressGZIP(file_increment_increment_proto_rawDescData)
	})
	return file_increment_increment_proto_rawDescData
}

var file_increment_increment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_increment_increment_proto_goTypes = []interface{}{
	(*IncrementRequest)(nil),  // 0: increment.IncrementRequest
	(*IncrementResponse)(nil), // 1: increment.IncrementResponse
	(*UptimeRequest)(nil),     // 2: increment.UptimeRequest
	(*UptimeResponse)(nil),    // 3: increment.UptimeResponse
}
var file_increment_increment_proto_depIdxs = []int32{
	0, // 0: increment.Incrementor.Increment:input_type -> increment.IncrementRequest
	2, // 1: increment.Incrementor.Uptime:input_type -> increment.UptimeRequest
	1, // 2: increment.Incrementor.Increment:output_type -> increment.IncrementResponse
	3, // 3: increment.Incrementor.Uptime:output_type -> increment.UptimeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_increment_increment_proto_init() }
func file_increment_increment_proto_init() {
	if File_increment_increment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_increment_increment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementRequest); i {
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
		file_increment_increment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementResponse); i {
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
		file_increment_increment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UptimeRequest); i {
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
		file_increment_increment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UptimeResponse); i {
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
			RawDescriptor: file_increment_increment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_increment_increment_proto_goTypes,
		DependencyIndexes: file_increment_increment_proto_depIdxs,
		MessageInfos:      file_increment_increment_proto_msgTypes,
	}.Build()
	File_increment_increment_proto = out.File
	file_increment_increment_proto_rawDesc = nil
	file_increment_increment_proto_goTypes = nil
	file_increment_increment_proto_depIdxs = nil
}

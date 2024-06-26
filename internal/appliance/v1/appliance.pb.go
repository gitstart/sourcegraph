// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: appliance.proto

package v1

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

type GetApplianceVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetApplianceVersionRequest) Reset() {
	*x = GetApplianceVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appliance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplianceVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplianceVersionRequest) ProtoMessage() {}

func (x *GetApplianceVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appliance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplianceVersionRequest.ProtoReflect.Descriptor instead.
func (*GetApplianceVersionRequest) Descriptor() ([]byte, []int) {
	return file_appliance_proto_rawDescGZIP(), []int{0}
}

type GetApplianceVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetApplianceVersionResponse) Reset() {
	*x = GetApplianceVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appliance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplianceVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplianceVersionResponse) ProtoMessage() {}

func (x *GetApplianceVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appliance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplianceVersionResponse.ProtoReflect.Descriptor instead.
func (*GetApplianceVersionResponse) Descriptor() ([]byte, []int) {
	return file_appliance_proto_rawDescGZIP(), []int{1}
}

func (x *GetApplianceVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetApplianceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetApplianceStatusRequest) Reset() {
	*x = GetApplianceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appliance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplianceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplianceStatusRequest) ProtoMessage() {}

func (x *GetApplianceStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appliance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplianceStatusRequest.ProtoReflect.Descriptor instead.
func (*GetApplianceStatusRequest) Descriptor() ([]byte, []int) {
	return file_appliance_proto_rawDescGZIP(), []int{2}
}

type GetApplianceStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetApplianceStatusResponse) Reset() {
	*x = GetApplianceStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appliance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplianceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplianceStatusResponse) ProtoMessage() {}

func (x *GetApplianceStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appliance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplianceStatusResponse.ProtoReflect.Descriptor instead.
func (*GetApplianceStatusResponse) Descriptor() ([]byte, []int) {
	return file_appliance_proto_rawDescGZIP(), []int{3}
}

func (x *GetApplianceStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_appliance_proto protoreflect.FileDescriptor

var file_appliance_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xf1, 0x01, 0x0a, 0x10, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12,
	0x6c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_appliance_proto_rawDescOnce sync.Once
	file_appliance_proto_rawDescData = file_appliance_proto_rawDesc
)

func file_appliance_proto_rawDescGZIP() []byte {
	file_appliance_proto_rawDescOnce.Do(func() {
		file_appliance_proto_rawDescData = protoimpl.X.CompressGZIP(file_appliance_proto_rawDescData)
	})
	return file_appliance_proto_rawDescData
}

var file_appliance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_appliance_proto_goTypes = []interface{}{
	(*GetApplianceVersionRequest)(nil),  // 0: appliance.v1.GetApplianceVersionRequest
	(*GetApplianceVersionResponse)(nil), // 1: appliance.v1.GetApplianceVersionResponse
	(*GetApplianceStatusRequest)(nil),   // 2: appliance.v1.GetApplianceStatusRequest
	(*GetApplianceStatusResponse)(nil),  // 3: appliance.v1.GetApplianceStatusResponse
}
var file_appliance_proto_depIdxs = []int32{
	0, // 0: appliance.v1.ApplianceService.GetApplianceVersion:input_type -> appliance.v1.GetApplianceVersionRequest
	2, // 1: appliance.v1.ApplianceService.GetApplianceStatus:input_type -> appliance.v1.GetApplianceStatusRequest
	1, // 2: appliance.v1.ApplianceService.GetApplianceVersion:output_type -> appliance.v1.GetApplianceVersionResponse
	3, // 3: appliance.v1.ApplianceService.GetApplianceStatus:output_type -> appliance.v1.GetApplianceStatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_appliance_proto_init() }
func file_appliance_proto_init() {
	if File_appliance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appliance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplianceVersionRequest); i {
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
		file_appliance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplianceVersionResponse); i {
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
		file_appliance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplianceStatusRequest); i {
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
		file_appliance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplianceStatusResponse); i {
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
			RawDescriptor: file_appliance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appliance_proto_goTypes,
		DependencyIndexes: file_appliance_proto_depIdxs,
		MessageInfos:      file_appliance_proto_msgTypes,
	}.Build()
	File_appliance_proto = out.File
	file_appliance_proto_rawDesc = nil
	file_appliance_proto_goTypes = nil
	file_appliance_proto_depIdxs = nil
}

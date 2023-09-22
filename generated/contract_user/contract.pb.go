// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: proto/contract_user/contract.proto

package contract_user

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

type GetGroupDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorId int64 `protobuf:"varint,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
}

func (x *GetGroupDetailsRequest) Reset() {
	*x = GetGroupDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_user_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupDetailsRequest) ProtoMessage() {}

func (x *GetGroupDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_user_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetGroupDetailsRequest) Descriptor() ([]byte, []int) {
	return file_proto_contract_user_contract_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupDetailsRequest) GetVendorId() int64 {
	if x != nil {
		return x.VendorId
	}
	return 0
}

type GetGroupDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorName    string `protobuf:"bytes,1,opt,name=vendor_name,json=vendorName,proto3" json:"vendor_name,omitempty"`
	VendorAddress string `protobuf:"bytes,2,opt,name=vendor_address,json=vendorAddress,proto3" json:"vendor_address,omitempty"`
	VendorEmail   string `protobuf:"bytes,3,opt,name=vendor_email,json=vendorEmail,proto3" json:"vendor_email,omitempty"`
}

func (x *GetGroupDetailsResponse) Reset() {
	*x = GetGroupDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_user_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupDetailsResponse) ProtoMessage() {}

func (x *GetGroupDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_user_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetGroupDetailsResponse) Descriptor() ([]byte, []int) {
	return file_proto_contract_user_contract_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupDetailsResponse) GetVendorName() string {
	if x != nil {
		return x.VendorName
	}
	return ""
}

func (x *GetGroupDetailsResponse) GetVendorAddress() string {
	if x != nil {
		return x.VendorAddress
	}
	return ""
}

func (x *GetGroupDetailsResponse) GetVendorEmail() string {
	if x != nil {
		return x.VendorEmail
	}
	return ""
}

var File_proto_contract_user_contract_proto protoreflect.FileDescriptor

var file_proto_contract_user_contract_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x32, 0x73, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x76, 0x69, 0x66, 0x79, 0x2d, 0x6f, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_contract_user_contract_proto_rawDescOnce sync.Once
	file_proto_contract_user_contract_proto_rawDescData = file_proto_contract_user_contract_proto_rawDesc
)

func file_proto_contract_user_contract_proto_rawDescGZIP() []byte {
	file_proto_contract_user_contract_proto_rawDescOnce.Do(func() {
		file_proto_contract_user_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_contract_user_contract_proto_rawDescData)
	})
	return file_proto_contract_user_contract_proto_rawDescData
}

var file_proto_contract_user_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_contract_user_contract_proto_goTypes = []interface{}{
	(*GetGroupDetailsRequest)(nil),  // 0: contract_user.GetGroupDetailsRequest
	(*GetGroupDetailsResponse)(nil), // 1: contract_user.GetGroupDetailsResponse
}
var file_proto_contract_user_contract_proto_depIdxs = []int32{
	0, // 0: contract_user.ContractService.GetGroupDetails:input_type -> contract_user.GetGroupDetailsRequest
	1, // 1: contract_user.ContractService.GetGroupDetails:output_type -> contract_user.GetGroupDetailsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_contract_user_contract_proto_init() }
func file_proto_contract_user_contract_proto_init() {
	if File_proto_contract_user_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_contract_user_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupDetailsRequest); i {
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
		file_proto_contract_user_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupDetailsResponse); i {
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
			RawDescriptor: file_proto_contract_user_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_contract_user_contract_proto_goTypes,
		DependencyIndexes: file_proto_contract_user_contract_proto_depIdxs,
		MessageInfos:      file_proto_contract_user_contract_proto_msgTypes,
	}.Build()
	File_proto_contract_user_contract_proto = out.File
	file_proto_contract_user_contract_proto_rawDesc = nil
	file_proto_contract_user_contract_proto_goTypes = nil
	file_proto_contract_user_contract_proto_depIdxs = nil
}

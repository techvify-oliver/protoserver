// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: contract_request_group.proto

package generated

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

type GetUserRequestGW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Req  *GetUserRequest `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *GetUserRequestGW) Reset() {
	*x = GetUserRequestGW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_request_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequestGW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequestGW) ProtoMessage() {}

func (x *GetUserRequestGW) ProtoReflect() protoreflect.Message {
	mi := &file_contract_request_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequestGW.ProtoReflect.Descriptor instead.
func (*GetUserRequestGW) Descriptor() ([]byte, []int) {
	return file_contract_request_group_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequestGW) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetUserRequestGW) GetReq() *GetUserRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

type GetGroupDetailsRequestGW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string                  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Req  *GetGroupDetailsRequest `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *GetGroupDetailsRequestGW) Reset() {
	*x = GetGroupDetailsRequestGW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_request_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupDetailsRequestGW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupDetailsRequestGW) ProtoMessage() {}

func (x *GetGroupDetailsRequestGW) ProtoReflect() protoreflect.Message {
	mi := &file_contract_request_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupDetailsRequestGW.ProtoReflect.Descriptor instead.
func (*GetGroupDetailsRequestGW) Descriptor() ([]byte, []int) {
	return file_contract_request_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupDetailsRequestGW) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetGroupDetailsRequestGW) GetReq() *GetGroupDetailsRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

type GetWithdrawsRequestGW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Req  *GetWithdrawsRequest `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *GetWithdrawsRequestGW) Reset() {
	*x = GetWithdrawsRequestGW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_request_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawsRequestGW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawsRequestGW) ProtoMessage() {}

func (x *GetWithdrawsRequestGW) ProtoReflect() protoreflect.Message {
	mi := &file_contract_request_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawsRequestGW.ProtoReflect.Descriptor instead.
func (*GetWithdrawsRequestGW) Descriptor() ([]byte, []int) {
	return file_contract_request_group_proto_rawDescGZIP(), []int{2}
}

func (x *GetWithdrawsRequestGW) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetWithdrawsRequestGW) GetReq() *GetWithdrawsRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

type GetContractRequestGW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string              `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Req  *GetContractRequest `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *GetContractRequestGW) Reset() {
	*x = GetContractRequestGW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_request_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractRequestGW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractRequestGW) ProtoMessage() {}

func (x *GetContractRequestGW) ProtoReflect() protoreflect.Message {
	mi := &file_contract_request_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractRequestGW.ProtoReflect.Descriptor instead.
func (*GetContractRequestGW) Descriptor() ([]byte, []int) {
	return file_contract_request_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetContractRequestGW) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetContractRequestGW) GetReq() *GetContractRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

var File_contract_request_group_proto protoreflect.FileDescriptor

var file_contract_request_group_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x47, 0x57, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x72, 0x65, 0x71,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52,
	0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x63, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x57, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x5d, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x47, 0x57, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50,
	0x43, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x5b, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x57, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x32, 0xd6, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x57, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x57, 0x1a,
	0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x57, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x63, 0x68, 0x76, 0x69, 0x66, 0x79, 0x2d, 0x6f, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_request_group_proto_rawDescOnce sync.Once
	file_contract_request_group_proto_rawDescData = file_contract_request_group_proto_rawDesc
)

func file_contract_request_group_proto_rawDescGZIP() []byte {
	file_contract_request_group_proto_rawDescOnce.Do(func() {
		file_contract_request_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_request_group_proto_rawDescData)
	})
	return file_contract_request_group_proto_rawDescData
}

var file_contract_request_group_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contract_request_group_proto_goTypes = []interface{}{
	(*GetUserRequestGW)(nil),         // 0: protoGRPC.GetUserRequestGW
	(*GetGroupDetailsRequestGW)(nil), // 1: protoGRPC.GetGroupDetailsRequestGW
	(*GetWithdrawsRequestGW)(nil),    // 2: protoGRPC.GetWithdrawsRequestGW
	(*GetContractRequestGW)(nil),     // 3: protoGRPC.GetContractRequestGW
	(*GetUserRequest)(nil),           // 4: protoGRPC.GetUserRequest
	(*GetGroupDetailsRequest)(nil),   // 5: protoGRPC.GetGroupDetailsRequest
	(*GetWithdrawsRequest)(nil),      // 6: protoGRPC.GetWithdrawsRequest
	(*GetContractRequest)(nil),       // 7: protoGRPC.GetContractRequest
	(*GetUserResponse)(nil),          // 8: protoGRPC.GetUserResponse
	(*GetGroupDetailsResponse)(nil),  // 9: protoGRPC.GetGroupDetailsResponse
	(*GetWithdrawsResponse)(nil),     // 10: protoGRPC.GetWithdrawsResponse
	(*GetContractResponse)(nil),      // 11: protoGRPC.GetContractResponse
}
var file_contract_request_group_proto_depIdxs = []int32{
	4,  // 0: protoGRPC.GetUserRequestGW.req:type_name -> protoGRPC.GetUserRequest
	5,  // 1: protoGRPC.GetGroupDetailsRequestGW.req:type_name -> protoGRPC.GetGroupDetailsRequest
	6,  // 2: protoGRPC.GetWithdrawsRequestGW.req:type_name -> protoGRPC.GetWithdrawsRequest
	7,  // 3: protoGRPC.GetContractRequestGW.req:type_name -> protoGRPC.GetContractRequest
	0,  // 4: protoGRPC.Service.GetUser:input_type -> protoGRPC.GetUserRequestGW
	1,  // 5: protoGRPC.Service.GetGroupDetails:input_type -> protoGRPC.GetGroupDetailsRequestGW
	2,  // 6: protoGRPC.Service.GetRequestDetails:input_type -> protoGRPC.GetWithdrawsRequestGW
	7,  // 7: protoGRPC.Service.GetContractDetails:input_type -> protoGRPC.GetContractRequest
	8,  // 8: protoGRPC.Service.GetUser:output_type -> protoGRPC.GetUserResponse
	9,  // 9: protoGRPC.Service.GetGroupDetails:output_type -> protoGRPC.GetGroupDetailsResponse
	10, // 10: protoGRPC.Service.GetRequestDetails:output_type -> protoGRPC.GetWithdrawsResponse
	11, // 11: protoGRPC.Service.GetContractDetails:output_type -> protoGRPC.GetContractResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_contract_request_group_proto_init() }
func file_contract_request_group_proto_init() {
	if File_contract_request_group_proto != nil {
		return
	}
	file_contract_details_proto_init()
	file_group_details_proto_init()
	file_user_details_proto_init()
	file_request_details_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_request_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequestGW); i {
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
		file_contract_request_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupDetailsRequestGW); i {
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
		file_contract_request_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawsRequestGW); i {
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
		file_contract_request_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractRequestGW); i {
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
			RawDescriptor: file_contract_request_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_request_group_proto_goTypes,
		DependencyIndexes: file_contract_request_group_proto_depIdxs,
		MessageInfos:      file_contract_request_group_proto_msgTypes,
	}.Build()
	File_contract_request_group_proto = out.File
	file_contract_request_group_proto_rawDesc = nil
	file_contract_request_group_proto_goTypes = nil
	file_contract_request_group_proto_depIdxs = nil
}

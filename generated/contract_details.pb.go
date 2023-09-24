// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: contract_details.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Get contract details
type GetContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId int64 `protobuf:"varint,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
}

func (x *GetContractRequest) Reset() {
	*x = GetContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractRequest) ProtoMessage() {}

func (x *GetContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractRequest.ProtoReflect.Descriptor instead.
func (*GetContractRequest) Descriptor() ([]byte, []int) {
	return file_contract_details_proto_rawDescGZIP(), []int{0}
}

func (x *GetContractRequest) GetContractId() int64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

type GetContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorId        int64                  `protobuf:"varint,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	Status          string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	StartDate       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	TotalAmount     int32                  `protobuf:"varint,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	RemainingAmount int32                  `protobuf:"varint,6,opt,name=remaining_amount,json=remainingAmount,proto3" json:"remaining_amount,omitempty"`
}

func (x *GetContractResponse) Reset() {
	*x = GetContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractResponse) ProtoMessage() {}

func (x *GetContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractResponse.ProtoReflect.Descriptor instead.
func (*GetContractResponse) Descriptor() ([]byte, []int) {
	return file_contract_details_proto_rawDescGZIP(), []int{1}
}

func (x *GetContractResponse) GetVendorId() int64 {
	if x != nil {
		return x.VendorId
	}
	return 0
}

func (x *GetContractResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetContractResponse) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *GetContractResponse) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *GetContractResponse) GetTotalAmount() int32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *GetContractResponse) GetRemainingAmount() int32 {
	if x != nil {
		return x.RemainingAmount
	}
	return 0
}

type GetContractByGroupIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetContractByGroupIDRequest) Reset() {
	*x = GetContractByGroupIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractByGroupIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractByGroupIDRequest) ProtoMessage() {}

func (x *GetContractByGroupIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractByGroupIDRequest.ProtoReflect.Descriptor instead.
func (*GetContractByGroupIDRequest) Descriptor() ([]byte, []int) {
	return file_contract_details_proto_rawDescGZIP(), []int{2}
}

func (x *GetContractByGroupIDRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GetContractByGroupID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VendorId        int64                  `protobuf:"varint,2,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	Status          string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	StartDate       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	TotalAmount     int32                  `protobuf:"varint,6,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	RemainingAmount int32                  `protobuf:"varint,7,opt,name=remaining_amount,json=remainingAmount,proto3" json:"remaining_amount,omitempty"`
}

func (x *GetContractByGroupID) Reset() {
	*x = GetContractByGroupID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractByGroupID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractByGroupID) ProtoMessage() {}

func (x *GetContractByGroupID) ProtoReflect() protoreflect.Message {
	mi := &file_contract_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractByGroupID.ProtoReflect.Descriptor instead.
func (*GetContractByGroupID) Descriptor() ([]byte, []int) {
	return file_contract_details_proto_rawDescGZIP(), []int{3}
}

func (x *GetContractByGroupID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetContractByGroupID) GetVendorId() int64 {
	if x != nil {
		return x.VendorId
	}
	return 0
}

func (x *GetContractByGroupID) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetContractByGroupID) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *GetContractByGroupID) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *GetContractByGroupID) GetTotalAmount() int32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *GetContractByGroupID) GetRemainingAmount() int32 {
	if x != nil {
		return x.RemainingAmount
	}
	return 0
}

type GetContractByGroupIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contracts []*GetContractByGroupID `protobuf:"bytes,1,rep,name=contracts,proto3" json:"contracts,omitempty"`
}

func (x *GetContractByGroupIDResponse) Reset() {
	*x = GetContractByGroupIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_details_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractByGroupIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractByGroupIDResponse) ProtoMessage() {}

func (x *GetContractByGroupIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_details_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractByGroupIDResponse.ProtoReflect.Descriptor instead.
func (*GetContractByGroupIDResponse) Descriptor() ([]byte, []int) {
	return file_contract_details_proto_rawDescGZIP(), []int{4}
}

func (x *GetContractByGroupIDResponse) GetContracts() []*GetContractByGroupID {
	if x != nil {
		return x.Contracts
	}
	return nil
}

var File_contract_details_proto protoreflect.FileDescriptor

var file_contract_details_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47,
	0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x8a,
	0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x9b, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52,
	0x50, 0x43, 0x61, 0x67, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x76, 0x69, 0x66, 0x79, 0x2d, 0x6f, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_contract_details_proto_rawDescOnce sync.Once
	file_contract_details_proto_rawDescData = file_contract_details_proto_rawDesc
)

func file_contract_details_proto_rawDescGZIP() []byte {
	file_contract_details_proto_rawDescOnce.Do(func() {
		file_contract_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_details_proto_rawDescData)
	})
	return file_contract_details_proto_rawDescData
}

var file_contract_details_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contract_details_proto_goTypes = []interface{}{
	(*GetContractRequest)(nil),           // 0: protoGRPCagw.GetContractRequest
	(*GetContractResponse)(nil),          // 1: protoGRPCagw.GetContractResponse
	(*GetContractByGroupIDRequest)(nil),  // 2: protoGRPCagw.GetContractByGroupIDRequest
	(*GetContractByGroupID)(nil),         // 3: protoGRPCagw.GetContractByGroupID
	(*GetContractByGroupIDResponse)(nil), // 4: protoGRPCagw.GetContractByGroupIDResponse
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
}
var file_contract_details_proto_depIdxs = []int32{
	5, // 0: protoGRPCagw.GetContractResponse.start_date:type_name -> google.protobuf.Timestamp
	5, // 1: protoGRPCagw.GetContractResponse.end_date:type_name -> google.protobuf.Timestamp
	5, // 2: protoGRPCagw.GetContractByGroupID.start_date:type_name -> google.protobuf.Timestamp
	5, // 3: protoGRPCagw.GetContractByGroupID.end_date:type_name -> google.protobuf.Timestamp
	3, // 4: protoGRPCagw.GetContractByGroupIDResponse.contracts:type_name -> protoGRPCagw.GetContractByGroupID
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_contract_details_proto_init() }
func file_contract_details_proto_init() {
	if File_contract_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractRequest); i {
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
		file_contract_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractResponse); i {
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
		file_contract_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractByGroupIDRequest); i {
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
		file_contract_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractByGroupID); i {
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
		file_contract_details_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractByGroupIDResponse); i {
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
			RawDescriptor: file_contract_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_details_proto_goTypes,
		DependencyIndexes: file_contract_details_proto_depIdxs,
		MessageInfos:      file_contract_details_proto_msgTypes,
	}.Build()
	File_contract_details_proto = out.File
	file_contract_details_proto_rawDesc = nil
	file_contract_details_proto_goTypes = nil
	file_contract_details_proto_depIdxs = nil
}

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_contract_request_group_proto protoreflect.FileDescriptor

var file_contract_request_group_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x1a, 0x16, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe8, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x61, 0x67, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x63, 0x68, 0x76, 0x69, 0x66, 0x79, 0x2d, 0x6f, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_contract_request_group_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),          // 0: protoGRPCagw.GetUserRequest
	(*GetGroupDetailsRequest)(nil),  // 1: protoGRPCagw.GetGroupDetailsRequest
	(*GetWithdrawsRequest)(nil),     // 2: protoGRPCagw.GetWithdrawsRequest
	(*GetContractRequest)(nil),      // 3: protoGRPCagw.GetContractRequest
	(*GetUserResponse)(nil),         // 4: protoGRPCagw.GetUserResponse
	(*GetGroupDetailsResponse)(nil), // 5: protoGRPCagw.GetGroupDetailsResponse
	(*GetWithdrawsResponse)(nil),    // 6: protoGRPCagw.GetWithdrawsResponse
	(*GetContractResponse)(nil),     // 7: protoGRPCagw.GetContractResponse
}
var file_contract_request_group_proto_depIdxs = []int32{
	0, // 0: protoGRPCagw.Service.GetUser:input_type -> protoGRPCagw.GetUserRequest
	1, // 1: protoGRPCagw.Service.GetGroupDetails:input_type -> protoGRPCagw.GetGroupDetailsRequest
	2, // 2: protoGRPCagw.Service.GetRequestDetails:input_type -> protoGRPCagw.GetWithdrawsRequest
	3, // 3: protoGRPCagw.Service.GetContractDetails:input_type -> protoGRPCagw.GetContractRequest
	4, // 4: protoGRPCagw.Service.GetUser:output_type -> protoGRPCagw.GetUserResponse
	5, // 5: protoGRPCagw.Service.GetGroupDetails:output_type -> protoGRPCagw.GetGroupDetailsResponse
	6, // 6: protoGRPCagw.Service.GetRequestDetails:output_type -> protoGRPCagw.GetWithdrawsResponse
	7, // 7: protoGRPCagw.Service.GetContractDetails:output_type -> protoGRPCagw.GetContractResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_request_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_request_group_proto_goTypes,
		DependencyIndexes: file_contract_request_group_proto_depIdxs,
	}.Build()
	File_contract_request_group_proto = out.File
	file_contract_request_group_proto_rawDesc = nil
	file_contract_request_group_proto_goTypes = nil
	file_contract_request_group_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: apps/usercenter/service.proto

package usercenter

import (
	context "context"
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

var File_apps_usercenter_service_proto protoreflect.FileDescriptor

var file_apps_usercenter_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x76, 0x0a, 0x0d, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x72, 0x76, 0x12, 0x30, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x6f, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apps_usercenter_service_proto_goTypes = []interface{}{
	(*UserGetReq)(nil),   // 0: user.UserGetReq
	(*UserListReq)(nil),  // 1: user.UserListReq
	(*UserGetResp)(nil),  // 2: user.UserGetResp
	(*UserListResp)(nil), // 3: user.UserListResp
}
var file_apps_usercenter_service_proto_depIdxs = []int32{
	0, // 0: UserCenterSrv.UserGet:input_type -> user.UserGetReq
	1, // 1: UserCenterSrv.UserList:input_type -> user.UserListReq
	2, // 2: UserCenterSrv.UserGet:output_type -> user.UserGetResp
	3, // 3: UserCenterSrv.UserList:output_type -> user.UserListResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_usercenter_service_proto_init() }
func file_apps_usercenter_service_proto_init() {
	if File_apps_usercenter_service_proto != nil {
		return
	}
	file_apps_usercenter_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_usercenter_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_usercenter_service_proto_goTypes,
		DependencyIndexes: file_apps_usercenter_service_proto_depIdxs,
	}.Build()
	File_apps_usercenter_service_proto = out.File
	file_apps_usercenter_service_proto_rawDesc = nil
	file_apps_usercenter_service_proto_goTypes = nil
	file_apps_usercenter_service_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.4. DO NOT EDIT.

type UserCenterSrv interface {
	UserGet(ctx context.Context, req *UserGetReq) (res *UserGetResp, err error)
	UserList(ctx context.Context, req *UserListReq) (res *UserListResp, err error)
}

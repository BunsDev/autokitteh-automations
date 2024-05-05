// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: autokitteh/auth/v1/svc.proto

package authv1

import (
	v1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/users/v1"
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

type WhoAmIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhoAmIRequest) Reset() {
	*x = WhoAmIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIRequest) ProtoMessage() {}

func (x *WhoAmIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIRequest.ProtoReflect.Descriptor instead.
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_auth_v1_svc_proto_rawDescGZIP(), []int{0}
}

type WhoAmIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *v1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *WhoAmIResponse) Reset() {
	*x = WhoAmIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIResponse) ProtoMessage() {}

func (x *WhoAmIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIResponse.ProtoReflect.Descriptor instead.
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_auth_v1_svc_proto_rawDescGZIP(), []int{1}
}

func (x *WhoAmIResponse) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_auth_v1_svc_proto_rawDescGZIP(), []int{2}
}

type CreateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateTokenResponse) Reset() {
	*x = CreateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenResponse) ProtoMessage() {}

func (x *CreateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_auth_v1_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_auth_v1_svc_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_autokitteh_auth_v1_svc_proto protoreflect.FileDescriptor

var file_autokitteh_auth_v1_svc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x1a, 0x1e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xbe, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d,
	0x49, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x68, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd1, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x42, 0x08, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b,
	0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa, 0x02, 0x12, 0x41, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x41, 0x75, 0x74,
	0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autokitteh_auth_v1_svc_proto_rawDescOnce sync.Once
	file_autokitteh_auth_v1_svc_proto_rawDescData = file_autokitteh_auth_v1_svc_proto_rawDesc
)

func file_autokitteh_auth_v1_svc_proto_rawDescGZIP() []byte {
	file_autokitteh_auth_v1_svc_proto_rawDescOnce.Do(func() {
		file_autokitteh_auth_v1_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_autokitteh_auth_v1_svc_proto_rawDescData)
	})
	return file_autokitteh_auth_v1_svc_proto_rawDescData
}

var file_autokitteh_auth_v1_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_autokitteh_auth_v1_svc_proto_goTypes = []interface{}{
	(*WhoAmIRequest)(nil),       // 0: autokitteh.auth.v1.WhoAmIRequest
	(*WhoAmIResponse)(nil),      // 1: autokitteh.auth.v1.WhoAmIResponse
	(*CreateTokenRequest)(nil),  // 2: autokitteh.auth.v1.CreateTokenRequest
	(*CreateTokenResponse)(nil), // 3: autokitteh.auth.v1.CreateTokenResponse
	(*v1.User)(nil),             // 4: autokitteh.users.v1.User
}
var file_autokitteh_auth_v1_svc_proto_depIdxs = []int32{
	4, // 0: autokitteh.auth.v1.WhoAmIResponse.user:type_name -> autokitteh.users.v1.User
	0, // 1: autokitteh.auth.v1.AuthService.WhoAmI:input_type -> autokitteh.auth.v1.WhoAmIRequest
	2, // 2: autokitteh.auth.v1.AuthService.CreateToken:input_type -> autokitteh.auth.v1.CreateTokenRequest
	1, // 3: autokitteh.auth.v1.AuthService.WhoAmI:output_type -> autokitteh.auth.v1.WhoAmIResponse
	3, // 4: autokitteh.auth.v1.AuthService.CreateToken:output_type -> autokitteh.auth.v1.CreateTokenResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_autokitteh_auth_v1_svc_proto_init() }
func file_autokitteh_auth_v1_svc_proto_init() {
	if File_autokitteh_auth_v1_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autokitteh_auth_v1_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIRequest); i {
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
		file_autokitteh_auth_v1_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIResponse); i {
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
		file_autokitteh_auth_v1_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRequest); i {
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
		file_autokitteh_auth_v1_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenResponse); i {
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
			RawDescriptor: file_autokitteh_auth_v1_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autokitteh_auth_v1_svc_proto_goTypes,
		DependencyIndexes: file_autokitteh_auth_v1_svc_proto_depIdxs,
		MessageInfos:      file_autokitteh_auth_v1_svc_proto_msgTypes,
	}.Build()
	File_autokitteh_auth_v1_svc_proto = out.File
	file_autokitteh_auth_v1_svc_proto_rawDesc = nil
	file_autokitteh_auth_v1_svc_proto_goTypes = nil
	file_autokitteh_auth_v1_svc_proto_depIdxs = nil
}

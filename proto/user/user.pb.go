// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: grpc-example/proto/user/user.proto

package user

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Account int32

const (
	Account_PERSONAL     Account = 0
	Account_PROFESSIONAL Account = 1
	Account_PUBLIC       Account = 2
)

// Enum value maps for Account.
var (
	Account_name = map[int32]string{
		0: "PERSONAL",
		1: "PROFESSIONAL",
		2: "PUBLIC",
	}
	Account_value = map[string]int32{
		"PERSONAL":     0,
		"PROFESSIONAL": 1,
		"PUBLIC":       2,
	}
)

func (x Account) Enum() *Account {
	p := new(Account)
	*p = x
	return p
}

func (x Account) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Account) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_example_proto_user_user_proto_enumTypes[0].Descriptor()
}

func (Account) Type() protoreflect.EnumType {
	return &file_grpc_example_proto_user_user_proto_enumTypes[0]
}

func (x Account) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Account.Descriptor instead.
func (Account) EnumDescriptor() ([]byte, []int) {
	return file_grpc_example_proto_user_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password  string               `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Birthdate *timestamp.Timestamp `protobuf:"bytes,3,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Friends   []*User              `protobuf:"bytes,4,rep,name=friends,proto3" json:"friends,omitempty"`
	Account   Account              `protobuf:"varint,5,opt,name=account,proto3,enum=user.Account" json:"account,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_example_proto_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_example_proto_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_grpc_example_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetBirthdate() *timestamp.Timestamp {
	if x != nil {
		return x.Birthdate
	}
	return nil
}

func (x *User) GetFriends() []*User {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *User) GetAccount() Account {
	if x != nil {
		return x.Account
	}
	return Account_PERSONAL
}

var File_grpc_example_proto_user_user_proto protoreflect.FileDescriptor

var file_grpc_example_proto_user_user_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24,
	0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x35, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x52, 0x53,
	0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x46, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x02, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_example_proto_user_user_proto_rawDescOnce sync.Once
	file_grpc_example_proto_user_user_proto_rawDescData = file_grpc_example_proto_user_user_proto_rawDesc
)

func file_grpc_example_proto_user_user_proto_rawDescGZIP() []byte {
	file_grpc_example_proto_user_user_proto_rawDescOnce.Do(func() {
		file_grpc_example_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_example_proto_user_user_proto_rawDescData)
	})
	return file_grpc_example_proto_user_user_proto_rawDescData
}

var file_grpc_example_proto_user_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_example_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_example_proto_user_user_proto_goTypes = []interface{}{
	(Account)(0),                // 0: user.Account
	(*User)(nil),                // 1: user.User
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_grpc_example_proto_user_user_proto_depIdxs = []int32{
	2, // 0: user.User.birthdate:type_name -> google.protobuf.Timestamp
	1, // 1: user.User.friends:type_name -> user.User
	0, // 2: user.User.account:type_name -> user.Account
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_example_proto_user_user_proto_init() }
func file_grpc_example_proto_user_user_proto_init() {
	if File_grpc_example_proto_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_example_proto_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_grpc_example_proto_user_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_example_proto_user_user_proto_goTypes,
		DependencyIndexes: file_grpc_example_proto_user_user_proto_depIdxs,
		EnumInfos:         file_grpc_example_proto_user_user_proto_enumTypes,
		MessageInfos:      file_grpc_example_proto_user_user_proto_msgTypes,
	}.Build()
	File_grpc_example_proto_user_user_proto = out.File
	file_grpc_example_proto_user_user_proto_rawDesc = nil
	file_grpc_example_proto_user_user_proto_goTypes = nil
	file_grpc_example_proto_user_user_proto_depIdxs = nil
}

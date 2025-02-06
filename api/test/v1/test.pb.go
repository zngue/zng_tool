// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: api/test/v1/test.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserType int32

const (
	UserType_USER_TYPE_UNSPECIFIED UserType = 0
	UserType_USER_TYPE_ADMIN       UserType = 1
	UserType_USER_TYPE_USER        UserType = 2
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "USER_TYPE_UNSPECIFIED",
		1: "USER_TYPE_ADMIN",
		2: "USER_TYPE_USER",
	}
	UserType_value = map[string]int32{
		"USER_TYPE_UNSPECIFIED": 0,
		"USER_TYPE_ADMIN":       1,
		"USER_TYPE_USER":        2,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_test_v1_test_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_api_test_v1_test_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{0}
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    // id
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // 姓名
	Age   int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Sex   int32  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Image string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateUserRequest) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UpdateUserRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetList2TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status   int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Age      int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Page     int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Order    int32  `protobuf:"varint,6,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetList2TestRequest) Reset() {
	*x = GetList2TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetList2TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetList2TestRequest) ProtoMessage() {}

func (x *GetList2TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetList2TestRequest.ProtoReflect.Descriptor instead.
func (*GetList2TestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *GetList2TestRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetList2TestRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetList2TestRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *GetList2TestRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetList2TestRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetList2TestRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type UserList2Rely struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Total     int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page      int32   `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int32   `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Order     int32   `protobuf:"varint,5,opt,name=order,proto3" json:"order,omitempty"`
	TotalPage int32   `protobuf:"varint,6,opt,name=totalPage,proto3" json:"totalPage,omitempty"`
	Status    int32   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Age       int32   `protobuf:"varint,8,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserList2Rely) Reset() {
	*x = UserList2Rely{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList2Rely) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList2Rely) ProtoMessage() {}

func (x *UserList2Rely) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList2Rely.ProtoReflect.Descriptor instead.
func (*UserList2Rely) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{2}
}

func (x *UserList2Rely) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserList2Rely) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserList2Rely) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UserList2Rely) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *UserList2Rely) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *UserList2Rely) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *UserList2Rely) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserList2Rely) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type GetListTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetListTestRequest) Reset() {
	*x = GetListTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListTestRequest) ProtoMessage() {}

func (x *GetListTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListTestRequest.ProtoReflect.Descriptor instead.
func (*GetListTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{3}
}

func (x *GetListTestRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{4}
}

func (x *GetTestRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTestRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetTestReply) Reset() {
	*x = GetTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestReply) ProtoMessage() {}

func (x *GetTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestReply.ProtoReflect.Descriptor instead.
func (*GetTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{5}
}

func (x *GetTestReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age      int32     `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	UserExit *UserExit `protobuf:"bytes,3,opt,name=userExit,proto3" json:"userExit,omitempty"`
	Data     *Data     `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[6]
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
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetUserExit() *UserExit {
	if x != nil {
		return x.UserExit
	}
	return nil
}

func (x *User) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{7}
}

func (x *Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserExit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image  string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Height int32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Weight int32  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *UserExit) Reset() {
	*x = UserExit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserExit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserExit) ProtoMessage() {}

func (x *UserExit) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserExit.ProtoReflect.Descriptor instead.
func (*UserExit) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{8}
}

func (x *UserExit) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UserExit) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UserExit) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Total int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Abc   *Abc    `protobuf:"bytes,3,opt,name=abc,proto3" json:"abc,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{9}
}

func (x *UserList) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserList) GetAbc() *Abc {
	if x != nil {
		return x.Abc
	}
	return nil
}

type Abc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Abc) Reset() {
	*x = Abc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Abc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Abc) ProtoMessage() {}

func (x *Abc) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Abc.ProtoReflect.Descriptor instead.
func (*Abc) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{10}
}

func (x *Abc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_test_v1_test_proto protoreflect.FileDescriptor

var file_api_test_v1_test_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xaa, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x1a, 0x0a, 0x30,
	0x00, 0x30, 0x01, 0x30, 0x02, 0x30, 0x03, 0x30, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x52, 0x65, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0f, 0xfa, 0x42, 0x0c, 0x1a, 0x0a, 0x30, 0x00, 0x30, 0x01, 0x30, 0x02, 0x30, 0x03, 0x30, 0x04,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x45, 0x78, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1a, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x78, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x6b, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x22, 0x0a, 0x03, 0x61, 0x62, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x62,
	0x63, 0x52, 0x03, 0x61, 0x62, 0x63, 0x22, 0x19, 0x0a, 0x03, 0x41, 0x62, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x2a, 0x4e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10,
	0x02, 0x32, 0xd0, 0x04, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x6e, 0x2d, 0x70, 0x62, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x5b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x6e, 0x2d, 0x70, 0x62, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x5e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x69, 0x6e, 0x2d, 0x70, 0x62, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x60, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x69, 0x6e, 0x2d, 0x70, 0x62, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x68, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x32, 0x12,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x52, 0x65, 0x6c, 0x79, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x6e, 0x2d, 0x70,
	0x62, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x32, 0x12, 0x66, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x69, 0x6e, 0x2d, 0x70, 0x62, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x3a, 0x01, 0x2a, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x6e, 0x67, 0x75, 0x65, 0x2f, 0x7a, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6f,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_test_v1_test_proto_rawDescOnce sync.Once
	file_api_test_v1_test_proto_rawDescData = file_api_test_v1_test_proto_rawDesc
)

func file_api_test_v1_test_proto_rawDescGZIP() []byte {
	file_api_test_v1_test_proto_rawDescOnce.Do(func() {
		file_api_test_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_test_v1_test_proto_rawDescData)
	})
	return file_api_test_v1_test_proto_rawDescData
}

var file_api_test_v1_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_test_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_test_v1_test_proto_goTypes = []interface{}{
	(UserType)(0),               // 0: api.test.v1.UserType
	(*UpdateUserRequest)(nil),   // 1: api.test.v1.UpdateUserRequest
	(*GetList2TestRequest)(nil), // 2: api.test.v1.GetList2TestRequest
	(*UserList2Rely)(nil),       // 3: api.test.v1.UserList2Rely
	(*GetListTestRequest)(nil),  // 4: api.test.v1.GetListTestRequest
	(*GetTestRequest)(nil),      // 5: api.test.v1.GetTestRequest
	(*GetTestReply)(nil),        // 6: api.test.v1.GetTestReply
	(*User)(nil),                // 7: api.test.v1.User
	(*Data)(nil),                // 8: api.test.v1.Data
	(*UserExit)(nil),            // 9: api.test.v1.UserExit
	(*UserList)(nil),            // 10: api.test.v1.UserList
	(*Abc)(nil),                 // 11: api.test.v1.Abc
	(*empty.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_api_test_v1_test_proto_depIdxs = []int32{
	7,  // 0: api.test.v1.UserList2Rely.user:type_name -> api.test.v1.User
	9,  // 1: api.test.v1.User.userExit:type_name -> api.test.v1.UserExit
	8,  // 2: api.test.v1.User.data:type_name -> api.test.v1.Data
	7,  // 3: api.test.v1.UserList.user:type_name -> api.test.v1.User
	11, // 4: api.test.v1.UserList.abc:type_name -> api.test.v1.Abc
	5,  // 5: api.test.v1.Test.Info:input_type -> api.test.v1.GetTestRequest
	4,  // 6: api.test.v1.Test.List:input_type -> api.test.v1.GetListTestRequest
	5,  // 7: api.test.v1.Test.Create:input_type -> api.test.v1.GetTestRequest
	4,  // 8: api.test.v1.Test.ListUser:input_type -> api.test.v1.GetListTestRequest
	2,  // 9: api.test.v1.Test.ListUser2:input_type -> api.test.v1.GetList2TestRequest
	1,  // 10: api.test.v1.Test.UpdateUser:input_type -> api.test.v1.UpdateUserRequest
	6,  // 11: api.test.v1.Test.Info:output_type -> api.test.v1.GetTestReply
	6,  // 12: api.test.v1.Test.List:output_type -> api.test.v1.GetTestReply
	6,  // 13: api.test.v1.Test.Create:output_type -> api.test.v1.GetTestReply
	10, // 14: api.test.v1.Test.ListUser:output_type -> api.test.v1.UserList
	3,  // 15: api.test.v1.Test.ListUser2:output_type -> api.test.v1.UserList2Rely
	12, // 16: api.test.v1.Test.UpdateUser:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_test_v1_test_proto_init() }
func file_api_test_v1_test_proto_init() {
	if File_api_test_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_test_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_api_test_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetList2TestRequest); i {
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
		file_api_test_v1_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList2Rely); i {
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
		file_api_test_v1_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListTestRequest); i {
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
		file_api_test_v1_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestRequest); i {
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
		file_api_test_v1_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestReply); i {
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
		file_api_test_v1_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_test_v1_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_api_test_v1_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserExit); i {
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
		file_api_test_v1_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
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
		file_api_test_v1_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Abc); i {
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
			RawDescriptor: file_api_test_v1_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_test_v1_test_proto_goTypes,
		DependencyIndexes: file_api_test_v1_test_proto_depIdxs,
		EnumInfos:         file_api_test_v1_test_proto_enumTypes,
		MessageInfos:      file_api_test_v1_test_proto_msgTypes,
	}.Build()
	File_api_test_v1_test_proto = out.File
	file_api_test_v1_test_proto_rawDesc = nil
	file_api_test_v1_test_proto_goTypes = nil
	file_api_test_v1_test_proto_depIdxs = nil
}

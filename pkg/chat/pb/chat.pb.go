// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: pkg/chat/pb/chat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User1ID int64 `protobuf:"varint,1,opt,name=user1ID,proto3" json:"user1ID,omitempty"`
	User2ID int64 `protobuf:"varint,2,opt,name=user2ID,proto3" json:"user2ID,omitempty"`
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChatRequest) GetUser1ID() int64 {
	if x != nil {
		return x.User1ID
	}
	return 0
}

func (x *CreateChatRequest) GetUser2ID() int64 {
	if x != nil {
		return x.User2ID
	}
	return 0
}

type CreateChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ChatId int64  `protobuf:"varint,3,opt,name=chatId,proto3" json:"chatId,omitempty"`
}

func (x *CreateChatResponse) Reset() {
	*x = CreateChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatResponse) ProtoMessage() {}

func (x *CreateChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatResponse.ProtoReflect.Descriptor instead.
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChatResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateChatResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateChatResponse) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type FindChatData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
}

func (x *FindChatData) Reset() {
	*x = FindChatData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindChatData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindChatData) ProtoMessage() {}

func (x *FindChatData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindChatData.ProtoReflect.Descriptor instead.
func (*FindChatData) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{2}
}

func (x *FindChatData) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type FindChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User1Id int64 `protobuf:"varint,1,opt,name=user1Id,proto3" json:"user1Id,omitempty"`
	User2Id int64 `protobuf:"varint,2,opt,name=user2Id,proto3" json:"user2Id,omitempty"`
}

func (x *FindChatRequest) Reset() {
	*x = FindChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindChatRequest) ProtoMessage() {}

func (x *FindChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindChatRequest.ProtoReflect.Descriptor instead.
func (*FindChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{3}
}

func (x *FindChatRequest) GetUser1Id() int64 {
	if x != nil {
		return x.User1Id
	}
	return 0
}

func (x *FindChatRequest) GetUser2Id() int64 {
	if x != nil {
		return x.User2Id
	}
	return 0
}

type FindChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *FindChatData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindChatResponse) Reset() {
	*x = FindChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindChatResponse) ProtoMessage() {}

func (x *FindChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindChatResponse.ProtoReflect.Descriptor instead.
func (*FindChatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{4}
}

func (x *FindChatResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindChatResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindChatResponse) GetData() *FindChatData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
}

func (x *DeleteChatRequest) Reset() {
	*x = DeleteChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRequest) ProtoMessage() {}

func (x *DeleteChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteChatRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type DeleteChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteChatResponse) Reset() {
	*x = DeleteChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatResponse) ProtoMessage() {}

func (x *DeleteChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatResponse.ProtoReflect.Descriptor instead.
func (*DeleteChatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteChatResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteChatResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AddMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFromId       int64  `protobuf:"varint,1,opt,name=userFromId,proto3" json:"userFromId,omitempty"`
	UserFromLanguage string `protobuf:"bytes,2,opt,name=userFromLanguage,proto3" json:"userFromLanguage,omitempty"`
	UserToId         int64  `protobuf:"varint,3,opt,name=userToId,proto3" json:"userToId,omitempty"`
	UserToLanguage   string `protobuf:"bytes,4,opt,name=userToLanguage,proto3" json:"userToLanguage,omitempty"`
	Message          string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddMessageRequest) Reset() {
	*x = AddMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageRequest) ProtoMessage() {}

func (x *AddMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageRequest.ProtoReflect.Descriptor instead.
func (*AddMessageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{7}
}

func (x *AddMessageRequest) GetUserFromId() int64 {
	if x != nil {
		return x.UserFromId
	}
	return 0
}

func (x *AddMessageRequest) GetUserFromLanguage() string {
	if x != nil {
		return x.UserFromLanguage
	}
	return ""
}

func (x *AddMessageRequest) GetUserToId() int64 {
	if x != nil {
		return x.UserToId
	}
	return 0
}

func (x *AddMessageRequest) GetUserToLanguage() string {
	if x != nil {
		return x.UserToLanguage
	}
	return ""
}

func (x *AddMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AddMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddMessageResponse) Reset() {
	*x = AddMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chat_pb_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageResponse) ProtoMessage() {}

func (x *AddMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chat_pb_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageResponse.ProtoReflect.Descriptor instead.
func (*AddMessageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_chat_pb_chat_proto_rawDescGZIP(), []int{8}
}

func (x *AddMessageResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddMessageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_chat_pb_chat_proto protoreflect.FileDescriptor

var file_pkg_chat_pb_chat_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x47,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x31, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x31, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x32, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x32, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x46,
	0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x31, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x31, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x32, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x32,
	0x49, 0x64, 0x22, 0x68, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xbd, 0x01,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75,
	0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x75,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x42, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0x93, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_chat_pb_chat_proto_rawDescOnce sync.Once
	file_pkg_chat_pb_chat_proto_rawDescData = file_pkg_chat_pb_chat_proto_rawDesc
)

func file_pkg_chat_pb_chat_proto_rawDescGZIP() []byte {
	file_pkg_chat_pb_chat_proto_rawDescOnce.Do(func() {
		file_pkg_chat_pb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_chat_pb_chat_proto_rawDescData)
	})
	return file_pkg_chat_pb_chat_proto_rawDescData
}

var file_pkg_chat_pb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_chat_pb_chat_proto_goTypes = []interface{}{
	(*CreateChatRequest)(nil),  // 0: chat.CreateChatRequest
	(*CreateChatResponse)(nil), // 1: chat.CreateChatResponse
	(*FindChatData)(nil),       // 2: chat.FindChatData
	(*FindChatRequest)(nil),    // 3: chat.FindChatRequest
	(*FindChatResponse)(nil),   // 4: chat.FindChatResponse
	(*DeleteChatRequest)(nil),  // 5: chat.DeleteChatRequest
	(*DeleteChatResponse)(nil), // 6: chat.DeleteChatResponse
	(*AddMessageRequest)(nil),  // 7: chat.AddMessageRequest
	(*AddMessageResponse)(nil), // 8: chat.AddMessageResponse
}
var file_pkg_chat_pb_chat_proto_depIdxs = []int32{
	2, // 0: chat.FindChatResponse.data:type_name -> chat.FindChatData
	0, // 1: chat.ChatService.CreateChat:input_type -> chat.CreateChatRequest
	3, // 2: chat.ChatService.FindChat:input_type -> chat.FindChatRequest
	5, // 3: chat.ChatService.DeleteChat:input_type -> chat.DeleteChatRequest
	7, // 4: chat.ChatService.AddMessage:input_type -> chat.AddMessageRequest
	1, // 5: chat.ChatService.CreateChat:output_type -> chat.CreateChatResponse
	4, // 6: chat.ChatService.FindChat:output_type -> chat.FindChatResponse
	6, // 7: chat.ChatService.DeleteChat:output_type -> chat.DeleteChatResponse
	8, // 8: chat.ChatService.AddMessage:output_type -> chat.AddMessageResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_chat_pb_chat_proto_init() }
func file_pkg_chat_pb_chat_proto_init() {
	if File_pkg_chat_pb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_chat_pb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatRequest); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatResponse); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindChatData); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindChatRequest); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindChatResponse); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRequest); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatResponse); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMessageRequest); i {
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
		file_pkg_chat_pb_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMessageResponse); i {
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
			RawDescriptor: file_pkg_chat_pb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_chat_pb_chat_proto_goTypes,
		DependencyIndexes: file_pkg_chat_pb_chat_proto_depIdxs,
		MessageInfos:      file_pkg_chat_pb_chat_proto_msgTypes,
	}.Build()
	File_pkg_chat_pb_chat_proto = out.File
	file_pkg_chat_pb_chat_proto_rawDesc = nil
	file_pkg_chat_pb_chat_proto_goTypes = nil
	file_pkg_chat_pb_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	FindChat(ctx context.Context, in *FindChatRequest, opts ...grpc.CallOption) (*FindChatResponse, error)
	DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*DeleteChatResponse, error)
	AddMessage(ctx context.Context, in *AddMessageRequest, opts ...grpc.CallOption) (*AddMessageResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) FindChat(ctx context.Context, in *FindChatRequest, opts ...grpc.CallOption) (*FindChatResponse, error) {
	out := new(FindChatResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/FindChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*DeleteChatResponse, error) {
	out := new(DeleteChatResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/DeleteChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddMessage(ctx context.Context, in *AddMessageRequest, opts ...grpc.CallOption) (*AddMessageResponse, error) {
	out := new(AddMessageResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/AddMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	FindChat(context.Context, *FindChatRequest) (*FindChatResponse, error)
	DeleteChat(context.Context, *DeleteChatRequest) (*DeleteChatResponse, error)
	AddMessage(context.Context, *AddMessageRequest) (*AddMessageResponse, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (*UnimplementedChatServiceServer) FindChat(context.Context, *FindChatRequest) (*FindChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindChat not implemented")
}
func (*UnimplementedChatServiceServer) DeleteChat(context.Context, *DeleteChatRequest) (*DeleteChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (*UnimplementedChatServiceServer) AddMessage(context.Context, *AddMessageRequest) (*AddMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_FindChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).FindChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/FindChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).FindChat(ctx, req.(*FindChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/DeleteChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteChat(ctx, req.(*DeleteChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/AddMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddMessage(ctx, req.(*AddMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "FindChat",
			Handler:    _ChatService_FindChat_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _ChatService_DeleteChat_Handler,
		},
		{
			MethodName: "AddMessage",
			Handler:    _ChatService_AddMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/chat/pb/chat.proto",
}

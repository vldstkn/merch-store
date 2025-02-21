// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: account.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	mi := &file_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRes) Reset() {
	*x = AuthRes{}
	mi := &file_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRes) ProtoMessage() {}

func (x *AuthRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRes.ProtoReflect.Descriptor instead.
func (*AuthRes) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetInfoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoReq) Reset() {
	*x = GetInfoReq{}
	mi := &file_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoReq) ProtoMessage() {}

func (x *GetInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoReq.ProtoReflect.Descriptor instead.
func (*GetInfoReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type GetInfoRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Coins         int64                  `protobuf:"varint,1,opt,name=coins,proto3" json:"coins,omitempty"`
	Inventory     []*Inventory           `protobuf:"bytes,2,rep,name=inventory,proto3" json:"inventory,omitempty"`
	CoinsHistory  *CoinsHistory          `protobuf:"bytes,3,opt,name=coinsHistory,json=coinHistory,proto3" json:"coinsHistory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoRes) Reset() {
	*x = GetInfoRes{}
	mi := &file_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRes) ProtoMessage() {}

func (x *GetInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRes.ProtoReflect.Descriptor instead.
func (*GetInfoRes) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetInfoRes) GetCoins() int64 {
	if x != nil {
		return x.Coins
	}
	return 0
}

func (x *GetInfoRes) GetInventory() []*Inventory {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *GetInfoRes) GetCoinsHistory() *CoinsHistory {
	if x != nil {
		return x.CoinsHistory
	}
	return nil
}

type DeductBalanceReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Amount        int64                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeductBalanceReq) Reset() {
	*x = DeductBalanceReq{}
	mi := &file_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductBalanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductBalanceReq) ProtoMessage() {}

func (x *DeductBalanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductBalanceReq.ProtoReflect.Descriptor instead.
func (*DeductBalanceReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *DeductBalanceReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *DeductBalanceReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DeductBalanceRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeductBalanceRes) Reset() {
	*x = DeductBalanceRes{}
	mi := &file_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductBalanceRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductBalanceRes) ProtoMessage() {}

func (x *DeductBalanceRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductBalanceRes.ProtoReflect.Descriptor instead.
func (*DeductBalanceRes) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

type RefundReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Amount        int64                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefundReq) Reset() {
	*x = RefundReq{}
	mi := &file_account_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefundReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundReq) ProtoMessage() {}

func (x *RefundReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundReq.ProtoReflect.Descriptor instead.
func (*RefundReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{6}
}

func (x *RefundReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RefundReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type RefundRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefundRes) Reset() {
	*x = RefundRes{}
	mi := &file_account_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefundRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRes) ProtoMessage() {}

func (x *RefundRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRes.ProtoReflect.Descriptor instead.
func (*RefundRes) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{7}
}

type TransferCoinsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserFromName  string                 `protobuf:"bytes,1,opt,name=userFromName,proto3" json:"userFromName,omitempty"`
	UserToName    string                 `protobuf:"bytes,2,opt,name=userToName,proto3" json:"userToName,omitempty"`
	Amount        int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferCoinsReq) Reset() {
	*x = TransferCoinsReq{}
	mi := &file_account_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferCoinsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferCoinsReq) ProtoMessage() {}

func (x *TransferCoinsReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferCoinsReq.ProtoReflect.Descriptor instead.
func (*TransferCoinsReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{8}
}

func (x *TransferCoinsReq) GetUserFromName() string {
	if x != nil {
		return x.UserFromName
	}
	return ""
}

func (x *TransferCoinsReq) GetUserToName() string {
	if x != nil {
		return x.UserToName
	}
	return ""
}

func (x *TransferCoinsReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferCoinsRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferCoinsRes) Reset() {
	*x = TransferCoinsRes{}
	mi := &file_account_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferCoinsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferCoinsRes) ProtoMessage() {}

func (x *TransferCoinsRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferCoinsRes.ProtoReflect.Descriptor instead.
func (*TransferCoinsRes) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{9}
}

var file_account_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "go_tag",
		Tag:           "bytes,50001,opt,name=go_tag",
		Filename:      "account.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string go_tag = 50001;
	E_GoTag = &file_account_proto_extTypes[0]
)

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x41, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1f, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x28, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x8f, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x42, 0x0f, 0x8a, 0xb5, 0x18, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x30, 0x0a, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x22, 0x46, 0x0a, 0x10, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x22, 0x3f,
	0x0a, 0x09, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x0b, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x32, 0xda, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x08, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x08,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a,
	0x0d, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x11,
	0x2e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x0a,
	0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x11, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x3a, 0x36, 0x0a,
	0x06, 0x67, 0x6f, 0x5f, 0x74, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x6f, 0x54, 0x61, 0x67, 0x42, 0x14, 0x5a, 0x12, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x2d, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData []byte
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_proto_rawDesc), len(file_account_proto_rawDesc)))
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_account_proto_goTypes = []any{
	(*AuthReq)(nil),                   // 0: AuthReq
	(*AuthRes)(nil),                   // 1: AuthRes
	(*GetInfoReq)(nil),                // 2: GetInfoReq
	(*GetInfoRes)(nil),                // 3: GetInfoRes
	(*DeductBalanceReq)(nil),          // 4: DeductBalanceReq
	(*DeductBalanceRes)(nil),          // 5: DeductBalanceRes
	(*RefundReq)(nil),                 // 6: RefundReq
	(*RefundRes)(nil),                 // 7: RefundRes
	(*TransferCoinsReq)(nil),          // 8: TransferCoinsReq
	(*TransferCoinsRes)(nil),          // 9: TransferCoinsRes
	(*Inventory)(nil),                 // 10: Inventory
	(*CoinsHistory)(nil),              // 11: CoinsHistory
	(*descriptorpb.FieldOptions)(nil), // 12: google.protobuf.FieldOptions
}
var file_account_proto_depIdxs = []int32{
	10, // 0: GetInfoRes.inventory:type_name -> Inventory
	11, // 1: GetInfoRes.coinsHistory:type_name -> CoinsHistory
	12, // 2: go_tag:extendee -> google.protobuf.FieldOptions
	0,  // 3: Account.Auth:input_type -> AuthReq
	2,  // 4: Account.GetInfo:input_type -> GetInfoReq
	4,  // 5: Account.DeductBalance:input_type -> DeductBalanceReq
	6,  // 6: Account.Refund:input_type -> RefundReq
	8,  // 7: Account.TransferCoins:input_type -> TransferCoinsReq
	1,  // 8: Account.Auth:output_type -> AuthRes
	3,  // 9: Account.GetInfo:output_type -> GetInfoRes
	5,  // 10: Account.DeductBalance:output_type -> DeductBalanceRes
	7,  // 11: Account.Refund:output_type -> RefundRes
	9,  // 12: Account.TransferCoins:output_type -> TransferCoinsRes
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	2,  // [2:3] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	file_products_proto_init()
	file_transfers_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_proto_rawDesc), len(file_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
		ExtensionInfos:    file_account_proto_extTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: api/grpc/captcha/captcha.proto

package captcha

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

// 通过purpose字段限制用途
type Purpose int32

const (
	Purpose_SIGN_UP      Purpose = 0 // 注册
	Purpose_SIGN_IN      Purpose = 1 // 登入
	Purpose_RESET_PASSWD Purpose = 2 // 重置密码
)

// Enum value maps for Purpose.
var (
	Purpose_name = map[int32]string{
		0: "SIGN_UP",
		1: "SIGN_IN",
		2: "RESET_PASSWD",
	}
	Purpose_value = map[string]int32{
		"SIGN_UP":      0,
		"SIGN_IN":      1,
		"RESET_PASSWD": 2,
	}
)

func (x Purpose) Enum() *Purpose {
	p := new(Purpose)
	*p = x
	return p
}

func (x Purpose) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Purpose) Descriptor() protoreflect.EnumDescriptor {
	return file_api_grpc_captcha_captcha_proto_enumTypes[0].Descriptor()
}

func (Purpose) Type() protoreflect.EnumType {
	return &file_api_grpc_captcha_captcha_proto_enumTypes[0]
}

func (x Purpose) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Purpose.Descriptor instead.
func (Purpose) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{0}
}

type Captcha struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Purpose  Purpose `protobuf:"varint,2,opt,name=purpose,proto3,enum=captcha.Purpose" json:"purpose,omitempty"`
	Content  string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Message  string  `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Deadline int64   `protobuf:"varint,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *Captcha) Reset() {
	*x = Captcha{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_captcha_captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Captcha) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Captcha) ProtoMessage() {}

func (x *Captcha) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_captcha_captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Captcha.ProtoReflect.Descriptor instead.
func (*Captcha) Descriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *Captcha) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Captcha) GetPurpose() Purpose {
	if x != nil {
		return x.Purpose
	}
	return Purpose_SIGN_UP
}

func (x *Captcha) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Captcha) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Captcha) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

type AsyncCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// email or phone
	Contact string  `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	Purpose Purpose `protobuf:"varint,2,opt,name=purpose,proto3,enum=captcha.Purpose" json:"purpose,omitempty"`
}

func (x *AsyncCodeReq) Reset() {
	*x = AsyncCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_captcha_captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncCodeReq) ProtoMessage() {}

func (x *AsyncCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_captcha_captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncCodeReq.ProtoReflect.Descriptor instead.
func (*AsyncCodeReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{1}
}

func (x *AsyncCodeReq) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *AsyncCodeReq) GetPurpose() Purpose {
	if x != nil {
		return x.Purpose
	}
	return Purpose_SIGN_UP
}

type AsyncCodeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AsyncCodeRsp) Reset() {
	*x = AsyncCodeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_captcha_captcha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncCodeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncCodeRsp) ProtoMessage() {}

func (x *AsyncCodeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_captcha_captcha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncCodeRsp.ProtoReflect.Descriptor instead.
func (*AsyncCodeRsp) Descriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{2}
}

type GraphCaptchaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purpose Purpose `protobuf:"varint,1,opt,name=purpose,proto3,enum=captcha.Purpose" json:"purpose,omitempty"`
}

func (x *GraphCaptchaReq) Reset() {
	*x = GraphCaptchaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_captcha_captcha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphCaptchaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphCaptchaReq) ProtoMessage() {}

func (x *GraphCaptchaReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_captcha_captcha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphCaptchaReq.ProtoReflect.Descriptor instead.
func (*GraphCaptchaReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{3}
}

func (x *GraphCaptchaReq) GetPurpose() Purpose {
	if x != nil {
		return x.Purpose
	}
	return Purpose_SIGN_UP
}

type GraphCaptchaRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Captcha *Captcha `protobuf:"bytes,1,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *GraphCaptchaRsp) Reset() {
	*x = GraphCaptchaRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_captcha_captcha_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphCaptchaRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphCaptchaRsp) ProtoMessage() {}

func (x *GraphCaptchaRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_captcha_captcha_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphCaptchaRsp.ProtoReflect.Descriptor instead.
func (*GraphCaptchaRsp) Descriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{4}
}

func (x *GraphCaptchaRsp) GetCaptcha() *Captcha {
	if x != nil {
		return x.Captcha
	}
	return nil
}

type VerifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   string  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Purpose Purpose `protobuf:"varint,3,opt,name=purpose,proto3,enum=captcha.Purpose" json:"purpose,omitempty"`
}

func (x *VerifyReq) Reset() {
	*x = VerifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_captcha_captcha_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyReq) ProtoMessage() {}

func (x *VerifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_captcha_captcha_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyReq.ProtoReflect.Descriptor instead.
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VerifyReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *VerifyReq) GetPurpose() Purpose {
	if x != nil {
		return x.Purpose
	}
	return Purpose_SIGN_UP
}

type VerifyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *VerifyRsp) Reset() {
	*x = VerifyRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_captcha_captcha_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRsp) ProtoMessage() {}

func (x *VerifyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_captcha_captcha_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRsp.ProtoReflect.Descriptor instead.
func (*VerifyRsp) Descriptor() ([]byte, []int) {
	return file_api_grpc_captcha_captcha_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyRsp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_grpc_captcha_captcha_proto protoreflect.FileDescriptor

var file_api_grpc_captcha_captcha_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x2e, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70,
	0x6f, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x54, 0x0a, 0x0c, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2a, 0x0a,
	0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x73, 0x79,
	0x6e, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x73, 0x70, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x07,
	0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x52,
	0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x07,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x5f, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x07,
	0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x52,
	0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a,
	0x35, 0x0a, 0x07, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x49,
	0x47, 0x4e, 0x5f, 0x55, 0x50, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x49, 0x47, 0x4e, 0x5f,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x44, 0x10, 0x02, 0x32, 0xc1, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x41, 0x73, 0x79,
	0x6e, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x18,
	0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x12, 0x12, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x73, 0x70, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x69, 0x6e, 0x2d, 0x63, 0x68,
	0x2f, 0x67, 0x72, 0x69, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_grpc_captcha_captcha_proto_rawDescOnce sync.Once
	file_api_grpc_captcha_captcha_proto_rawDescData = file_api_grpc_captcha_captcha_proto_rawDesc
)

func file_api_grpc_captcha_captcha_proto_rawDescGZIP() []byte {
	file_api_grpc_captcha_captcha_proto_rawDescOnce.Do(func() {
		file_api_grpc_captcha_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_captcha_captcha_proto_rawDescData)
	})
	return file_api_grpc_captcha_captcha_proto_rawDescData
}

var file_api_grpc_captcha_captcha_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_grpc_captcha_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_grpc_captcha_captcha_proto_goTypes = []interface{}{
	(Purpose)(0),            // 0: captcha.Purpose
	(*Captcha)(nil),         // 1: captcha.Captcha
	(*AsyncCodeReq)(nil),    // 2: captcha.AsyncCodeReq
	(*AsyncCodeRsp)(nil),    // 3: captcha.AsyncCodeRsp
	(*GraphCaptchaReq)(nil), // 4: captcha.GraphCaptchaReq
	(*GraphCaptchaRsp)(nil), // 5: captcha.GraphCaptchaRsp
	(*VerifyReq)(nil),       // 6: captcha.VerifyReq
	(*VerifyRsp)(nil),       // 7: captcha.VerifyRsp
}
var file_api_grpc_captcha_captcha_proto_depIdxs = []int32{
	0, // 0: captcha.Captcha.purpose:type_name -> captcha.Purpose
	0, // 1: captcha.AsyncCodeReq.purpose:type_name -> captcha.Purpose
	0, // 2: captcha.GraphCaptchaReq.purpose:type_name -> captcha.Purpose
	1, // 3: captcha.GraphCaptchaRsp.captcha:type_name -> captcha.Captcha
	0, // 4: captcha.VerifyReq.purpose:type_name -> captcha.Purpose
	2, // 5: captcha.CaptchaService.AsyncCode:input_type -> captcha.AsyncCodeReq
	4, // 6: captcha.CaptchaService.GraphCaptcha:input_type -> captcha.GraphCaptchaReq
	6, // 7: captcha.CaptchaService.Verify:input_type -> captcha.VerifyReq
	3, // 8: captcha.CaptchaService.AsyncCode:output_type -> captcha.AsyncCodeRsp
	5, // 9: captcha.CaptchaService.GraphCaptcha:output_type -> captcha.GraphCaptchaRsp
	7, // 10: captcha.CaptchaService.Verify:output_type -> captcha.VerifyRsp
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_grpc_captcha_captcha_proto_init() }
func file_api_grpc_captcha_captcha_proto_init() {
	if File_api_grpc_captcha_captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_captcha_captcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Captcha); i {
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
		file_api_grpc_captcha_captcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncCodeReq); i {
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
		file_api_grpc_captcha_captcha_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncCodeRsp); i {
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
		file_api_grpc_captcha_captcha_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphCaptchaReq); i {
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
		file_api_grpc_captcha_captcha_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphCaptchaRsp); i {
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
		file_api_grpc_captcha_captcha_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyReq); i {
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
		file_api_grpc_captcha_captcha_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRsp); i {
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
			RawDescriptor: file_api_grpc_captcha_captcha_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_captcha_captcha_proto_goTypes,
		DependencyIndexes: file_api_grpc_captcha_captcha_proto_depIdxs,
		EnumInfos:         file_api_grpc_captcha_captcha_proto_enumTypes,
		MessageInfos:      file_api_grpc_captcha_captcha_proto_msgTypes,
	}.Build()
	File_api_grpc_captcha_captcha_proto = out.File
	file_api_grpc_captcha_captcha_proto_rawDesc = nil
	file_api_grpc_captcha_captcha_proto_goTypes = nil
	file_api_grpc_captcha_captcha_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CaptchaServiceClient is the client API for CaptchaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CaptchaServiceClient interface {
	// 异步验证码 用于短信/邮箱等
	AsyncCode(ctx context.Context, in *AsyncCodeReq, opts ...grpc.CallOption) (*AsyncCodeRsp, error)
	// 图形验证码
	GraphCaptcha(ctx context.Context, in *GraphCaptchaReq, opts ...grpc.CallOption) (*GraphCaptchaRsp, error)
	// 验证码校验
	Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRsp, error)
}

type captchaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptchaServiceClient(cc grpc.ClientConnInterface) CaptchaServiceClient {
	return &captchaServiceClient{cc}
}

func (c *captchaServiceClient) AsyncCode(ctx context.Context, in *AsyncCodeReq, opts ...grpc.CallOption) (*AsyncCodeRsp, error) {
	out := new(AsyncCodeRsp)
	err := c.cc.Invoke(ctx, "/captcha.CaptchaService/AsyncCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) GraphCaptcha(ctx context.Context, in *GraphCaptchaReq, opts ...grpc.CallOption) (*GraphCaptchaRsp, error) {
	out := new(GraphCaptchaRsp)
	err := c.cc.Invoke(ctx, "/captcha.CaptchaService/GraphCaptcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRsp, error) {
	out := new(VerifyRsp)
	err := c.cc.Invoke(ctx, "/captcha.CaptchaService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptchaServiceServer is the server API for CaptchaService service.
type CaptchaServiceServer interface {
	// 异步验证码 用于短信/邮箱等
	AsyncCode(context.Context, *AsyncCodeReq) (*AsyncCodeRsp, error)
	// 图形验证码
	GraphCaptcha(context.Context, *GraphCaptchaReq) (*GraphCaptchaRsp, error)
	// 验证码校验
	Verify(context.Context, *VerifyReq) (*VerifyRsp, error)
}

// UnimplementedCaptchaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCaptchaServiceServer struct {
}

func (*UnimplementedCaptchaServiceServer) AsyncCode(context.Context, *AsyncCodeReq) (*AsyncCodeRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsyncCode not implemented")
}
func (*UnimplementedCaptchaServiceServer) GraphCaptcha(context.Context, *GraphCaptchaReq) (*GraphCaptchaRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GraphCaptcha not implemented")
}
func (*UnimplementedCaptchaServiceServer) Verify(context.Context, *VerifyReq) (*VerifyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterCaptchaServiceServer(s *grpc.Server, srv CaptchaServiceServer) {
	s.RegisterService(&_CaptchaService_serviceDesc, srv)
}

func _CaptchaService_AsyncCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AsyncCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).AsyncCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/captcha.CaptchaService/AsyncCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).AsyncCode(ctx, req.(*AsyncCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_GraphCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphCaptchaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).GraphCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/captcha.CaptchaService/GraphCaptcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).GraphCaptcha(ctx, req.(*GraphCaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/captcha.CaptchaService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).Verify(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CaptchaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "captcha.CaptchaService",
	HandlerType: (*CaptchaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AsyncCode",
			Handler:    _CaptchaService_AsyncCode_Handler,
		},
		{
			MethodName: "GraphCaptcha",
			Handler:    _CaptchaService_GraphCaptcha_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _CaptchaService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/captcha/captcha.proto",
}

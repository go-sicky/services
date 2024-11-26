//
// The MIT License (MIT)
//
// Copyright (c) 2021 HereweTech Co.LTD
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: proto/setting.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitDBResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InitDBResp) Reset() {
	*x = InitDBResp{}
	mi := &file_proto_setting_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitDBResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDBResp) ProtoMessage() {}

func (x *InitDBResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDBResp.ProtoReflect.Descriptor instead.
func (*InitDBResp) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{0}
}

func (x *InitDBResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type SettingDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value  string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Raw    bool   `protobuf:"varint,4,opt,name=raw,proto3" json:"raw,omitempty"`
	Status int64  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SettingDef) Reset() {
	*x = SettingDef{}
	mi := &file_proto_setting_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SettingDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingDef) ProtoMessage() {}

func (x *SettingDef) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingDef.ProtoReflect.Descriptor instead.
func (*SettingDef) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{1}
}

func (x *SettingDef) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SettingDef) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SettingDef) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SettingDef) GetRaw() bool {
	if x != nil {
		return x.Raw
	}
	return false
}

func (x *SettingDef) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *SettingDef `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"`
}

func (x *SetReq) Reset() {
	*x = SetReq{}
	mi := &file_proto_setting_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReq) ProtoMessage() {}

func (x *SetReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReq.ProtoReflect.Descriptor instead.
func (*SetReq) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{2}
}

func (x *SetReq) GetSetting() *SettingDef {
	if x != nil {
		return x.Setting
	}
	return nil
}

type SetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,127,opt,name=id,proto3" json:"id,omitempty"`
	IsNew bool   `protobuf:"varint,128,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
}

func (x *SetResp) Reset() {
	*x = SetResp{}
	mi := &file_proto_setting_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResp) ProtoMessage() {}

func (x *SetResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResp.ProtoReflect.Descriptor instead.
func (*SetResp) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{3}
}

func (x *SetResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetResp) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Id  string `protobuf:"bytes,127,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	mi := &file_proto_setting_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{4}
}

func (x *GetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *SettingDef `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	mi := &file_proto_setting_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{5}
}

func (x *GetResp) GetSetting() *SettingDef {
	if x != nil {
		return x.Setting
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Id  string `protobuf:"bytes,127,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	mi := &file_proto_setting_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DeleteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteResp) Reset() {
	*x = DeleteResp{}
	mi := &file_proto_setting_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResp) ProtoMessage() {}

func (x *DeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_setting_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResp.ProtoReflect.Descriptor instead.
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return file_proto_setting_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResp) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

var File_proto_setting_proto protoreflect.FileDescriptor

var file_proto_setting_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x76, 0x63, 0x2e, 0x73, 0x69, 0x63, 0x6b, 0x79,
	0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6e, 0x0a, 0x0a, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x72, 0x61, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x41, 0x0a, 0x06, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x73, 0x69, 0x63,
	0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x66, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x31,
	0x0a, 0x07, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x6e, 0x65, 0x77, 0x18, 0x80, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4e, 0x65,
	0x77, 0x22, 0x2a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x73, 0x69, 0x63, 0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x66, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x2d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x26, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0x8d, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x06, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x73, 0x69, 0x63,
	0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x44,
	0x42, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x73, 0x69, 0x63, 0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x73, 0x69,
	0x63, 0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x73, 0x69, 0x63, 0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x73, 0x69, 0x63, 0x6b,
	0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x45, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x73, 0x69, 0x63, 0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x73, 0x69, 0x63, 0x6b, 0x79, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_setting_proto_rawDescOnce sync.Once
	file_proto_setting_proto_rawDescData = file_proto_setting_proto_rawDesc
)

func file_proto_setting_proto_rawDescGZIP() []byte {
	file_proto_setting_proto_rawDescOnce.Do(func() {
		file_proto_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_setting_proto_rawDescData)
	})
	return file_proto_setting_proto_rawDescData
}

var file_proto_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_setting_proto_goTypes = []any{
	(*InitDBResp)(nil),    // 0: svc.sicky.setting.InitDBResp
	(*SettingDef)(nil),    // 1: svc.sicky.setting.SettingDef
	(*SetReq)(nil),        // 2: svc.sicky.setting.SetReq
	(*SetResp)(nil),       // 3: svc.sicky.setting.SetResp
	(*GetReq)(nil),        // 4: svc.sicky.setting.GetReq
	(*GetResp)(nil),       // 5: svc.sicky.setting.GetResp
	(*DeleteReq)(nil),     // 6: svc.sicky.setting.DeleteReq
	(*DeleteResp)(nil),    // 7: svc.sicky.setting.DeleteResp
	(*emptypb.Empty)(nil), // 8: google.protobuf.Empty
}
var file_proto_setting_proto_depIdxs = []int32{
	1, // 0: svc.sicky.setting.SetReq.setting:type_name -> svc.sicky.setting.SettingDef
	1, // 1: svc.sicky.setting.GetResp.setting:type_name -> svc.sicky.setting.SettingDef
	8, // 2: svc.sicky.setting.Setting.InitDB:input_type -> google.protobuf.Empty
	2, // 3: svc.sicky.setting.Setting.Set:input_type -> svc.sicky.setting.SetReq
	4, // 4: svc.sicky.setting.Setting.Get:input_type -> svc.sicky.setting.GetReq
	6, // 5: svc.sicky.setting.Setting.Delete:input_type -> svc.sicky.setting.DeleteReq
	0, // 6: svc.sicky.setting.Setting.InitDB:output_type -> svc.sicky.setting.InitDBResp
	3, // 7: svc.sicky.setting.Setting.Set:output_type -> svc.sicky.setting.SetResp
	5, // 8: svc.sicky.setting.Setting.Get:output_type -> svc.sicky.setting.GetResp
	7, // 9: svc.sicky.setting.Setting.Delete:output_type -> svc.sicky.setting.DeleteResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_setting_proto_init() }
func file_proto_setting_proto_init() {
	if File_proto_setting_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_setting_proto_goTypes,
		DependencyIndexes: file_proto_setting_proto_depIdxs,
		MessageInfos:      file_proto_setting_proto_msgTypes,
	}.Build()
	File_proto_setting_proto = out.File
	file_proto_setting_proto_rawDesc = nil
	file_proto_setting_proto_goTypes = nil
	file_proto_setting_proto_depIdxs = nil
}

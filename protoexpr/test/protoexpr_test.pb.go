// Copyright 2022 The Go Protobuf Expression Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protoexpr/protoexpr_test.proto

package test

import (
	_ "github.com/kagadar/go_proto_expression/genproto/options"
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

type TestFiltering_TestEnum int32

const (
	TestFiltering_VALUE_0 TestFiltering_TestEnum = 0
	TestFiltering_VALUE_1 TestFiltering_TestEnum = 1
)

// Enum value maps for TestFiltering_TestEnum.
var (
	TestFiltering_TestEnum_name = map[int32]string{
		0: "VALUE_0",
		1: "VALUE_1",
	}
	TestFiltering_TestEnum_value = map[string]int32{
		"VALUE_0": 0,
		"VALUE_1": 1,
	}
)

func (x TestFiltering_TestEnum) Enum() *TestFiltering_TestEnum {
	p := new(TestFiltering_TestEnum)
	*p = x
	return p
}

func (x TestFiltering_TestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestFiltering_TestEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_protoexpr_protoexpr_test_proto_enumTypes[0].Descriptor()
}

func (TestFiltering_TestEnum) Type() protoreflect.EnumType {
	return &file_protoexpr_protoexpr_test_proto_enumTypes[0]
}

func (x TestFiltering_TestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestFiltering_TestEnum.Descriptor instead.
func (TestFiltering_TestEnum) EnumDescriptor() ([]byte, []int) {
	return file_protoexpr_protoexpr_test_proto_rawDescGZIP(), []int{2, 0}
}

type ListTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter    string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListTestRequest) Reset() {
	*x = ListTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoexpr_protoexpr_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestRequest) ProtoMessage() {}

func (x *ListTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoexpr_protoexpr_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestRequest.ProtoReflect.Descriptor instead.
func (*ListTestRequest) Descriptor() ([]byte, []int) {
	return file_protoexpr_protoexpr_test_proto_rawDescGZIP(), []int{0}
}

func (x *ListTestRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListTestRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTestRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTestRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tests         []*TestFiltering `protobuf:"bytes,2,rep,name=tests,proto3" json:"tests,omitempty"`
	NextPageToken string           `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTestResponse) Reset() {
	*x = ListTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoexpr_protoexpr_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestResponse) ProtoMessage() {}

func (x *ListTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoexpr_protoexpr_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestResponse.ProtoReflect.Descriptor instead.
func (*ListTestResponse) Descriptor() ([]byte, []int) {
	return file_protoexpr_protoexpr_test_proto_rawDescGZIP(), []int{1}
}

func (x *ListTestResponse) GetTests() []*TestFiltering {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *ListTestResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type TestFiltering struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterableSubmessage   *TestFiltering_SubMessage `protobuf:"bytes,1,opt,name=filterable_submessage,json=filterableSubmessage,proto3" json:"filterable_submessage,omitempty"`
	UnfilterableSubmessage *TestFiltering_SubMessage `protobuf:"bytes,2,opt,name=unfilterable_submessage,json=unfilterableSubmessage,proto3" json:"unfilterable_submessage,omitempty"`
	DefaultSubmessage      *TestFiltering_SubMessage `protobuf:"bytes,3,opt,name=default_submessage,json=defaultSubmessage,proto3" json:"default_submessage,omitempty"`
	FilterablePrimitive    string                    `protobuf:"bytes,4,opt,name=filterable_primitive,json=filterablePrimitive,proto3" json:"filterable_primitive,omitempty"`
	UnfilterablePrimitive  int64                     `protobuf:"varint,5,opt,name=unfilterable_primitive,json=unfilterablePrimitive,proto3" json:"unfilterable_primitive,omitempty"`
	DefaultFloat           float32                   `protobuf:"fixed32,6,opt,name=default_float,json=defaultFloat,proto3" json:"default_float,omitempty"`
	DefaultBool            bool                      `protobuf:"varint,7,opt,name=default_bool,json=defaultBool,proto3" json:"default_bool,omitempty"`
	DefaultEnum            TestFiltering_TestEnum    `protobuf:"varint,8,opt,name=default_enum,json=defaultEnum,proto3,enum=kagadar.protoexpr.options.TestFiltering_TestEnum" json:"default_enum,omitempty"`
}

func (x *TestFiltering) Reset() {
	*x = TestFiltering{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoexpr_protoexpr_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestFiltering) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestFiltering) ProtoMessage() {}

func (x *TestFiltering) ProtoReflect() protoreflect.Message {
	mi := &file_protoexpr_protoexpr_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestFiltering.ProtoReflect.Descriptor instead.
func (*TestFiltering) Descriptor() ([]byte, []int) {
	return file_protoexpr_protoexpr_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestFiltering) GetFilterableSubmessage() *TestFiltering_SubMessage {
	if x != nil {
		return x.FilterableSubmessage
	}
	return nil
}

func (x *TestFiltering) GetUnfilterableSubmessage() *TestFiltering_SubMessage {
	if x != nil {
		return x.UnfilterableSubmessage
	}
	return nil
}

func (x *TestFiltering) GetDefaultSubmessage() *TestFiltering_SubMessage {
	if x != nil {
		return x.DefaultSubmessage
	}
	return nil
}

func (x *TestFiltering) GetFilterablePrimitive() string {
	if x != nil {
		return x.FilterablePrimitive
	}
	return ""
}

func (x *TestFiltering) GetUnfilterablePrimitive() int64 {
	if x != nil {
		return x.UnfilterablePrimitive
	}
	return 0
}

func (x *TestFiltering) GetDefaultFloat() float32 {
	if x != nil {
		return x.DefaultFloat
	}
	return 0
}

func (x *TestFiltering) GetDefaultBool() bool {
	if x != nil {
		return x.DefaultBool
	}
	return false
}

func (x *TestFiltering) GetDefaultEnum() TestFiltering_TestEnum {
	if x != nil {
		return x.DefaultEnum
	}
	return TestFiltering_VALUE_0
}

type TestFiltering_SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterablePrimitive   int64  `protobuf:"varint,1,opt,name=filterable_primitive,json=filterablePrimitive,proto3" json:"filterable_primitive,omitempty"`
	UnfilterablePrimitive string `protobuf:"bytes,2,opt,name=unfilterable_primitive,json=unfilterablePrimitive,proto3" json:"unfilterable_primitive,omitempty"`
}

func (x *TestFiltering_SubMessage) Reset() {
	*x = TestFiltering_SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoexpr_protoexpr_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestFiltering_SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestFiltering_SubMessage) ProtoMessage() {}

func (x *TestFiltering_SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protoexpr_protoexpr_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestFiltering_SubMessage.ProtoReflect.Descriptor instead.
func (*TestFiltering_SubMessage) Descriptor() ([]byte, []int) {
	return file_protoexpr_protoexpr_test_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TestFiltering_SubMessage) GetFilterablePrimitive() int64 {
	if x != nil {
		return x.FilterablePrimitive
	}
	return 0
}

func (x *TestFiltering_SubMessage) GetUnfilterablePrimitive() string {
	if x != nil {
		return x.UnfilterablePrimitive
	}
	return ""
}

var File_protoexpr_protoexpr_test_proto protoreflect.FileDescriptor

var file_protoexpr_protoexpr_test_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x65, 0x78, 0x70, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6b, 0x61,
	0x67, 0x61, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x70, 0x72, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x07, 0xea, 0x9f, 0x91, 0x4d, 0x02, 0x08, 0x02, 0x22, 0xa8, 0x06,
	0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x71, 0x0a, 0x15, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x75,
	0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78,
	0x70, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x07, 0xea, 0x9f, 0x91, 0x4d, 0x02, 0x08, 0x01, 0x52, 0x14, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x75, 0x0a, 0x17, 0x75, 0x6e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75,
	0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x07, 0xea, 0x9f, 0x91, 0x4d, 0x02, 0x08,
	0x00, 0x52, 0x16, 0x75, 0x6e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x62, 0x0a, 0x12, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a,
	0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xea, 0x9f, 0x91,
	0x4d, 0x02, 0x08, 0x01, 0x52, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x3e, 0x0a, 0x16, 0x75, 0x6e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xea, 0x9f, 0x91, 0x4d, 0x02,
	0x08, 0x00, 0x52, 0x15, 0x75, 0x6e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x6f, 0x6f,
	0x6c, 0x12, 0x54, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x1a, 0x88, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xea, 0x9f, 0x91, 0x4d, 0x02, 0x08, 0x01, 0x52, 0x13, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x3e, 0x0a, 0x16, 0x75, 0x6e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xea, 0x9f, 0x91, 0x4d, 0x02, 0x08, 0x00, 0x52, 0x15, 0x75, 0x6e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b,
	0x0a, 0x07, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x30, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x31, 0x10, 0x01, 0x32, 0x7f, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0b, 0xea, 0x9f,
	0x91, 0x4d, 0x06, 0x08, 0xe8, 0x07, 0x10, 0x90, 0x4e, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x67, 0x61, 0x64, 0x61, 0x72, 0x2f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoexpr_protoexpr_test_proto_rawDescOnce sync.Once
	file_protoexpr_protoexpr_test_proto_rawDescData = file_protoexpr_protoexpr_test_proto_rawDesc
)

func file_protoexpr_protoexpr_test_proto_rawDescGZIP() []byte {
	file_protoexpr_protoexpr_test_proto_rawDescOnce.Do(func() {
		file_protoexpr_protoexpr_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoexpr_protoexpr_test_proto_rawDescData)
	})
	return file_protoexpr_protoexpr_test_proto_rawDescData
}

var file_protoexpr_protoexpr_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protoexpr_protoexpr_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protoexpr_protoexpr_test_proto_goTypes = []interface{}{
	(TestFiltering_TestEnum)(0),      // 0: kagadar.protoexpr.options.TestFiltering.TestEnum
	(*ListTestRequest)(nil),          // 1: kagadar.protoexpr.options.ListTestRequest
	(*ListTestResponse)(nil),         // 2: kagadar.protoexpr.options.ListTestResponse
	(*TestFiltering)(nil),            // 3: kagadar.protoexpr.options.TestFiltering
	(*TestFiltering_SubMessage)(nil), // 4: kagadar.protoexpr.options.TestFiltering.SubMessage
}
var file_protoexpr_protoexpr_test_proto_depIdxs = []int32{
	3, // 0: kagadar.protoexpr.options.ListTestResponse.tests:type_name -> kagadar.protoexpr.options.TestFiltering
	4, // 1: kagadar.protoexpr.options.TestFiltering.filterable_submessage:type_name -> kagadar.protoexpr.options.TestFiltering.SubMessage
	4, // 2: kagadar.protoexpr.options.TestFiltering.unfilterable_submessage:type_name -> kagadar.protoexpr.options.TestFiltering.SubMessage
	4, // 3: kagadar.protoexpr.options.TestFiltering.default_submessage:type_name -> kagadar.protoexpr.options.TestFiltering.SubMessage
	0, // 4: kagadar.protoexpr.options.TestFiltering.default_enum:type_name -> kagadar.protoexpr.options.TestFiltering.TestEnum
	1, // 5: kagadar.protoexpr.options.TestService.ListTest:input_type -> kagadar.protoexpr.options.ListTestRequest
	2, // 6: kagadar.protoexpr.options.TestService.ListTest:output_type -> kagadar.protoexpr.options.ListTestResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protoexpr_protoexpr_test_proto_init() }
func file_protoexpr_protoexpr_test_proto_init() {
	if File_protoexpr_protoexpr_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoexpr_protoexpr_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestRequest); i {
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
		file_protoexpr_protoexpr_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestResponse); i {
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
		file_protoexpr_protoexpr_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestFiltering); i {
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
		file_protoexpr_protoexpr_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestFiltering_SubMessage); i {
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
			RawDescriptor: file_protoexpr_protoexpr_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoexpr_protoexpr_test_proto_goTypes,
		DependencyIndexes: file_protoexpr_protoexpr_test_proto_depIdxs,
		EnumInfos:         file_protoexpr_protoexpr_test_proto_enumTypes,
		MessageInfos:      file_protoexpr_protoexpr_test_proto_msgTypes,
	}.Build()
	File_protoexpr_protoexpr_test_proto = out.File
	file_protoexpr_protoexpr_test_proto_rawDesc = nil
	file_protoexpr_protoexpr_test_proto_goTypes = nil
	file_protoexpr_protoexpr_test_proto_depIdxs = nil
}

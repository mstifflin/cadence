// Copyright (c) 2020 Uber Technologies, Inc.
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
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: uber/cadence/api/v1/visibility.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IndexedValueType int32

const (
	IndexedValueType_INDEXED_VALUE_TYPE_INVALID  IndexedValueType = 0
	IndexedValueType_INDEXED_VALUE_TYPE_STRING   IndexedValueType = 1
	IndexedValueType_INDEXED_VALUE_TYPE_KEYWORD  IndexedValueType = 2
	IndexedValueType_INDEXED_VALUE_TYPE_INT      IndexedValueType = 3
	IndexedValueType_INDEXED_VALUE_TYPE_DOUBLE   IndexedValueType = 4
	IndexedValueType_INDEXED_VALUE_TYPE_BOOL     IndexedValueType = 5
	IndexedValueType_INDEXED_VALUE_TYPE_DATETIME IndexedValueType = 6
)

// Enum value maps for IndexedValueType.
var (
	IndexedValueType_name = map[int32]string{
		0: "INDEXED_VALUE_TYPE_INVALID",
		1: "INDEXED_VALUE_TYPE_STRING",
		2: "INDEXED_VALUE_TYPE_KEYWORD",
		3: "INDEXED_VALUE_TYPE_INT",
		4: "INDEXED_VALUE_TYPE_DOUBLE",
		5: "INDEXED_VALUE_TYPE_BOOL",
		6: "INDEXED_VALUE_TYPE_DATETIME",
	}
	IndexedValueType_value = map[string]int32{
		"INDEXED_VALUE_TYPE_INVALID":  0,
		"INDEXED_VALUE_TYPE_STRING":   1,
		"INDEXED_VALUE_TYPE_KEYWORD":  2,
		"INDEXED_VALUE_TYPE_INT":      3,
		"INDEXED_VALUE_TYPE_DOUBLE":   4,
		"INDEXED_VALUE_TYPE_BOOL":     5,
		"INDEXED_VALUE_TYPE_DATETIME": 6,
	}
)

func (x IndexedValueType) Enum() *IndexedValueType {
	p := new(IndexedValueType)
	*p = x
	return p
}

func (x IndexedValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexedValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_uber_cadence_api_v1_visibility_proto_enumTypes[0].Descriptor()
}

func (IndexedValueType) Type() protoreflect.EnumType {
	return &file_uber_cadence_api_v1_visibility_proto_enumTypes[0]
}

func (x IndexedValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexedValueType.Descriptor instead.
func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_visibility_proto_rawDescGZIP(), []int{0}
}

type WorkflowExecutionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId      string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *WorkflowExecutionFilter) Reset() {
	*x = WorkflowExecutionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowExecutionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecutionFilter) ProtoMessage() {}

func (x *WorkflowExecutionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecutionFilter.ProtoReflect.Descriptor instead.
func (*WorkflowExecutionFilter) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_visibility_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowExecutionFilter) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowExecutionFilter) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type WorkflowTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WorkflowTypeFilter) Reset() {
	*x = WorkflowTypeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowTypeFilter) ProtoMessage() {}

func (x *WorkflowTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowTypeFilter.ProtoReflect.Descriptor instead.
func (*WorkflowTypeFilter) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_visibility_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowTypeFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StartTimeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EarliestTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=earliest_time,json=earliestTime,proto3" json:"earliest_time,omitempty"`
	LatestTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
}

func (x *StartTimeFilter) Reset() {
	*x = StartTimeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTimeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTimeFilter) ProtoMessage() {}

func (x *StartTimeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTimeFilter.ProtoReflect.Descriptor instead.
func (*StartTimeFilter) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_visibility_proto_rawDescGZIP(), []int{2}
}

func (x *StartTimeFilter) GetEarliestTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EarliestTime
	}
	return nil
}

func (x *StartTimeFilter) GetLatestTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LatestTime
	}
	return nil
}

type StatusFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status WorkflowExecutionCloseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=uber.cadence.api.v1.WorkflowExecutionCloseStatus" json:"status,omitempty"`
}

func (x *StatusFilter) Reset() {
	*x = StatusFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusFilter) ProtoMessage() {}

func (x *StatusFilter) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_visibility_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusFilter.ProtoReflect.Descriptor instead.
func (*StatusFilter) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_visibility_proto_rawDescGZIP(), []int{3}
}

func (x *StatusFilter) GetStatus() WorkflowExecutionCloseStatus {
	if x != nil {
		return x.Status
	}
	return WorkflowExecutionCloseStatus_WORKFLOW_EXECUTION_CLOSE_STATUS_INVALID
}

var File_uber_cadence_api_v1_visibility_proto protoreflect.FileDescriptor

var file_uber_cadence_api_v1_visibility_proto_rawDesc = []byte{
	0x0a, 0x24, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x75, 0x62,
	0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x51, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75,
	0x6e, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54,
	0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8f, 0x01,
	0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x3f, 0x0a, 0x0d, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x59, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x49, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x31, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xea, 0x01, 0x0a, 0x10, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1e,
	0x0a, 0x1a, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x1a,
	0x0a, 0x16, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e,
	0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x44,
	0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45,
	0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54,
	0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x06, 0x42, 0x56, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x75,
	0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x62, 0x65, 0x72, 0x2f,
	0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x2e, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uber_cadence_api_v1_visibility_proto_rawDescOnce sync.Once
	file_uber_cadence_api_v1_visibility_proto_rawDescData = file_uber_cadence_api_v1_visibility_proto_rawDesc
)

func file_uber_cadence_api_v1_visibility_proto_rawDescGZIP() []byte {
	file_uber_cadence_api_v1_visibility_proto_rawDescOnce.Do(func() {
		file_uber_cadence_api_v1_visibility_proto_rawDescData = protoimpl.X.CompressGZIP(file_uber_cadence_api_v1_visibility_proto_rawDescData)
	})
	return file_uber_cadence_api_v1_visibility_proto_rawDescData
}

var file_uber_cadence_api_v1_visibility_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_uber_cadence_api_v1_visibility_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_uber_cadence_api_v1_visibility_proto_goTypes = []interface{}{
	(IndexedValueType)(0),             // 0: uber.cadence.api.v1.IndexedValueType
	(*WorkflowExecutionFilter)(nil),   // 1: uber.cadence.api.v1.WorkflowExecutionFilter
	(*WorkflowTypeFilter)(nil),        // 2: uber.cadence.api.v1.WorkflowTypeFilter
	(*StartTimeFilter)(nil),           // 3: uber.cadence.api.v1.StartTimeFilter
	(*StatusFilter)(nil),              // 4: uber.cadence.api.v1.StatusFilter
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
	(WorkflowExecutionCloseStatus)(0), // 6: uber.cadence.api.v1.WorkflowExecutionCloseStatus
}
var file_uber_cadence_api_v1_visibility_proto_depIdxs = []int32{
	5, // 0: uber.cadence.api.v1.StartTimeFilter.earliest_time:type_name -> google.protobuf.Timestamp
	5, // 1: uber.cadence.api.v1.StartTimeFilter.latest_time:type_name -> google.protobuf.Timestamp
	6, // 2: uber.cadence.api.v1.StatusFilter.status:type_name -> uber.cadence.api.v1.WorkflowExecutionCloseStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_uber_cadence_api_v1_visibility_proto_init() }
func file_uber_cadence_api_v1_visibility_proto_init() {
	if File_uber_cadence_api_v1_visibility_proto != nil {
		return
	}
	file_uber_cadence_api_v1_workflow_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_uber_cadence_api_v1_visibility_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowExecutionFilter); i {
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
		file_uber_cadence_api_v1_visibility_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowTypeFilter); i {
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
		file_uber_cadence_api_v1_visibility_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTimeFilter); i {
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
		file_uber_cadence_api_v1_visibility_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusFilter); i {
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
			RawDescriptor: file_uber_cadence_api_v1_visibility_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uber_cadence_api_v1_visibility_proto_goTypes,
		DependencyIndexes: file_uber_cadence_api_v1_visibility_proto_depIdxs,
		EnumInfos:         file_uber_cadence_api_v1_visibility_proto_enumTypes,
		MessageInfos:      file_uber_cadence_api_v1_visibility_proto_msgTypes,
	}.Build()
	File_uber_cadence_api_v1_visibility_proto = out.File
	file_uber_cadence_api_v1_visibility_proto_rawDesc = nil
	file_uber_cadence_api_v1_visibility_proto_goTypes = nil
	file_uber_cadence_api_v1_visibility_proto_depIdxs = nil
}
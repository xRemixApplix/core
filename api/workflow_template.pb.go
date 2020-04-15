// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.3
// source: workflow_template.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateWorkflowTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace        string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplate *WorkflowTemplate `protobuf:"bytes,2,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
}

func (x *CreateWorkflowTemplateRequest) Reset() {
	*x = CreateWorkflowTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkflowTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkflowTemplateRequest) ProtoMessage() {}

func (x *CreateWorkflowTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkflowTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWorkflowTemplateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateWorkflowTemplateRequest) GetWorkflowTemplate() *WorkflowTemplate {
	if x != nil {
		return x.WorkflowTemplate
	}
	return nil
}

type UpdateWorkflowTemplateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace        string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplate *WorkflowTemplate `protobuf:"bytes,2,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
}

func (x *UpdateWorkflowTemplateVersionRequest) Reset() {
	*x = UpdateWorkflowTemplateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkflowTemplateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkflowTemplateVersionRequest) ProtoMessage() {}

func (x *UpdateWorkflowTemplateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkflowTemplateVersionRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkflowTemplateVersionRequest) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateWorkflowTemplateVersionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *UpdateWorkflowTemplateVersionRequest) GetWorkflowTemplate() *WorkflowTemplate {
	if x != nil {
		return x.WorkflowTemplate
	}
	return nil
}

type GetWorkflowTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid       string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Version   int32  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetWorkflowTemplateRequest) Reset() {
	*x = GetWorkflowTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkflowTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowTemplateRequest) ProtoMessage() {}

func (x *GetWorkflowTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{2}
}

func (x *GetWorkflowTemplateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetWorkflowTemplateRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GetWorkflowTemplateRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type CloneWorkflowTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid       string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version   int32  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CloneWorkflowTemplateRequest) Reset() {
	*x = CloneWorkflowTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloneWorkflowTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloneWorkflowTemplateRequest) ProtoMessage() {}

func (x *CloneWorkflowTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloneWorkflowTemplateRequest.ProtoReflect.Descriptor instead.
func (*CloneWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{3}
}

func (x *CloneWorkflowTemplateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CloneWorkflowTemplateRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CloneWorkflowTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CloneWorkflowTemplateRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type ListWorkflowTemplateVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid       string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ListWorkflowTemplateVersionsRequest) Reset() {
	*x = ListWorkflowTemplateVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkflowTemplateVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkflowTemplateVersionsRequest) ProtoMessage() {}

func (x *ListWorkflowTemplateVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkflowTemplateVersionsRequest.ProtoReflect.Descriptor instead.
func (*ListWorkflowTemplateVersionsRequest) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{4}
}

func (x *ListWorkflowTemplateVersionsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListWorkflowTemplateVersionsRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ListWorkflowTemplateVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count             int32               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WorkflowTemplates []*WorkflowTemplate `protobuf:"bytes,2,rep,name=workflowTemplates,proto3" json:"workflowTemplates,omitempty"`
}

func (x *ListWorkflowTemplateVersionsResponse) Reset() {
	*x = ListWorkflowTemplateVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkflowTemplateVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkflowTemplateVersionsResponse) ProtoMessage() {}

func (x *ListWorkflowTemplateVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkflowTemplateVersionsResponse.ProtoReflect.Descriptor instead.
func (*ListWorkflowTemplateVersionsResponse) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{5}
}

func (x *ListWorkflowTemplateVersionsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListWorkflowTemplateVersionsResponse) GetWorkflowTemplates() []*WorkflowTemplate {
	if x != nil {
		return x.WorkflowTemplates
	}
	return nil
}

type ListWorkflowTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ListWorkflowTemplatesRequest) Reset() {
	*x = ListWorkflowTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkflowTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkflowTemplatesRequest) ProtoMessage() {}

func (x *ListWorkflowTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkflowTemplatesRequest.ProtoReflect.Descriptor instead.
func (*ListWorkflowTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{6}
}

func (x *ListWorkflowTemplatesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ListWorkflowTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count             int32               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WorkflowTemplates []*WorkflowTemplate `protobuf:"bytes,2,rep,name=workflowTemplates,proto3" json:"workflowTemplates,omitempty"`
}

func (x *ListWorkflowTemplatesResponse) Reset() {
	*x = ListWorkflowTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkflowTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkflowTemplatesResponse) ProtoMessage() {}

func (x *ListWorkflowTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkflowTemplatesResponse.ProtoReflect.Descriptor instead.
func (*ListWorkflowTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{7}
}

func (x *ListWorkflowTemplatesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListWorkflowTemplatesResponse) GetWorkflowTemplates() []*WorkflowTemplate {
	if x != nil {
		return x.WorkflowTemplates
	}
	return nil
}

type ArchiveWorkflowTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid       string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ArchiveWorkflowTemplateRequest) Reset() {
	*x = ArchiveWorkflowTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveWorkflowTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveWorkflowTemplateRequest) ProtoMessage() {}

func (x *ArchiveWorkflowTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveWorkflowTemplateRequest.ProtoReflect.Descriptor instead.
func (*ArchiveWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{8}
}

func (x *ArchiveWorkflowTemplateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ArchiveWorkflowTemplateRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ArchiveWorkflowTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowTemplate *WorkflowTemplate `protobuf:"bytes,1,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
}

func (x *ArchiveWorkflowTemplateResponse) Reset() {
	*x = ArchiveWorkflowTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveWorkflowTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveWorkflowTemplateResponse) ProtoMessage() {}

func (x *ArchiveWorkflowTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveWorkflowTemplateResponse.ProtoReflect.Descriptor instead.
func (*ArchiveWorkflowTemplateResponse) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{9}
}

func (x *ArchiveWorkflowTemplateResponse) GetWorkflowTemplate() *WorkflowTemplate {
	if x != nil {
		return x.WorkflowTemplate
	}
	return nil
}

type WorkflowExecutionStatisticReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total        int32  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	LastExecuted string `protobuf:"bytes,2,opt,name=lastExecuted,proto3" json:"lastExecuted,omitempty"`
	Running      int32  `protobuf:"varint,3,opt,name=running,proto3" json:"running,omitempty"`
	Completed    int32  `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	Failed       int32  `protobuf:"varint,5,opt,name=failed,proto3" json:"failed,omitempty"`
}

func (x *WorkflowExecutionStatisticReport) Reset() {
	*x = WorkflowExecutionStatisticReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowExecutionStatisticReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecutionStatisticReport) ProtoMessage() {}

func (x *WorkflowExecutionStatisticReport) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecutionStatisticReport.ProtoReflect.Descriptor instead.
func (*WorkflowExecutionStatisticReport) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{10}
}

func (x *WorkflowExecutionStatisticReport) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *WorkflowExecutionStatisticReport) GetLastExecuted() string {
	if x != nil {
		return x.LastExecuted
	}
	return ""
}

func (x *WorkflowExecutionStatisticReport) GetRunning() int32 {
	if x != nil {
		return x.Running
	}
	return 0
}

func (x *WorkflowExecutionStatisticReport) GetCompleted() int32 {
	if x != nil {
		return x.Completed
	}
	return 0
}

func (x *WorkflowExecutionStatisticReport) GetFailed() int32 {
	if x != nil {
		return x.Failed
	}
	return 0
}

type WorkflowTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt  string                            `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Uid        string                            `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name       string                            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version    int32                             `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Manifest   string                            `protobuf:"bytes,5,opt,name=manifest,proto3" json:"manifest,omitempty"`
	IsLatest   bool                              `protobuf:"varint,6,opt,name=isLatest,proto3" json:"isLatest,omitempty"`
	IsArchived bool                              `protobuf:"varint,7,opt,name=isArchived,proto3" json:"isArchived,omitempty"`
	Stats      *WorkflowExecutionStatisticReport `protobuf:"bytes,8,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *WorkflowTemplate) Reset() {
	*x = WorkflowTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_template_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowTemplate) ProtoMessage() {}

func (x *WorkflowTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_template_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowTemplate.ProtoReflect.Descriptor instead.
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return file_workflow_template_proto_rawDescGZIP(), []int{11}
}

func (x *WorkflowTemplate) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WorkflowTemplate) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *WorkflowTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowTemplate) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WorkflowTemplate) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

func (x *WorkflowTemplate) GetIsLatest() bool {
	if x != nil {
		return x.IsLatest
	}
	return false
}

func (x *WorkflowTemplate) GetIsArchived() bool {
	if x != nil {
		return x.IsArchived
	}
	return false
}

func (x *WorkflowTemplate) GetStats() *WorkflowExecutionStatisticReport {
	if x != nil {
		return x.Stats
	}
	return nil
}

var File_workflow_template_proto protoreflect.FileDescriptor

var file_workflow_template_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x80,
	0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x22, 0x87, 0x01, 0x0a, 0x24, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x66, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x1c, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x55, 0x0a, 0x23, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x24, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x1c,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x7a, 0x0a, 0x1d, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x43, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x1e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x1f, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x10, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0xac,
	0x01, 0x0a, 0x20, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x85, 0x02,
	0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_template_proto_rawDescOnce sync.Once
	file_workflow_template_proto_rawDescData = file_workflow_template_proto_rawDesc
)

func file_workflow_template_proto_rawDescGZIP() []byte {
	file_workflow_template_proto_rawDescOnce.Do(func() {
		file_workflow_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_template_proto_rawDescData)
	})
	return file_workflow_template_proto_rawDescData
}

var file_workflow_template_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_workflow_template_proto_goTypes = []interface{}{
	(*CreateWorkflowTemplateRequest)(nil),        // 0: api.CreateWorkflowTemplateRequest
	(*UpdateWorkflowTemplateVersionRequest)(nil), // 1: api.UpdateWorkflowTemplateVersionRequest
	(*GetWorkflowTemplateRequest)(nil),           // 2: api.GetWorkflowTemplateRequest
	(*CloneWorkflowTemplateRequest)(nil),         // 3: api.CloneWorkflowTemplateRequest
	(*ListWorkflowTemplateVersionsRequest)(nil),  // 4: api.ListWorkflowTemplateVersionsRequest
	(*ListWorkflowTemplateVersionsResponse)(nil), // 5: api.ListWorkflowTemplateVersionsResponse
	(*ListWorkflowTemplatesRequest)(nil),         // 6: api.ListWorkflowTemplatesRequest
	(*ListWorkflowTemplatesResponse)(nil),        // 7: api.ListWorkflowTemplatesResponse
	(*ArchiveWorkflowTemplateRequest)(nil),       // 8: api.ArchiveWorkflowTemplateRequest
	(*ArchiveWorkflowTemplateResponse)(nil),      // 9: api.ArchiveWorkflowTemplateResponse
	(*WorkflowExecutionStatisticReport)(nil),     // 10: api.WorkflowExecutionStatisticReport
	(*WorkflowTemplate)(nil),                     // 11: api.WorkflowTemplate
}
var file_workflow_template_proto_depIdxs = []int32{
	11, // 0: api.CreateWorkflowTemplateRequest.workflowTemplate:type_name -> api.WorkflowTemplate
	11, // 1: api.UpdateWorkflowTemplateVersionRequest.workflowTemplate:type_name -> api.WorkflowTemplate
	11, // 2: api.ListWorkflowTemplateVersionsResponse.workflowTemplates:type_name -> api.WorkflowTemplate
	11, // 3: api.ListWorkflowTemplatesResponse.workflowTemplates:type_name -> api.WorkflowTemplate
	11, // 4: api.ArchiveWorkflowTemplateResponse.workflowTemplate:type_name -> api.WorkflowTemplate
	10, // 5: api.WorkflowTemplate.stats:type_name -> api.WorkflowExecutionStatisticReport
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_workflow_template_proto_init() }
func file_workflow_template_proto_init() {
	if File_workflow_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflow_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkflowTemplateRequest); i {
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
		file_workflow_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkflowTemplateVersionRequest); i {
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
		file_workflow_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkflowTemplateRequest); i {
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
		file_workflow_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloneWorkflowTemplateRequest); i {
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
		file_workflow_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkflowTemplateVersionsRequest); i {
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
		file_workflow_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkflowTemplateVersionsResponse); i {
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
		file_workflow_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkflowTemplatesRequest); i {
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
		file_workflow_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkflowTemplatesResponse); i {
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
		file_workflow_template_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveWorkflowTemplateRequest); i {
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
		file_workflow_template_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveWorkflowTemplateResponse); i {
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
		file_workflow_template_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowExecutionStatisticReport); i {
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
		file_workflow_template_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowTemplate); i {
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
			RawDescriptor: file_workflow_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflow_template_proto_goTypes,
		DependencyIndexes: file_workflow_template_proto_depIdxs,
		MessageInfos:      file_workflow_template_proto_msgTypes,
	}.Build()
	File_workflow_template_proto = out.File
	file_workflow_template_proto_rawDesc = nil
	file_workflow_template_proto_goTypes = nil
	file_workflow_template_proto_depIdxs = nil
}

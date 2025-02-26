// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: policies/pb/policies.proto

package pb

import (
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

type PolicyByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyID string `protobuf:"bytes,1,opt,name=policyID,proto3" json:"policyID,omitempty"`
	OwnerID  string `protobuf:"bytes,2,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
}

func (x *PolicyByIDReq) Reset() {
	*x = PolicyByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_pb_policies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyByIDReq) ProtoMessage() {}

func (x *PolicyByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_policies_pb_policies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyByIDReq.ProtoReflect.Descriptor instead.
func (*PolicyByIDReq) Descriptor() ([]byte, []int) {
	return file_policies_pb_policies_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyByIDReq) GetPolicyID() string {
	if x != nil {
		return x.PolicyID
	}
	return ""
}

func (x *PolicyByIDReq) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

type PoliciesByGroupsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupIDs []string `protobuf:"bytes,1,rep,name=groupIDs,proto3" json:"groupIDs,omitempty"`
	OwnerID  string   `protobuf:"bytes,2,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
}

func (x *PoliciesByGroupsReq) Reset() {
	*x = PoliciesByGroupsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_pb_policies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoliciesByGroupsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoliciesByGroupsReq) ProtoMessage() {}

func (x *PoliciesByGroupsReq) ProtoReflect() protoreflect.Message {
	mi := &file_policies_pb_policies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoliciesByGroupsReq.ProtoReflect.Descriptor instead.
func (*PoliciesByGroupsReq) Descriptor() ([]byte, []int) {
	return file_policies_pb_policies_proto_rawDescGZIP(), []int{1}
}

func (x *PoliciesByGroupsReq) GetGroupIDs() []string {
	if x != nil {
		return x.GroupIDs
	}
	return nil
}

func (x *PoliciesByGroupsReq) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

type PolicyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Backend string `protobuf:"bytes,3,opt,name=backend,proto3" json:"backend,omitempty"`
	Version int32  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Data    []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PolicyRes) Reset() {
	*x = PolicyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_pb_policies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRes) ProtoMessage() {}

func (x *PolicyRes) ProtoReflect() protoreflect.Message {
	mi := &file_policies_pb_policies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRes.ProtoReflect.Descriptor instead.
func (*PolicyRes) Descriptor() ([]byte, []int) {
	return file_policies_pb_policies_proto_rawDescGZIP(), []int{2}
}

func (x *PolicyRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PolicyRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyRes) GetBackend() string {
	if x != nil {
		return x.Backend
	}
	return ""
}

func (x *PolicyRes) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PolicyRes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PolicyInDSRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Backend   string `protobuf:"bytes,3,opt,name=backend,proto3" json:"backend,omitempty"`
	Version   int32  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Data      []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	DatasetId string `protobuf:"bytes,6,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
}

func (x *PolicyInDSRes) Reset() {
	*x = PolicyInDSRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_pb_policies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyInDSRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyInDSRes) ProtoMessage() {}

func (x *PolicyInDSRes) ProtoReflect() protoreflect.Message {
	mi := &file_policies_pb_policies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyInDSRes.ProtoReflect.Descriptor instead.
func (*PolicyInDSRes) Descriptor() ([]byte, []int) {
	return file_policies_pb_policies_proto_rawDescGZIP(), []int{3}
}

func (x *PolicyInDSRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PolicyInDSRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyInDSRes) GetBackend() string {
	if x != nil {
		return x.Backend
	}
	return ""
}

func (x *PolicyInDSRes) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PolicyInDSRes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PolicyInDSRes) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

type PolicyInDSListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*PolicyInDSRes `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *PolicyInDSListRes) Reset() {
	*x = PolicyInDSListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_pb_policies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyInDSListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyInDSListRes) ProtoMessage() {}

func (x *PolicyInDSListRes) ProtoReflect() protoreflect.Message {
	mi := &file_policies_pb_policies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyInDSListRes.ProtoReflect.Descriptor instead.
func (*PolicyInDSListRes) Descriptor() ([]byte, []int) {
	return file_policies_pb_policies_proto_rawDescGZIP(), []int{4}
}

func (x *PolicyInDSListRes) GetPolicies() []*PolicyInDSRes {
	if x != nil {
		return x.Policies
	}
	return nil
}

var File_policies_pb_policies_proto protoreflect.FileDescriptor

var file_policies_pb_policies_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4b, 0x0a,
	0x13, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x22, 0x77, 0x0a, 0x09, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e,
	0x44, 0x53, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x22, 0x48, 0x0a, 0x11, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x44, 0x53, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x44, 0x53, 0x52, 0x65, 0x73,
	0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x32, 0xab, 0x01, 0x0a, 0x0d, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x17,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x18, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x42, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x44, 0x53, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policies_pb_policies_proto_rawDescOnce sync.Once
	file_policies_pb_policies_proto_rawDescData = file_policies_pb_policies_proto_rawDesc
)

func file_policies_pb_policies_proto_rawDescGZIP() []byte {
	file_policies_pb_policies_proto_rawDescOnce.Do(func() {
		file_policies_pb_policies_proto_rawDescData = protoimpl.X.CompressGZIP(file_policies_pb_policies_proto_rawDescData)
	})
	return file_policies_pb_policies_proto_rawDescData
}

var file_policies_pb_policies_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_policies_pb_policies_proto_goTypes = []interface{}{
	(*PolicyByIDReq)(nil),       // 0: policies.PolicyByIDReq
	(*PoliciesByGroupsReq)(nil), // 1: policies.PoliciesByGroupsReq
	(*PolicyRes)(nil),           // 2: policies.PolicyRes
	(*PolicyInDSRes)(nil),       // 3: policies.PolicyInDSRes
	(*PolicyInDSListRes)(nil),   // 4: policies.PolicyInDSListRes
}
var file_policies_pb_policies_proto_depIdxs = []int32{
	3, // 0: policies.PolicyInDSListRes.policies:type_name -> policies.PolicyInDSRes
	0, // 1: policies.PolicyService.RetrievePolicy:input_type -> policies.PolicyByIDReq
	1, // 2: policies.PolicyService.RetrievePoliciesByGroups:input_type -> policies.PoliciesByGroupsReq
	2, // 3: policies.PolicyService.RetrievePolicy:output_type -> policies.PolicyRes
	4, // 4: policies.PolicyService.RetrievePoliciesByGroups:output_type -> policies.PolicyInDSListRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_policies_pb_policies_proto_init() }
func file_policies_pb_policies_proto_init() {
	if File_policies_pb_policies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policies_pb_policies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyByIDReq); i {
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
		file_policies_pb_policies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoliciesByGroupsReq); i {
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
		file_policies_pb_policies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyRes); i {
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
		file_policies_pb_policies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyInDSRes); i {
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
		file_policies_pb_policies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyInDSListRes); i {
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
			RawDescriptor: file_policies_pb_policies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policies_pb_policies_proto_goTypes,
		DependencyIndexes: file_policies_pb_policies_proto_depIdxs,
		MessageInfos:      file_policies_pb_policies_proto_msgTypes,
	}.Build()
	File_policies_pb_policies_proto = out.File
	file_policies_pb_policies_proto_rawDesc = nil
	file_policies_pb_policies_proto_goTypes = nil
	file_policies_pb_policies_proto_depIdxs = nil
}

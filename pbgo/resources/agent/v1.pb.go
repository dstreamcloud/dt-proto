// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: resources/agent/v1.proto

package agent_pb

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

type V1_Status int32

const (
	V1_UNKNOWN V1_Status = 0
	V1_IDLE    V1_Status = 1
	V1_BUSY    V1_Status = 2
)

// Enum value maps for V1_Status.
var (
	V1_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "IDLE",
		2: "BUSY",
	}
	V1_Status_value = map[string]int32{
		"UNKNOWN": 0,
		"IDLE":    1,
		"BUSY":    2,
	}
)

func (x V1_Status) Enum() *V1_Status {
	p := new(V1_Status)
	*p = x
	return p
}

func (x V1_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (V1_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_agent_v1_proto_enumTypes[0].Descriptor()
}

func (V1_Status) Type() protoreflect.EnumType {
	return &file_resources_agent_v1_proto_enumTypes[0]
}

func (x V1_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use V1_Status.Descriptor instead.
func (V1_Status) EnumDescriptor() ([]byte, []int) {
	return file_resources_agent_v1_proto_rawDescGZIP(), []int{0, 0}
}

type V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *V1) Reset() {
	*x = V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_agent_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1) ProtoMessage() {}

func (x *V1) ProtoReflect() protoreflect.Message {
	mi := &file_resources_agent_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1.ProtoReflect.Descriptor instead.
func (*V1) Descriptor() ([]byte, []int) {
	return file_resources_agent_v1_proto_rawDescGZIP(), []int{0}
}

type V1_Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt      int64     `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      int64     `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	OrganizationID string    `protobuf:"bytes,4,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
	OwnerID        string    `protobuf:"bytes,5,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
	Name           string    `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Token          string    `protobuf:"bytes,12,opt,name=token,proto3" json:"token,omitempty"`
	Status         V1_Status `protobuf:"varint,13,opt,name=status,proto3,enum=agent.V1_Status" json:"status,omitempty"`
}

func (x *V1_Agent) Reset() {
	*x = V1_Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_agent_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1_Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1_Agent) ProtoMessage() {}

func (x *V1_Agent) ProtoReflect() protoreflect.Message {
	mi := &file_resources_agent_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1_Agent.ProtoReflect.Descriptor instead.
func (*V1_Agent) Descriptor() ([]byte, []int) {
	return file_resources_agent_v1_proto_rawDescGZIP(), []int{0, 0}
}

func (x *V1_Agent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *V1_Agent) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *V1_Agent) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *V1_Agent) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

func (x *V1_Agent) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

func (x *V1_Agent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *V1_Agent) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *V1_Agent) GetStatus() V1_Status {
	if x != nil {
		return x.Status
	}
	return V1_UNKNOWN
}

var File_resources_agent_v1_proto protoreflect.FileDescriptor

var file_resources_agent_v1_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x22, 0x9b, 0x02, 0x0a, 0x02, 0x56, 0x31, 0x1a, 0xe9, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x44, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x59, 0x10, 0x02, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_agent_v1_proto_rawDescOnce sync.Once
	file_resources_agent_v1_proto_rawDescData = file_resources_agent_v1_proto_rawDesc
)

func file_resources_agent_v1_proto_rawDescGZIP() []byte {
	file_resources_agent_v1_proto_rawDescOnce.Do(func() {
		file_resources_agent_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_agent_v1_proto_rawDescData)
	})
	return file_resources_agent_v1_proto_rawDescData
}

var file_resources_agent_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_agent_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_agent_v1_proto_goTypes = []interface{}{
	(V1_Status)(0),   // 0: agent.V1.Status
	(*V1)(nil),       // 1: agent.V1
	(*V1_Agent)(nil), // 2: agent.V1.Agent
}
var file_resources_agent_v1_proto_depIdxs = []int32{
	0, // 0: agent.V1.Agent.status:type_name -> agent.V1.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resources_agent_v1_proto_init() }
func file_resources_agent_v1_proto_init() {
	if File_resources_agent_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_agent_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1); i {
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
		file_resources_agent_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1_Agent); i {
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
			RawDescriptor: file_resources_agent_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_agent_v1_proto_goTypes,
		DependencyIndexes: file_resources_agent_v1_proto_depIdxs,
		EnumInfos:         file_resources_agent_v1_proto_enumTypes,
		MessageInfos:      file_resources_agent_v1_proto_msgTypes,
	}.Build()
	File_resources_agent_v1_proto = out.File
	file_resources_agent_v1_proto_rawDesc = nil
	file_resources_agent_v1_proto_goTypes = nil
	file_resources_agent_v1_proto_depIdxs = nil
}

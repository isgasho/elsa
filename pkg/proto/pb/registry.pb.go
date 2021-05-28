// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActionType int32

const (
	ActionType_None        ActionType = 0
	ActionType_Normal      ActionType = 1
	ActionType_Replication ActionType = 2
)

var ActionType_name = map[int32]string{
	0: "None",
	1: "Normal",
	2: "Replication",
}

var ActionType_value = map[string]int32{
	"None":        0,
	"Normal":      1,
	"Replication": 2,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{0}
}

type RegisterRequest struct {
	Action               ActionType       `protobuf:"varint,1,opt,name=action,proto3,enum=com.busgo.registry.proto.ActionType" json:"action,omitempty"`
	Instance             *ServiceInstance `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetAction() ActionType {
	if m != nil {
		return m.Action
	}
	return ActionType_None
}

func (m *RegisterRequest) GetInstance() *ServiceInstance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type RegisterResponse struct {
	Code                 int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Instance             *ServiceInstance `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RegisterResponse) GetInstance() *ServiceInstance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type FetchRequest struct {
	Segment              string   `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{2}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetSegment() string {
	if m != nil {
		return m.Segment
	}
	return ""
}

func (m *FetchRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type FetchResponse struct {
	Code                 int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Instances            []*ServiceInstance `protobuf:"bytes,3,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{3}
}

func (m *FetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchResponse.Unmarshal(m, b)
}
func (m *FetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchResponse.Marshal(b, m, deterministic)
}
func (m *FetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchResponse.Merge(m, src)
}
func (m *FetchResponse) XXX_Size() int {
	return xxx_messageInfo_FetchResponse.Size(m)
}
func (m *FetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchResponse proto.InternalMessageInfo

func (m *FetchResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FetchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FetchResponse) GetInstances() []*ServiceInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type RenewRequest struct {
	Segment              string   `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewRequest) Reset()         { *m = RenewRequest{} }
func (m *RenewRequest) String() string { return proto.CompactTextString(m) }
func (*RenewRequest) ProtoMessage()    {}
func (*RenewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{4}
}

func (m *RenewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewRequest.Unmarshal(m, b)
}
func (m *RenewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewRequest.Marshal(b, m, deterministic)
}
func (m *RenewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewRequest.Merge(m, src)
}
func (m *RenewRequest) XXX_Size() int {
	return xxx_messageInfo_RenewRequest.Size(m)
}
func (m *RenewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenewRequest proto.InternalMessageInfo

func (m *RenewRequest) GetSegment() string {
	if m != nil {
		return m.Segment
	}
	return ""
}

func (m *RenewRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RenewRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RenewRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type RenewResponse struct {
	Code                 int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Instance             *ServiceInstance `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RenewResponse) Reset()         { *m = RenewResponse{} }
func (m *RenewResponse) String() string { return proto.CompactTextString(m) }
func (*RenewResponse) ProtoMessage()    {}
func (*RenewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{5}
}

func (m *RenewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewResponse.Unmarshal(m, b)
}
func (m *RenewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewResponse.Marshal(b, m, deterministic)
}
func (m *RenewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewResponse.Merge(m, src)
}
func (m *RenewResponse) XXX_Size() int {
	return xxx_messageInfo_RenewResponse.Size(m)
}
func (m *RenewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RenewResponse proto.InternalMessageInfo

func (m *RenewResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RenewResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RenewResponse) GetInstance() *ServiceInstance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type CancelRequest struct {
	Segment              string   `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelRequest) Reset()         { *m = CancelRequest{} }
func (m *CancelRequest) String() string { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()    {}
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{6}
}

func (m *CancelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelRequest.Unmarshal(m, b)
}
func (m *CancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelRequest.Marshal(b, m, deterministic)
}
func (m *CancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRequest.Merge(m, src)
}
func (m *CancelRequest) XXX_Size() int {
	return xxx_messageInfo_CancelRequest.Size(m)
}
func (m *CancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRequest proto.InternalMessageInfo

func (m *CancelRequest) GetSegment() string {
	if m != nil {
		return m.Segment
	}
	return ""
}

func (m *CancelRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *CancelRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CancelRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type CancelResponse struct {
	Code                 int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Instance             *ServiceInstance `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CancelResponse) Reset()         { *m = CancelResponse{} }
func (m *CancelResponse) String() string { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()    {}
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{7}
}

func (m *CancelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelResponse.Unmarshal(m, b)
}
func (m *CancelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelResponse.Marshal(b, m, deterministic)
}
func (m *CancelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResponse.Merge(m, src)
}
func (m *CancelResponse) XXX_Size() int {
	return xxx_messageInfo_CancelResponse.Size(m)
}
func (m *CancelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResponse proto.InternalMessageInfo

func (m *CancelResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CancelResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CancelResponse) GetInstance() *ServiceInstance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type ServiceInstance struct {
	Segment              string            `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	ServiceName          string            `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Ip                   string            `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32             `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RegTimestamp         int64             `protobuf:"varint,6,opt,name=regTimestamp,proto3" json:"regTimestamp,omitempty"`
	UpTimestamp          int64             `protobuf:"varint,7,opt,name=upTimestamp,proto3" json:"upTimestamp,omitempty"`
	RenewTimestamp       int64             `protobuf:"varint,8,opt,name=renewTimestamp,proto3" json:"renewTimestamp,omitempty"`
	DirtyTimestamp       int64             `protobuf:"varint,9,opt,name=dirtyTimestamp,proto3" json:"dirtyTimestamp,omitempty"`
	LatestTimestamp      int64             `protobuf:"varint,10,opt,name=latestTimestamp,proto3" json:"latestTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceInstance) Reset()         { *m = ServiceInstance{} }
func (m *ServiceInstance) String() string { return proto.CompactTextString(m) }
func (*ServiceInstance) ProtoMessage()    {}
func (*ServiceInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{8}
}

func (m *ServiceInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInstance.Unmarshal(m, b)
}
func (m *ServiceInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInstance.Marshal(b, m, deterministic)
}
func (m *ServiceInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInstance.Merge(m, src)
}
func (m *ServiceInstance) XXX_Size() int {
	return xxx_messageInfo_ServiceInstance.Size(m)
}
func (m *ServiceInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInstance.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInstance proto.InternalMessageInfo

func (m *ServiceInstance) GetSegment() string {
	if m != nil {
		return m.Segment
	}
	return ""
}

func (m *ServiceInstance) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceInstance) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ServiceInstance) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServiceInstance) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ServiceInstance) GetRegTimestamp() int64 {
	if m != nil {
		return m.RegTimestamp
	}
	return 0
}

func (m *ServiceInstance) GetUpTimestamp() int64 {
	if m != nil {
		return m.UpTimestamp
	}
	return 0
}

func (m *ServiceInstance) GetRenewTimestamp() int64 {
	if m != nil {
		return m.RenewTimestamp
	}
	return 0
}

func (m *ServiceInstance) GetDirtyTimestamp() int64 {
	if m != nil {
		return m.DirtyTimestamp
	}
	return 0
}

func (m *ServiceInstance) GetLatestTimestamp() int64 {
	if m != nil {
		return m.LatestTimestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("com.busgo.registry.proto.ActionType", ActionType_name, ActionType_value)
	proto.RegisterType((*RegisterRequest)(nil), "com.busgo.registry.proto.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "com.busgo.registry.proto.RegisterResponse")
	proto.RegisterType((*FetchRequest)(nil), "com.busgo.registry.proto.FetchRequest")
	proto.RegisterType((*FetchResponse)(nil), "com.busgo.registry.proto.FetchResponse")
	proto.RegisterType((*RenewRequest)(nil), "com.busgo.registry.proto.RenewRequest")
	proto.RegisterType((*RenewResponse)(nil), "com.busgo.registry.proto.RenewResponse")
	proto.RegisterType((*CancelRequest)(nil), "com.busgo.registry.proto.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "com.busgo.registry.proto.CancelResponse")
	proto.RegisterType((*ServiceInstance)(nil), "com.busgo.registry.proto.ServiceInstance")
	proto.RegisterMapType((map[string]string)(nil), "com.busgo.registry.proto.ServiceInstance.MetadataEntry")
}

func init() { proto.RegisterFile("registry.proto", fileDescriptor_41af05d40a615591) }

var fileDescriptor_41af05d40a615591 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0xc5, 0x49, 0x7f, 0x7e, 0xfd, 0x29, 0x8b, 0x43, 0xd4, 0x53, 0x15, 0xa1, 0x51, 0x76, 0xc8,
	0xa1, 0x3b, 0x80, 0x80, 0x0b, 0xa0, 0x81, 0x40, 0xa2, 0x07, 0x6f, 0x07, 0x04, 0x27, 0x37, 0xfd,
	0x28, 0x11, 0x4d, 0x1c, 0x6c, 0x77, 0x53, 0xee, 0x0c, 0xfe, 0x03, 0xfe, 0x21, 0xfe, 0x31, 0x14,
	0x27, 0x59, 0x9a, 0x4a, 0xed, 0xc6, 0x04, 0xdb, 0xcd, 0xfe, 0xfc, 0xde, 0xe7, 0xf7, 0x6c, 0xe7,
	0x05, 0xfa, 0x12, 0x97, 0x81, 0xd2, 0x32, 0xf1, 0x62, 0x29, 0xb4, 0xa0, 0x8e, 0x2f, 0x42, 0x6f,
	0xbe, 0x56, 0x4b, 0xe1, 0x55, 0x57, 0xdc, 0x5f, 0x04, 0x06, 0xcc, 0x94, 0x50, 0x32, 0xfc, 0xb6,
	0x46, 0xa5, 0xe9, 0x73, 0x68, 0x70, 0x5f, 0x07, 0x22, 0x72, 0xc8, 0x98, 0x4c, 0xfa, 0xd3, 0x07,
	0xde, 0x2e, 0xba, 0xf7, 0xc2, 0xe0, 0x4e, 0x93, 0x18, 0x59, 0xce, 0xa1, 0xc7, 0xd0, 0x0a, 0x22,
	0xa5, 0x79, 0xe4, 0xa3, 0x63, 0x8d, 0xc9, 0xa4, 0x33, 0x7d, 0xb4, 0x9b, 0x7f, 0x82, 0xf2, 0x2c,
	0xf0, 0xf1, 0x6d, 0x4e, 0x60, 0x97, 0x54, 0xf7, 0x27, 0x81, 0x61, 0x29, 0x4c, 0xc5, 0x22, 0x52,
	0x48, 0x29, 0xd4, 0x7c, 0xb1, 0x40, 0xa3, 0xab, 0xce, 0xcc, 0x98, 0x3a, 0xd0, 0x0c, 0x51, 0x29,
	0xbe, 0xcc, 0xb6, 0x6b, 0xb3, 0x62, 0x5a, 0x51, 0x62, 0xdf, 0x5c, 0xc9, 0x3b, 0xe8, 0xbe, 0x46,
	0xed, 0x7f, 0x29, 0x8e, 0xc7, 0x81, 0xa6, 0xc2, 0x65, 0x88, 0x91, 0x36, 0x3a, 0xda, 0xac, 0x98,
	0xd2, 0x31, 0x74, 0x54, 0xd6, 0x66, 0xc6, 0xc3, 0x42, 0xce, 0x66, 0xc9, 0xfd, 0x41, 0xa0, 0x97,
	0x37, 0xbb, 0x91, 0xa5, 0x37, 0xd0, 0x2e, 0x74, 0x29, 0xc7, 0x1e, 0xdb, 0x7f, 0xe7, 0xa9, 0xe4,
	0xba, 0x11, 0x74, 0x19, 0x46, 0x78, 0xfe, 0x0f, 0x4c, 0xd1, 0x3e, 0x58, 0x41, 0x6c, 0x4e, 0xb8,
	0xcd, 0xac, 0x20, 0x4e, 0x2d, 0xc5, 0x42, 0x6a, 0xa7, 0x96, 0x59, 0x4a, 0xc7, 0xee, 0x77, 0x02,
	0xbd, 0x7c, 0xc3, 0xbb, 0xbc, 0x4b, 0x01, 0xbd, 0x57, 0xe9, 0x60, 0x75, 0x5b, 0xbe, 0x2f, 0x08,
	0xf4, 0x8b, 0x1d, 0xef, 0xd2, 0xf8, 0x6f, 0x1b, 0x06, 0x5b, 0xab, 0xff, 0xdb, 0x3b, 0x3d, 0x81,
	0x56, 0x88, 0x9a, 0x2f, 0xb8, 0xe6, 0x4e, 0xdd, 0xbc, 0xd5, 0xc7, 0xd7, 0x96, 0xee, 0xbd, 0xcf,
	0x99, 0xc7, 0x91, 0x96, 0x09, 0xbb, 0x6c, 0x44, 0x5d, 0xe8, 0x4a, 0x5c, 0x9e, 0x06, 0x21, 0x2a,
	0xcd, 0xc3, 0xd8, 0x69, 0x8c, 0xc9, 0xc4, 0x66, 0x95, 0x5a, 0x2a, 0x7f, 0x1d, 0x97, 0x90, 0xa6,
	0x81, 0x6c, 0x96, 0xe8, 0x41, 0x1a, 0x91, 0x11, 0x9e, 0x97, 0xa0, 0x96, 0x01, 0x6d, 0x55, 0x53,
	0xdc, 0x22, 0x90, 0x3a, 0x29, 0x71, 0xed, 0x0c, 0x57, 0xad, 0xd2, 0x09, 0x0c, 0x56, 0x5c, 0xa3,
	0xd2, 0x25, 0x10, 0x0c, 0x70, 0xbb, 0x3c, 0x7a, 0x06, 0xbd, 0x8a, 0x35, 0x3a, 0x04, 0xfb, 0x2b,
	0x26, 0xf9, 0x0d, 0xa4, 0x43, 0x7a, 0x1f, 0xea, 0x67, 0x7c, 0xb5, 0x2e, 0xce, 0x3d, 0x9b, 0x3c,
	0xb5, 0x9e, 0x90, 0xc3, 0x23, 0x80, 0x32, 0x71, 0x69, 0x0b, 0x6a, 0x33, 0x11, 0xe1, 0xf0, 0x1e,
	0x05, 0x68, 0xcc, 0x84, 0x0c, 0xf9, 0x6a, 0x48, 0xe8, 0x00, 0x3a, 0x0c, 0xe3, 0x55, 0xe0, 0xf3,
	0x14, 0x38, 0xb4, 0xa6, 0x17, 0x76, 0x11, 0xf1, 0x32, 0xc9, 0x4f, 0x99, 0x72, 0x68, 0xc9, 0x3c,
	0x5c, 0xe9, 0x9e, 0xf7, 0xb4, 0xf5, 0x67, 0x18, 0x1d, 0x5e, 0x07, 0x9a, 0x3f, 0xf3, 0x0f, 0x50,
	0xff, 0x9c, 0x26, 0x1d, 0x3d, 0xd8, 0x4d, 0xda, 0xcc, 0xd5, 0xd1, 0xc3, 0x2b, 0x71, 0x65, 0x67,
	0x73, 0x4d, 0xfb, 0x3a, 0x6f, 0x86, 0xdb, 0xbe, 0xce, 0xd5, 0x4c, 0xfa, 0x04, 0x0d, 0xdf, 0x7c,
	0xac, 0x74, 0x0f, 0xa5, 0x12, 0x20, 0xa3, 0xc9, 0xd5, 0xc0, 0xac, 0xf9, 0xcb, 0xda, 0x47, 0x2b,
	0x9e, 0xcf, 0x1b, 0x66, 0xed, 0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x2b, 0x6f, 0x69,
	0xa3, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryServiceClient is the client API for RegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryServiceClient interface {
	// register a service instance
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// fetch the instance with segment and service name
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	// renew the instance
	Renew(ctx context.Context, in *RenewRequest, opts ...grpc.CallOption) (*RenewResponse, error)
	// cancel the instance
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type registryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistryServiceClient(cc *grpc.ClientConn) RegistryServiceClient {
	return &registryServiceClient{cc}
}

func (c *registryServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/com.busgo.registry.proto.RegistryService/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/com.busgo.registry.proto.RegistryService/fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) Renew(ctx context.Context, in *RenewRequest, opts ...grpc.CallOption) (*RenewResponse, error) {
	out := new(RenewResponse)
	err := c.cc.Invoke(ctx, "/com.busgo.registry.proto.RegistryService/renew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/com.busgo.registry.proto.RegistryService/cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServiceServer is the server API for RegistryService service.
type RegistryServiceServer interface {
	// register a service instance
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// fetch the instance with segment and service name
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	// renew the instance
	Renew(context.Context, *RenewRequest) (*RenewResponse, error)
	// cancel the instance
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
}

// UnimplementedRegistryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryServiceServer struct {
}

func (*UnimplementedRegistryServiceServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedRegistryServiceServer) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedRegistryServiceServer) Renew(ctx context.Context, req *RenewRequest) (*RenewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Renew not implemented")
}
func (*UnimplementedRegistryServiceServer) Cancel(ctx context.Context, req *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterRegistryServiceServer(s *grpc.Server, srv RegistryServiceServer) {
	s.RegisterService(&_RegistryService_serviceDesc, srv)
}

func _RegistryService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.busgo.registry.proto.RegistryService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.busgo.registry.proto.RegistryService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.busgo.registry.proto.RegistryService/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Renew(ctx, req.(*RenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.busgo.registry.proto.RegistryService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.busgo.registry.proto.RegistryService",
	HandlerType: (*RegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _RegistryService_Register_Handler,
		},
		{
			MethodName: "fetch",
			Handler:    _RegistryService_Fetch_Handler,
		},
		{
			MethodName: "renew",
			Handler:    _RegistryService_Renew_Handler,
		},
		{
			MethodName: "cancel",
			Handler:    _RegistryService_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}

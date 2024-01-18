// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_availabilityset.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type AvailabilitySetRequest struct {
	AvailabilitySets     []*AvailabilitySet `protobuf:"bytes,1,rep,name=AvailabilitySets,proto3" json:"AvailabilitySets,omitempty"`
	OperationType        common.Operation   `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AvailabilitySetRequest) Reset()         { *m = AvailabilitySetRequest{} }
func (m *AvailabilitySetRequest) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySetRequest) ProtoMessage()    {}
func (*AvailabilitySetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{0}
}

func (m *AvailabilitySetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySetRequest.Unmarshal(m, b)
}
func (m *AvailabilitySetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySetRequest.Marshal(b, m, deterministic)
}
func (m *AvailabilitySetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySetRequest.Merge(m, src)
}
func (m *AvailabilitySetRequest) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySetRequest.Size(m)
}
func (m *AvailabilitySetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySetRequest proto.InternalMessageInfo

func (m *AvailabilitySetRequest) GetAvailabilitySets() []*AvailabilitySet {
	if m != nil {
		return m.AvailabilitySets
	}
	return nil
}

func (m *AvailabilitySetRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type AvailabilitySetResponse struct {
	AvailabilitySets     []*AvailabilitySet  `protobuf:"bytes,1,rep,name=AvailabilitySets,proto3" json:"AvailabilitySets,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AvailabilitySetResponse) Reset()         { *m = AvailabilitySetResponse{} }
func (m *AvailabilitySetResponse) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySetResponse) ProtoMessage()    {}
func (*AvailabilitySetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{1}
}

func (m *AvailabilitySetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySetResponse.Unmarshal(m, b)
}
func (m *AvailabilitySetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySetResponse.Marshal(b, m, deterministic)
}
func (m *AvailabilitySetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySetResponse.Merge(m, src)
}
func (m *AvailabilitySetResponse) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySetResponse.Size(m)
}
func (m *AvailabilitySetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySetResponse proto.InternalMessageInfo

func (m *AvailabilitySetResponse) GetAvailabilitySets() []*AvailabilitySet {
	if m != nil {
		return m.AvailabilitySets
	}
	return nil
}

func (m *AvailabilitySetResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *AvailabilitySetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// avset structure is modeled after the Azure sdk for go at
// https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/resourcemanager/compute/armcompute/models.go
type AvailabilitySet struct {
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Tags                 *common.Tags               `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	Entity               *common.Entity             `protobuf:"bytes,7,opt,name=entity,proto3" json:"entity,omitempty"`
	Properties           *AvailabilitySetProperties `protobuf:"bytes,8,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AvailabilitySet) Reset()         { *m = AvailabilitySet{} }
func (m *AvailabilitySet) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySet) ProtoMessage()    {}
func (*AvailabilitySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{2}
}

func (m *AvailabilitySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySet.Unmarshal(m, b)
}
func (m *AvailabilitySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySet.Marshal(b, m, deterministic)
}
func (m *AvailabilitySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySet.Merge(m, src)
}
func (m *AvailabilitySet) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySet.Size(m)
}
func (m *AvailabilitySet) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySet.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySet proto.InternalMessageInfo

func (m *AvailabilitySet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AvailabilitySet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AvailabilitySet) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *AvailabilitySet) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *AvailabilitySet) GetProperties() *AvailabilitySetProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

type AvailabilitySetProperties struct {
	PlatformFaultDomainCount string                    `protobuf:"bytes,1,opt,name=platformFaultDomainCount,proto3" json:"platformFaultDomainCount,omitempty"`
	VirtualMachines          []*common.NodeSubResource `protobuf:"bytes,2,rep,name=virtualMachines,proto3" json:"virtualMachines,omitempty"`
	Status                   *common.Status            `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *AvailabilitySetProperties) Reset()         { *m = AvailabilitySetProperties{} }
func (m *AvailabilitySetProperties) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySetProperties) ProtoMessage()    {}
func (*AvailabilitySetProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{3}
}

func (m *AvailabilitySetProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySetProperties.Unmarshal(m, b)
}
func (m *AvailabilitySetProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySetProperties.Marshal(b, m, deterministic)
}
func (m *AvailabilitySetProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySetProperties.Merge(m, src)
}
func (m *AvailabilitySetProperties) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySetProperties.Size(m)
}
func (m *AvailabilitySetProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySetProperties.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySetProperties proto.InternalMessageInfo

func (m *AvailabilitySetProperties) GetPlatformFaultDomainCount() string {
	if m != nil {
		return m.PlatformFaultDomainCount
	}
	return ""
}

func (m *AvailabilitySetProperties) GetVirtualMachines() []*common.NodeSubResource {
	if m != nil {
		return m.VirtualMachines
	}
	return nil
}

func (m *AvailabilitySetProperties) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AvailabilitySetRequest)(nil), "moc.nodeagent.compute.AvailabilitySetRequest")
	proto.RegisterType((*AvailabilitySetResponse)(nil), "moc.nodeagent.compute.AvailabilitySetResponse")
	proto.RegisterType((*AvailabilitySet)(nil), "moc.nodeagent.compute.AvailabilitySet")
	proto.RegisterType((*AvailabilitySetProperties)(nil), "moc.nodeagent.compute.AvailabilitySetProperties")
}

func init() {
	proto.RegisterFile("moc_nodeagent_availabilityset.proto", fileDescriptor_d615df2311cce965)
}

var fileDescriptor_d615df2311cce965 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0x49, 0x37, 0x0a, 0x75, 0x45, 0x87, 0xac, 0xc2, 0x42, 0x25, 0x50, 0xd5, 0x49, 0xa8,
	0x17, 0x92, 0xa9, 0x70, 0xe2, 0x80, 0xb4, 0xc1, 0x90, 0x38, 0x00, 0x93, 0x3b, 0x71, 0xe0, 0x32,
	0xb9, 0xe9, 0xd7, 0xcc, 0x22, 0xf1, 0x67, 0xec, 0xcf, 0x45, 0x3d, 0xf2, 0x57, 0xf8, 0x13, 0x5c,
	0xf8, 0x07, 0xfc, 0x29, 0x54, 0x27, 0xab, 0x46, 0xc6, 0xa4, 0x5e, 0x38, 0xb5, 0xf1, 0xfb, 0xf8,
	0xf5, 0xeb, 0x37, 0x5f, 0xd8, 0x41, 0x89, 0xd9, 0xb9, 0xc6, 0x39, 0xc8, 0x1c, 0x34, 0x9d, 0xcb,
	0xa5, 0x54, 0x85, 0x9c, 0xa9, 0x42, 0xd1, 0xca, 0x01, 0x25, 0xc6, 0x22, 0x21, 0x7f, 0x50, 0x62,
	0x96, 0x6c, 0xa0, 0x24, 0xc3, 0xd2, 0x78, 0x82, 0xc1, 0x93, 0x1c, 0x31, 0x2f, 0x20, 0x0d, 0xd0,
	0xcc, 0x2f, 0xd2, 0x6f, 0x56, 0x1a, 0x03, 0xd6, 0x55, 0xdb, 0x06, 0xfb, 0x6b, 0xef, 0x0c, 0xcb,
	0x12, 0x75, 0xfd, 0x53, 0x09, 0xa3, 0x1f, 0x11, 0x7b, 0x78, 0x74, 0xe5, 0xa4, 0x29, 0x90, 0x80,
	0xaf, 0x1e, 0x1c, 0x71, 0xc1, 0xee, 0x37, 0x14, 0x17, 0x47, 0xc3, 0x9d, 0x71, 0x77, 0xf2, 0x34,
	0xf9, 0x67, 0x8a, 0xa4, 0x69, 0x74, 0x6d, 0x3f, 0x7f, 0xc1, 0xee, 0x7d, 0x34, 0x60, 0x25, 0x29,
	0xd4, 0x67, 0x2b, 0x03, 0x71, 0x6b, 0x18, 0x8d, 0x7b, 0x93, 0x5e, 0x30, 0xdc, 0x28, 0xe2, 0x6f,
	0x68, 0xf4, 0x33, 0x62, 0xfb, 0xd7, 0x42, 0x3a, 0x83, 0xda, 0xc1, 0x7f, 0x49, 0x39, 0x61, 0x6d,
	0x01, 0xce, 0x17, 0x14, 0xe2, 0x75, 0x27, 0x83, 0xa4, 0xaa, 0x37, 0xb9, 0xac, 0x37, 0x39, 0x46,
	0x2c, 0x3e, 0xc9, 0xc2, 0x83, 0xa8, 0x49, 0xde, 0x67, 0xb7, 0x4f, 0xac, 0x45, 0x1b, 0xef, 0x0c,
	0xa3, 0x71, 0x47, 0x54, 0x0f, 0xa3, 0xdf, 0x11, 0xdb, 0x6b, 0xd8, 0x73, 0xce, 0x76, 0xb5, 0x2c,
	0x21, 0x8e, 0x02, 0x18, 0xfe, 0xf3, 0x1e, 0x6b, 0xa9, 0x79, 0x38, 0xad, 0x23, 0x5a, 0x6a, 0xce,
	0x1f, 0xb3, 0x5d, 0x92, 0xb9, 0x8b, 0xdb, 0xe1, 0xfc, 0x4e, 0xb8, 0xc9, 0x99, 0xcc, 0x9d, 0x08,
	0xcb, 0xfc, 0x80, 0xb5, 0x41, 0x93, 0xa2, 0x55, 0x7c, 0x27, 0x00, 0xdd, 0x00, 0x9c, 0x84, 0x25,
	0x51, 0x4b, 0xfc, 0x94, 0x31, 0x63, 0xd1, 0x80, 0x25, 0x05, 0x2e, 0xbe, 0x1b, 0xc0, 0xc3, 0xed,
	0x3a, 0x39, 0xdd, 0xec, 0x13, 0x57, 0x3c, 0x46, 0xbf, 0x22, 0xf6, 0xe8, 0x46, 0x92, 0xbf, 0x64,
	0xb1, 0x29, 0x24, 0x2d, 0xd0, 0x96, 0x6f, 0xa5, 0x2f, 0xe8, 0x0d, 0x96, 0x52, 0xe9, 0xd7, 0xe8,
	0x35, 0xd5, 0x77, 0xbd, 0x51, 0xe7, 0xaf, 0xd8, 0xde, 0x52, 0x59, 0xf2, 0xb2, 0x78, 0x2f, 0xb3,
	0x0b, 0xa5, 0xc1, 0xc5, 0xad, 0xf0, 0x12, 0xfb, 0x21, 0xf0, 0x07, 0x9c, 0xc3, 0xd4, 0xcf, 0x04,
	0x38, 0xf4, 0x36, 0x03, 0xd1, 0x84, 0xd7, 0x85, 0x38, 0x92, 0xe4, 0x5d, 0xa8, 0xff, 0xb2, 0x90,
	0x69, 0x58, 0x12, 0xb5, 0x34, 0xf9, 0x1e, 0xb1, 0x7e, 0x23, 0xfe, 0xd1, 0xba, 0x05, 0xae, 0x58,
	0xfb, 0x9d, 0x5e, 0xe2, 0x17, 0xe0, 0xcf, 0xb6, 0x9c, 0x99, 0xea, 0x13, 0x19, 0x24, 0xdb, 0xe2,
	0xd5, 0xb0, 0x8e, 0x6e, 0x1d, 0x1f, 0x7e, 0x4e, 0x72, 0x45, 0x17, 0x7e, 0xb6, 0x66, 0xd3, 0x52,
	0x65, 0x16, 0x1d, 0x2e, 0x28, 0x2d, 0x31, 0x4b, 0xad, 0xc9, 0xd2, 0x8d, 0x57, 0x5a, 0x7b, 0xcd,
	0xda, 0x61, 0xe8, 0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x97, 0x9e, 0x3d, 0xb6, 0x1f, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AvailabilitySetAgentClient is the client API for AvailabilitySetAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AvailabilitySetAgentClient interface {
	Invoke(ctx context.Context, in *AvailabilitySetRequest, opts ...grpc.CallOption) (*AvailabilitySetResponse, error)
}

type availabilitySetAgentClient struct {
	cc *grpc.ClientConn
}

func NewAvailabilitySetAgentClient(cc *grpc.ClientConn) AvailabilitySetAgentClient {
	return &availabilitySetAgentClient{cc}
}

func (c *availabilitySetAgentClient) Invoke(ctx context.Context, in *AvailabilitySetRequest, opts ...grpc.CallOption) (*AvailabilitySetResponse, error) {
	out := new(AvailabilitySetResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.compute.AvailabilitySetAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvailabilitySetAgentServer is the server API for AvailabilitySetAgent service.
type AvailabilitySetAgentServer interface {
	Invoke(context.Context, *AvailabilitySetRequest) (*AvailabilitySetResponse, error)
}

// UnimplementedAvailabilitySetAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAvailabilitySetAgentServer struct {
}

func (*UnimplementedAvailabilitySetAgentServer) Invoke(ctx context.Context, req *AvailabilitySetRequest) (*AvailabilitySetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterAvailabilitySetAgentServer(s *grpc.Server, srv AvailabilitySetAgentServer) {
	s.RegisterService(&_AvailabilitySetAgent_serviceDesc, srv)
}

func _AvailabilitySetAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailabilitySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilitySetAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.compute.AvailabilitySetAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilitySetAgentServer).Invoke(ctx, req.(*AvailabilitySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AvailabilitySetAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.compute.AvailabilitySetAgent",
	HandlerType: (*AvailabilitySetAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _AvailabilitySetAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_availabilityset.proto",
}

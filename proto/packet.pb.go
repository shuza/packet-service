// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/packet.proto

package packet

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c344942e71fec3df, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_c344942e71fec3df, []int{1}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Item) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Item) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Packet struct {
	Is                   string   `protobuf:"bytes,1,opt,name=is,proto3" json:"is,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32    `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Items                []*Item  `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	VesselId             string   `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c344942e71fec3df, []int{2}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetIs() string {
	if m != nil {
		return m.Is
	}
	return ""
}

func (m *Packet) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Packet) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Packet) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Packet) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Response struct {
	Created              bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Packet               *Packet   `protobuf:"bytes,2,opt,name=packet,proto3" json:"packet,omitempty"`
	Packets              []*Packet `protobuf:"bytes,3,rep,name=packets,proto3" json:"packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c344942e71fec3df, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetPacket() *Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *Response) GetPackets() []*Packet {
	if m != nil {
		return m.Packets
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "packet.Empty")
	proto.RegisterType((*Item)(nil), "packet.Item")
	proto.RegisterType((*Packet)(nil), "packet.Packet")
	proto.RegisterType((*Response)(nil), "packet.Response")
}

func init() { proto.RegisterFile("proto/packet.proto", fileDescriptor_c344942e71fec3df) }

var fileDescriptor_c344942e71fec3df = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x4a, 0xf3, 0x40,
	0x10, 0xc5, 0xbf, 0xfe, 0x49, 0xd2, 0x4e, 0xda, 0xf2, 0x31, 0x17, 0x1a, 0xf4, 0xc2, 0x92, 0x0b,
	0xe9, 0x55, 0x0b, 0xf1, 0x11, 0x44, 0x24, 0x77, 0x12, 0x1f, 0x40, 0x6a, 0x76, 0x68, 0x17, 0x4d,
	0x13, 0x76, 0xb6, 0x15, 0x9f, 0xc3, 0x17, 0x96, 0xcc, 0xee, 0x8a, 0x8a, 0x77, 0x3d, 0xe7, 0x74,
	0xf2, 0x3b, 0x33, 0x09, 0x60, 0x67, 0x5a, 0xdb, 0x6e, 0xba, 0x6d, 0xfd, 0x42, 0x76, 0x2d, 0x02,
	0x63, 0xa7, 0xf2, 0x04, 0xa2, 0xbb, 0xa6, 0xb3, 0xef, 0xf9, 0x1e, 0xc6, 0xa5, 0xa5, 0x06, 0x17,
	0x30, 0xd4, 0x2a, 0x1b, 0x2c, 0x07, 0xab, 0x69, 0x35, 0xd4, 0x0a, 0xaf, 0x20, 0xad, 0x8f, 0x6c,
	0xdb, 0x86, 0xcc, 0x93, 0x56, 0xd9, 0x50, 0x02, 0x08, 0x56, 0xa9, 0xf0, 0x0c, 0xe2, 0xd6, 0xe8,
	0x9d, 0x3e, 0x64, 0x23, 0xc9, 0xbc, 0xc2, 0x73, 0x48, 0x8e, 0xec, 0x86, 0xc6, 0x2e, 0xe8, 0x65,
	0xa9, 0xf2, 0x8f, 0x01, 0xc4, 0x0f, 0x42, 0x17, 0x18, 0x7f, 0xc1, 0x18, 0x97, 0x90, 0x2a, 0xe2,
	0xda, 0xe8, 0xce, 0xea, 0xf6, 0xe0, 0x61, 0xdf, 0xad, 0x9e, 0xf6, 0x46, 0x7a, 0xb7, 0xb7, 0x42,
	0x8b, 0x2a, 0xaf, 0x30, 0x87, 0x48, 0x5b, 0x6a, 0x38, 0x1b, 0x2f, 0x47, 0xab, 0xb4, 0x98, 0xad,
	0xfd, 0xb6, 0xfd, 0x4e, 0x95, 0x8b, 0xf0, 0x12, 0xa6, 0x27, 0x62, 0xa6, 0xd7, 0xbe, 0x53, 0x24,
	0xcf, 0x9e, 0x38, 0xa3, 0x54, 0xf9, 0x09, 0x26, 0x15, 0x71, 0xd7, 0x1e, 0x98, 0x30, 0x83, 0xa4,
	0x36, 0xb4, 0xb5, 0xe4, 0x0e, 0x31, 0xa9, 0x82, 0xc4, 0x6b, 0xf0, 0x87, 0x93, 0x6e, 0x69, 0xb1,
	0x08, 0x1c, 0xb7, 0x50, 0xe5, 0x53, 0x5c, 0x41, 0xe2, 0x7e, 0x71, 0x36, 0x92, 0x42, 0xbf, 0xff,
	0x18, 0xe2, 0xc2, 0xc2, 0xdc, 0x59, 0x8f, 0x64, 0x4e, 0xba, 0x26, 0xdc, 0x00, 0xdc, 0x93, 0x75,
	0x1e, 0xe3, 0x3c, 0xcc, 0xc9, 0x5b, 0xba, 0xf8, 0x1f, 0x64, 0xe8, 0x9a, 0xff, 0xc3, 0x02, 0x66,
	0xb7, 0x52, 0x2f, 0x1c, 0xf5, 0x27, 0xea, 0xaf, 0x99, 0xe7, 0x58, 0xbe, 0x82, 0x9b, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x31, 0x11, 0x99, 0x1b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PacketServiceClient is the client API for PacketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PacketServiceClient interface {
	GetPackets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	CreatePacket(ctx context.Context, in *Packet, opts ...grpc.CallOption) (*Response, error)
}

type packetServiceClient struct {
	cc *grpc.ClientConn
}

func NewPacketServiceClient(cc *grpc.ClientConn) PacketServiceClient {
	return &packetServiceClient{cc}
}

func (c *packetServiceClient) GetPackets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/packet.PacketService/GetPackets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetServiceClient) CreatePacket(ctx context.Context, in *Packet, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/packet.PacketService/CreatePacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PacketServiceServer is the server API for PacketService service.
type PacketServiceServer interface {
	GetPackets(context.Context, *Empty) (*Response, error)
	CreatePacket(context.Context, *Packet) (*Response, error)
}

// UnimplementedPacketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPacketServiceServer struct {
}

func (*UnimplementedPacketServiceServer) GetPackets(ctx context.Context, req *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackets not implemented")
}
func (*UnimplementedPacketServiceServer) CreatePacket(ctx context.Context, req *Packet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePacket not implemented")
}

func RegisterPacketServiceServer(s *grpc.Server, srv PacketServiceServer) {
	s.RegisterService(&_PacketService_serviceDesc, srv)
}

func _PacketService_GetPackets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketServiceServer).GetPackets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packet.PacketService/GetPackets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketServiceServer).GetPackets(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PacketService_CreatePacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Packet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketServiceServer).CreatePacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packet.PacketService/CreatePacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketServiceServer).CreatePacket(ctx, req.(*Packet))
	}
	return interceptor(ctx, in, info, handler)
}

var _PacketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "packet.PacketService",
	HandlerType: (*PacketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPackets",
			Handler:    _PacketService_GetPackets_Handler,
		},
		{
			MethodName: "CreatePacket",
			Handler:    _PacketService_CreatePacket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/packet.proto",
}
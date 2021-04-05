// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

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

type HBRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mid                  int64    `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HBRequest) Reset()         { *m = HBRequest{} }
func (m *HBRequest) String() string { return proto.CompactTextString(m) }
func (*HBRequest) ProtoMessage()    {}
func (*HBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{0}
}

func (m *HBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HBRequest.Unmarshal(m, b)
}
func (m *HBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HBRequest.Marshal(b, m, deterministic)
}
func (m *HBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HBRequest.Merge(m, src)
}
func (m *HBRequest) XXX_Size() int {
	return xxx_messageInfo_HBRequest.Size(m)
}
func (m *HBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HBRequest proto.InternalMessageInfo

func (m *HBRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HBRequest) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

type HBReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HBReply) Reset()         { *m = HBReply{} }
func (m *HBReply) String() string { return proto.CompactTextString(m) }
func (*HBReply) ProtoMessage()    {}
func (*HBReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{1}
}

func (m *HBReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HBReply.Unmarshal(m, b)
}
func (m *HBReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HBReply.Marshal(b, m, deterministic)
}
func (m *HBReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HBReply.Merge(m, src)
}
func (m *HBReply) XXX_Size() int {
	return xxx_messageInfo_HBReply.Size(m)
}
func (m *HBReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HBReply.DiscardUnknown(m)
}

var xxx_messageInfo_HBReply proto.InternalMessageInfo

func (m *HBReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*HBRequest)(nil), "service.HBRequest")
	proto.RegisterType((*HBReply)(nil), "service.HBReply")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_3c667767fb9826a9) }

var fileDescriptor_3c667767fb9826a9 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0xd2, 0xe5, 0xe2, 0xf4, 0x70, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca,
	0x4c, 0x11, 0x12, 0xe0, 0x62, 0xce, 0xcd, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x02,
	0x31, 0x95, 0x64, 0xb9, 0xd8, 0x41, 0xca, 0x0b, 0x72, 0x2a, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf3,
	0x53, 0x52, 0xc1, 0xca, 0x59, 0x83, 0xc0, 0x6c, 0x23, 0x07, 0x2e, 0x4e, 0x0f, 0x90, 0x4d, 0x4e,
	0xa9, 0x89, 0x25, 0x42, 0xc6, 0xc8, 0x1c, 0x21, 0x3d, 0xa8, 0x8d, 0x7a, 0x70, 0xeb, 0xa4, 0x04,
	0x50, 0xc4, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x9c, 0xd8, 0xa2, 0x58, 0xf4, 0xac, 0x0b, 0x92, 0x92,
	0xd8, 0xc0, 0xee, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x67, 0xb1, 0x98, 0xba, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartBeatClient is the client API for HeartBeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartBeatClient interface {
	HeartBeat(ctx context.Context, in *HBRequest, opts ...grpc.CallOption) (*HBReply, error)
}

type heartBeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartBeatClient(cc *grpc.ClientConn) HeartBeatClient {
	return &heartBeatClient{cc}
}

func (c *heartBeatClient) HeartBeat(ctx context.Context, in *HBRequest, opts ...grpc.CallOption) (*HBReply, error) {
	out := new(HBReply)
	err := c.cc.Invoke(ctx, "/service.HeartBeat/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartBeatServer is the server API for HeartBeat service.
type HeartBeatServer interface {
	HeartBeat(context.Context, *HBRequest) (*HBReply, error)
}

// UnimplementedHeartBeatServer can be embedded to have forward compatible implementations.
type UnimplementedHeartBeatServer struct {
}

func (*UnimplementedHeartBeatServer) HeartBeat(ctx context.Context, req *HBRequest) (*HBReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}

func RegisterHeartBeatServer(s *grpc.Server, srv HeartBeatServer) {
	s.RegisterService(&_HeartBeat_serviceDesc, srv)
}

func _HeartBeat_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeatServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.HeartBeat/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeatServer).HeartBeat(ctx, req.(*HBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeartBeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.HeartBeat",
	HandlerType: (*HeartBeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _HeartBeat_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}
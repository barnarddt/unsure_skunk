// Code generated by protoc-gen-go. DO NOT EDIT.
// source: skunk.proto

package skunkpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import reflexpb "github.com/luno/reflex/reflexpb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_skunk_009bec3b4856f880, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type GetDataReq struct {
	RoundId              int64    `protobuf:"varint,1,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
	Player               string   `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataReq) Reset()         { *m = GetDataReq{} }
func (m *GetDataReq) String() string { return proto.CompactTextString(m) }
func (*GetDataReq) ProtoMessage()    {}
func (*GetDataReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_skunk_009bec3b4856f880, []int{1}
}
func (m *GetDataReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataReq.Unmarshal(m, b)
}
func (m *GetDataReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataReq.Marshal(b, m, deterministic)
}
func (dst *GetDataReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataReq.Merge(dst, src)
}
func (m *GetDataReq) XXX_Size() int {
	return xxx_messageInfo_GetDataReq.Size(m)
}
func (m *GetDataReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataReq proto.InternalMessageInfo

func (m *GetDataReq) GetRoundId() int64 {
	if m != nil {
		return m.RoundId
	}
	return 0
}

func (m *GetDataReq) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

type GetDataRes struct {
	Part                 []*Part  `protobuf:"bytes,1,rep,name=part,proto3" json:"part,omitempty"`
	Rank                 int32    `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataRes) Reset()         { *m = GetDataRes{} }
func (m *GetDataRes) String() string { return proto.CompactTextString(m) }
func (*GetDataRes) ProtoMessage()    {}
func (*GetDataRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_skunk_009bec3b4856f880, []int{2}
}
func (m *GetDataRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataRes.Unmarshal(m, b)
}
func (m *GetDataRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataRes.Marshal(b, m, deterministic)
}
func (dst *GetDataRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRes.Merge(dst, src)
}
func (m *GetDataRes) XXX_Size() int {
	return xxx_messageInfo_GetDataRes.Size(m)
}
func (m *GetDataRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRes proto.InternalMessageInfo

func (m *GetDataRes) GetPart() []*Part {
	if m != nil {
		return m.Part
	}
	return nil
}

func (m *GetDataRes) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type Part struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RoundId              int64    `protobuf:"varint,2,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
	Player               string   `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
	Part                 int64    `protobuf:"varint,4,opt,name=part,proto3" json:"part,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Part) Reset()         { *m = Part{} }
func (m *Part) String() string { return proto.CompactTextString(m) }
func (*Part) ProtoMessage()    {}
func (*Part) Descriptor() ([]byte, []int) {
	return fileDescriptor_skunk_009bec3b4856f880, []int{3}
}
func (m *Part) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Part.Unmarshal(m, b)
}
func (m *Part) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Part.Marshal(b, m, deterministic)
}
func (dst *Part) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Part.Merge(dst, src)
}
func (m *Part) XXX_Size() int {
	return xxx_messageInfo_Part.Size(m)
}
func (m *Part) XXX_DiscardUnknown() {
	xxx_messageInfo_Part.DiscardUnknown(m)
}

var xxx_messageInfo_Part proto.InternalMessageInfo

func (m *Part) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Part) GetRoundId() int64 {
	if m != nil {
		return m.RoundId
	}
	return 0
}

func (m *Part) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

func (m *Part) GetPart() int64 {
	if m != nil {
		return m.Part
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "skunkpb.Empty")
	proto.RegisterType((*GetDataReq)(nil), "skunkpb.GetDataReq")
	proto.RegisterType((*GetDataRes)(nil), "skunkpb.GetDataRes")
	proto.RegisterType((*Part)(nil), "skunkpb.Part")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SkunkClient is the client API for Skunk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SkunkClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetData(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetDataRes, error)
	Stream(ctx context.Context, in *reflexpb.StreamRequest, opts ...grpc.CallOption) (Skunk_StreamClient, error)
}

type skunkClient struct {
	cc *grpc.ClientConn
}

func NewSkunkClient(cc *grpc.ClientConn) SkunkClient {
	return &skunkClient{cc}
}

func (c *skunkClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/skunkpb.Skunk/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skunkClient) GetData(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetDataRes, error) {
	out := new(GetDataRes)
	err := c.cc.Invoke(ctx, "/skunkpb.Skunk/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skunkClient) Stream(ctx context.Context, in *reflexpb.StreamRequest, opts ...grpc.CallOption) (Skunk_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Skunk_serviceDesc.Streams[0], "/skunkpb.Skunk/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &skunkStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Skunk_StreamClient interface {
	Recv() (*reflexpb.Event, error)
	grpc.ClientStream
}

type skunkStreamClient struct {
	grpc.ClientStream
}

func (x *skunkStreamClient) Recv() (*reflexpb.Event, error) {
	m := new(reflexpb.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SkunkServer is the server API for Skunk service.
type SkunkServer interface {
	Ping(context.Context, *Empty) (*Empty, error)
	GetData(context.Context, *GetDataReq) (*GetDataRes, error)
	Stream(*reflexpb.StreamRequest, Skunk_StreamServer) error
}

func RegisterSkunkServer(s *grpc.Server, srv SkunkServer) {
	s.RegisterService(&_Skunk_serviceDesc, srv)
}

func _Skunk_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkunkServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skunkpb.Skunk/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkunkServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Skunk_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkunkServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skunkpb.Skunk/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkunkServer).GetData(ctx, req.(*GetDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Skunk_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(reflexpb.StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SkunkServer).Stream(m, &skunkStreamServer{stream})
}

type Skunk_StreamServer interface {
	Send(*reflexpb.Event) error
	grpc.ServerStream
}

type skunkStreamServer struct {
	grpc.ServerStream
}

func (x *skunkStreamServer) Send(m *reflexpb.Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Skunk_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skunkpb.Skunk",
	HandlerType: (*SkunkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Skunk_Ping_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _Skunk_GetData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Skunk_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "skunk.proto",
}

func init() { proto.RegisterFile("skunk.proto", fileDescriptor_skunk_009bec3b4856f880) }

var fileDescriptor_skunk_009bec3b4856f880 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xeb, 0x34, 0x6d, 0xfe, 0xff, 0xab, 0x28, 0xd2, 0x21, 0x41, 0xc9, 0x54, 0x3c, 0x65,
	0x40, 0x2e, 0x2a, 0x82, 0x95, 0x01, 0x2a, 0xc4, 0x56, 0xb9, 0x33, 0x42, 0x0e, 0x31, 0x25, 0x6a,
	0xe2, 0xb8, 0x8e, 0x83, 0xe8, 0xb7, 0xe1, 0xa3, 0xa2, 0xb8, 0x21, 0x2d, 0x08, 0xa6, 0xdc, 0xbb,
	0xbb, 0x3c, 0xff, 0xfc, 0x0c, 0x83, 0x72, 0x55, 0xa9, 0x15, 0xd3, 0xa6, 0xb0, 0x05, 0x06, 0x4e,
	0xe8, 0x38, 0x3c, 0x5f, 0xa6, 0xf6, 0xb5, 0x8a, 0xd9, 0x73, 0x91, 0x4f, 0xb2, 0x4a, 0x15, 0x13,
	0x23, 0x5f, 0x32, 0xf9, 0xde, 0x7c, 0x74, 0xdc, 0x14, 0xdb, 0xdf, 0x68, 0x00, 0xbd, 0x59, 0xae,
	0xed, 0x86, 0xde, 0x00, 0xdc, 0x4b, 0x7b, 0x27, 0xac, 0xe0, 0x72, 0x8d, 0xa7, 0xf0, 0xcf, 0x14,
	0x95, 0x4a, 0x9e, 0xd2, 0x64, 0x44, 0xc6, 0x24, 0xea, 0xf2, 0xc0, 0xe9, 0x87, 0x04, 0x8f, 0xa1,
	0xaf, 0x33, 0xb1, 0x91, 0x66, 0xe4, 0x8d, 0x49, 0xf4, 0x9f, 0x37, 0x8a, 0xde, 0xee, 0x19, 0x94,
	0x78, 0x06, 0xbe, 0x16, 0xc6, 0x8e, 0xc8, 0xb8, 0x1b, 0x0d, 0xa6, 0x07, 0xac, 0xa1, 0x63, 0x73,
	0x61, 0x2c, 0x77, 0x23, 0x44, 0xf0, 0x8d, 0x50, 0x2b, 0x67, 0xd3, 0xe3, 0xae, 0xa6, 0x8f, 0xe0,
	0xd7, 0x1b, 0x38, 0x04, 0xaf, 0x3d, 0xd9, 0x4b, 0x93, 0x6f, 0x3c, 0xde, 0x5f, 0x3c, 0xdd, 0x7d,
	0x9e, 0xda, 0xde, 0x11, 0xf8, 0x6e, 0xdd, 0xd5, 0xd3, 0x0f, 0x02, 0xbd, 0x45, 0x4d, 0x82, 0x11,
	0xf8, 0xf3, 0x54, 0x2d, 0x71, 0xd8, 0x92, 0xb9, 0x18, 0xc2, 0x1f, 0x9a, 0x76, 0xf0, 0x0a, 0x82,
	0xe6, 0x5e, 0x78, 0xd4, 0x0e, 0x77, 0x51, 0x85, 0xbf, 0x34, 0x4b, 0xda, 0xc1, 0x6b, 0xe8, 0x2f,
	0xac, 0x91, 0x22, 0xc7, 0x13, 0xf6, 0x15, 0x3d, 0xdb, 0x76, 0xb8, 0x5c, 0x57, 0xb2, 0xb4, 0xe1,
	0xe1, 0x6e, 0x30, 0x7b, 0x93, 0xca, 0xd2, 0xce, 0x05, 0x89, 0xfb, 0xee, 0x5d, 0x2e, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x23, 0x56, 0x11, 0x56, 0xdd, 0x01, 0x00, 0x00,
}

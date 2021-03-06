// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/webpush.proto

package webpush101

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type WebpushRequest struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebpushRequest) Reset()         { *m = WebpushRequest{} }
func (m *WebpushRequest) String() string { return proto.CompactTextString(m) }
func (*WebpushRequest) ProtoMessage()    {}
func (*WebpushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webpush_df7cff66462f91ce, []int{0}
}
func (m *WebpushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebpushRequest.Unmarshal(m, b)
}
func (m *WebpushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebpushRequest.Marshal(b, m, deterministic)
}
func (dst *WebpushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebpushRequest.Merge(dst, src)
}
func (m *WebpushRequest) XXX_Size() int {
	return xxx_messageInfo_WebpushRequest.Size(m)
}
func (m *WebpushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebpushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebpushRequest proto.InternalMessageInfo

func (m *WebpushRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type PushSubscriptionNotification struct {
	Subscription         []*PushSubscription `protobuf:"bytes,1,rep,name=subscription,proto3" json:"subscription,omitempty"`
	Request              *WebpushRequest     `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PushSubscriptionNotification) Reset()         { *m = PushSubscriptionNotification{} }
func (m *PushSubscriptionNotification) String() string { return proto.CompactTextString(m) }
func (*PushSubscriptionNotification) ProtoMessage()    {}
func (*PushSubscriptionNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_webpush_df7cff66462f91ce, []int{1}
}
func (m *PushSubscriptionNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushSubscriptionNotification.Unmarshal(m, b)
}
func (m *PushSubscriptionNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushSubscriptionNotification.Marshal(b, m, deterministic)
}
func (dst *PushSubscriptionNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushSubscriptionNotification.Merge(dst, src)
}
func (m *PushSubscriptionNotification) XXX_Size() int {
	return xxx_messageInfo_PushSubscriptionNotification.Size(m)
}
func (m *PushSubscriptionNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_PushSubscriptionNotification.DiscardUnknown(m)
}

var xxx_messageInfo_PushSubscriptionNotification proto.InternalMessageInfo

func (m *PushSubscriptionNotification) GetSubscription() []*PushSubscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *PushSubscriptionNotification) GetRequest() *WebpushRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*WebpushRequest)(nil), "WebpushRequest")
	proto.RegisterType((*PushSubscriptionNotification)(nil), "PushSubscriptionNotification")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WebpushServiceClient is the client API for WebpushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebpushServiceClient interface {
	SendPushSubscriptionNotification(ctx context.Context, in *PushSubscriptionNotification, opts ...grpc.CallOption) (*empty.Empty, error)
}

type webpushServiceClient struct {
	cc *grpc.ClientConn
}

func NewWebpushServiceClient(cc *grpc.ClientConn) WebpushServiceClient {
	return &webpushServiceClient{cc}
}

func (c *webpushServiceClient) SendPushSubscriptionNotification(ctx context.Context, in *PushSubscriptionNotification, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/WebpushService/SendPushSubscriptionNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebpushServiceServer is the server API for WebpushService service.
type WebpushServiceServer interface {
	SendPushSubscriptionNotification(context.Context, *PushSubscriptionNotification) (*empty.Empty, error)
}

func RegisterWebpushServiceServer(s *grpc.Server, srv WebpushServiceServer) {
	s.RegisterService(&_WebpushService_serviceDesc, srv)
}

func _WebpushService_SendPushSubscriptionNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushSubscriptionNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebpushServiceServer).SendPushSubscriptionNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WebpushService/SendPushSubscriptionNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebpushServiceServer).SendPushSubscriptionNotification(ctx, req.(*PushSubscriptionNotification))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebpushService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WebpushService",
	HandlerType: (*WebpushServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPushSubscriptionNotification",
			Handler:    _WebpushService_SendPushSubscriptionNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/webpush.proto",
}

func init() { proto.RegisterFile("protobuf/webpush.proto", fileDescriptor_webpush_df7cff66462f91ce) }

var fileDescriptor_webpush_df7cff66462f91ce = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4f, 0x4d, 0x2a, 0x28, 0x2d, 0xce, 0xd0, 0x03, 0x0b, 0x48,
	0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0xc3, 0xa5, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0xa1,
	0x92, 0x0a, 0x70, 0x51, 0x90, 0x8e, 0xf8, 0xe2, 0xd2, 0xa4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92,
	0xcc, 0xfc, 0x3c, 0x88, 0x0a, 0x25, 0x2d, 0x2e, 0xbe, 0x70, 0x88, 0x79, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0xf9, 0x79, 0x25, 0xa9, 0x79, 0x25, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x52, 0x03, 0x23, 0x97, 0x4c, 0x40, 0x69, 0x71, 0x46,
	0x30, 0x92, 0x31, 0x7e, 0xf9, 0x25, 0x99, 0x69, 0x99, 0xc9, 0x89, 0x20, 0xb6, 0x90, 0x29, 0x17,
	0x0f, 0xb2, 0x15, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x82, 0x7a, 0xe8, 0x9a, 0x82, 0x50,
	0x94, 0x09, 0x69, 0x72, 0xb1, 0x17, 0x41, 0x2c, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0xe2,
	0xd7, 0x43, 0x75, 0x53, 0x10, 0x4c, 0xde, 0x28, 0x1b, 0xee, 0xdc, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc,
	0xe4, 0x54, 0xa1, 0x48, 0x2e, 0x85, 0xe0, 0xd4, 0xbc, 0x14, 0xbc, 0xee, 0x92, 0xd5, 0xc3, 0x27,
	0x2d, 0x25, 0xa6, 0x07, 0x09, 0x43, 0x3d, 0x58, 0x68, 0xe9, 0xb9, 0x82, 0xc2, 0xd0, 0x49, 0x9e,
	0x4b, 0x30, 0x2f, 0x3f, 0x3b, 0x31, 0x37, 0xbf, 0x24, 0x1f, 0x2e, 0x15, 0xc5, 0x05, 0x0d, 0x7e,
	0x43, 0x03, 0xc3, 0x24, 0x36, 0xb0, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x3c, 0xa6,
	0xde, 0x9c, 0x01, 0x00, 0x00,
}

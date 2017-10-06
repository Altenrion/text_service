// Code generated by protoc-gen-go. DO NOT EDIT.
// source: text_service.proto

/*
Package text_service is a generated protocol buffer package.

It is generated from these files:
	text_service.proto

It has these top-level messages:
	Message
*/
package text_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Types
type Message struct {
	Sender string `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "text_service.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TextService service

type TextServiceClient interface {
	TransformMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type textServiceClient struct {
	cc *grpc.ClientConn
}

func NewTextServiceClient(cc *grpc.ClientConn) TextServiceClient {
	return &textServiceClient{cc}
}

func (c *textServiceClient) TransformMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/text_service.textService/TransformMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TextService service

type TextServiceServer interface {
	TransformMessage(context.Context, *Message) (*Message, error)
}

func RegisterTextServiceServer(s *grpc.Server, srv TextServiceServer) {
	s.RegisterService(&_TextService_serviceDesc, srv)
}

func _TextService_TransformMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextServiceServer).TransformMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/text_service.textService/TransformMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextServiceServer).TransformMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _TextService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "text_service.textService",
	HandlerType: (*TextServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransformMessage",
			Handler:    _TextService_TransformMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "text_service.proto",
}

func init() { proto.RegisterFile("text_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x49, 0xad, 0x28,
	0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x32, 0xe5, 0x62, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe3,
	0x62, 0x2b, 0x4e, 0xcd, 0x4b, 0x49, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2,
	0x84, 0x84, 0xb8, 0x58, 0x40, 0x5a, 0x24, 0x98, 0xc0, 0xa2, 0x60, 0xb6, 0x51, 0x20, 0x17, 0x37,
	0x88, 0x0e, 0x86, 0x98, 0x22, 0xe4, 0xc4, 0x25, 0x10, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x96, 0x5f,
	0x94, 0x0b, 0x33, 0x4e, 0x54, 0x0f, 0xc5, 0x72, 0xa8, 0xb0, 0x14, 0x76, 0x61, 0x25, 0x86, 0x24,
	0x36, 0xb0, 0xf3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xaf, 0xb6, 0x66, 0xb4, 0x00,
	0x00, 0x00,
}
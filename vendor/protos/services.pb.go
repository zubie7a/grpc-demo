// Code generated by protoc-gen-go.
// source: services.proto
// DO NOT EDIT!

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	services.proto

It has these top-level messages:
	CodeRequest
	CodeResponse
	StatusRequest
	StatusResponse
	Attendant
	AppData
*/
package services

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

// The request message for requesting a code, contains nothing.
type CodeRequest struct {
}

func (m *CodeRequest) Reset()                    { *m = CodeRequest{} }
func (m *CodeRequest) String() string            { return proto.CompactTextString(m) }
func (*CodeRequest) ProtoMessage()               {}
func (*CodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the code.
type CodeResponse struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *CodeResponse) Reset()                    { *m = CodeResponse{} }
func (m *CodeResponse) String() string            { return proto.CompactTextString(m) }
func (*CodeResponse) ProtoMessage()               {}
func (*CodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The request message for requesting the status.
type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// The response message containing the status.
type StatusResponse struct {
	AppData *AppData `protobuf:"bytes,1,opt,name=appData" json:"appData,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusResponse) GetAppData() *AppData {
	if m != nil {
		return m.AppData
	}
	return nil
}

// The struct that will define what an attendant is.
type Attendant struct {
	FirstName string `protobuf:"bytes,1,opt,name=firstName" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=lastName" json:"lastName,omitempty"`
	Code      string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *Attendant) Reset()                    { *m = Attendant{} }
func (m *Attendant) String() string            { return proto.CompactTextString(m) }
func (*Attendant) ProtoMessage()               {}
func (*Attendant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// The struct that will define the basic structure of the application data.
// It will be a counter of attendants, and a map between their unique codes
// and their name.
type AppData struct {
	Counter    int32             `protobuf:"varint,1,opt,name=counter" json:"counter,omitempty"`
	Attendants map[string]string `protobuf:"bytes,2,rep,name=attendants" json:"attendants,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AppData) Reset()                    { *m = AppData{} }
func (m *AppData) String() string            { return proto.CompactTextString(m) }
func (*AppData) ProtoMessage()               {}
func (*AppData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AppData) GetAttendants() map[string]string {
	if m != nil {
		return m.Attendants
	}
	return nil
}

func init() {
	proto.RegisterType((*CodeRequest)(nil), "services.CodeRequest")
	proto.RegisterType((*CodeResponse)(nil), "services.CodeResponse")
	proto.RegisterType((*StatusRequest)(nil), "services.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "services.StatusResponse")
	proto.RegisterType((*Attendant)(nil), "services.Attendant")
	proto.RegisterType((*AppData)(nil), "services.AppData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RegisterService service

type RegisterServiceClient interface {
	// Sends a greeting
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error)
}

type registerServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegisterServiceClient(cc *grpc.ClientConn) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/services.RegisterService/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerServiceClient) GetCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := grpc.Invoke(ctx, "/services.RegisterService/GetCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegisterService service

type RegisterServiceServer interface {
	// Sends a greeting
	GetStatus(context.Context, *StatusRequest) (*StatusResponse, error)
	GetCode(context.Context, *CodeRequest) (*CodeResponse, error)
}

func RegisterRegisterServiceServer(s *grpc.Server, srv RegisterServiceServer) {
	s.RegisterService(&_RegisterService_serviceDesc, srv)
}

func _RegisterService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.RegisterService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterService_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.RegisterService/GetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).GetCode(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _RegisterService_GetStatus_Handler,
		},
		{
			MethodName: "GetCode",
			Handler:    _RegisterService_GetCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("services.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0xc2, 0x8b, 0xa5, 0x83, 0x80, 0x4e, 0xfc, 0xd3, 0x34, 0x1e, 0xb0, 0x27, 0x12,
	0x13, 0x0e, 0x78, 0x31, 0x24, 0x24, 0x12, 0x35, 0xdc, 0x3c, 0x2c, 0x27, 0x8f, 0x2b, 0x8c, 0x86,
	0x88, 0xdd, 0xba, 0x3b, 0x25, 0xe1, 0x53, 0xf8, 0x29, 0xfc, 0x9e, 0x86, 0x6e, 0xb7, 0x54, 0xb8,
	0xed, 0x33, 0xcf, 0x3e, 0x9d, 0x5f, 0x67, 0x16, 0x3a, 0x86, 0xf4, 0x7a, 0x39, 0x27, 0x33, 0x48,
	0xb5, 0x62, 0x85, 0x4d, 0xa7, 0xe3, 0x36, 0xb4, 0x1e, 0xd4, 0x82, 0x04, 0x7d, 0x65, 0x64, 0x38,
	0x8e, 0xe1, 0xd8, 0x4a, 0x93, 0xaa, 0xc4, 0x10, 0x22, 0xfc, 0x9f, 0xab, 0x05, 0x85, 0x5e, 0xcf,
	0xeb, 0x07, 0x22, 0x3f, 0xc7, 0x5d, 0x68, 0xcf, 0x58, 0x72, 0x66, 0x5c, 0x68, 0x0c, 0x1d, 0x57,
	0x28, 0x62, 0x37, 0xe0, 0xcb, 0x34, 0x7d, 0x94, 0x2c, 0xf3, 0x64, 0x6b, 0x78, 0x3a, 0x28, 0x09,
	0x26, 0xd6, 0x10, 0xee, 0x46, 0xfc, 0x02, 0xc1, 0x84, 0x99, 0x92, 0x85, 0x4c, 0x18, 0xaf, 0x20,
	0x78, 0x5b, 0x6a, 0xc3, 0xcf, 0xf2, 0xd3, 0x75, 0xdd, 0x15, 0x30, 0x82, 0xe6, 0x4a, 0x16, 0x66,
	0x2d, 0x37, 0x4b, 0x5d, 0xa2, 0xd6, 0x2b, 0xa8, 0x3f, 0x1e, 0xf8, 0x45, 0x3f, 0x0c, 0xc1, 0x9f,
	0xab, 0x2c, 0x61, 0xd2, 0xf9, 0x77, 0x1b, 0xc2, 0x49, 0x9c, 0x00, 0x48, 0x07, 0x60, 0xc2, 0x5a,
	0xaf, 0xde, 0x6f, 0x0d, 0xaf, 0x0f, 0x80, 0x07, 0x25, 0xa4, 0x79, 0x4a, 0x58, 0x6f, 0x44, 0x25,
	0x14, 0x8d, 0xa1, 0xbb, 0x67, 0xe3, 0x09, 0xd4, 0x3f, 0x68, 0x53, 0xfc, 0xc3, 0xf6, 0x88, 0x67,
	0xd0, 0x58, 0xcb, 0x55, 0xe6, 0xd0, 0xad, 0x18, 0xd5, 0xee, 0xbc, 0xe1, 0xb7, 0x07, 0x5d, 0x41,
	0xef, 0x4b, 0xc3, 0xa4, 0x67, 0xb6, 0x2f, 0xde, 0x43, 0x30, 0x25, 0xb6, 0x83, 0xc5, 0xcb, 0x1d,
	0xce, 0x9f, 0xd9, 0x47, 0xe1, 0xa1, 0x61, 0x77, 0x10, 0xff, 0xc3, 0x11, 0xf8, 0x53, 0xe2, 0xed,
	0x3e, 0xf1, 0x7c, 0x77, 0xad, 0xb2, 0xee, 0xe8, 0x62, 0xbf, 0xec, 0xb2, 0xaf, 0x47, 0xf9, 0x43,
	0xb9, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x16, 0xc1, 0xfb, 0xea, 0x3a, 0x02, 0x00, 0x00,
}

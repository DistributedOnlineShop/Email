// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: service_verification.proto

package verification

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Verification_CheckEmail_FullMethodName = "/pb.Verification/CheckEmail"
)

// VerificationClient is the client API for Verification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerificationClient interface {
	CheckEmail(ctx context.Context, in *CheckEmailRequest, opts ...grpc.CallOption) (*CheckEmailResponse, error)
}

type verificationClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificationClient(cc grpc.ClientConnInterface) VerificationClient {
	return &verificationClient{cc}
}

func (c *verificationClient) CheckEmail(ctx context.Context, in *CheckEmailRequest, opts ...grpc.CallOption) (*CheckEmailResponse, error) {
	out := new(CheckEmailResponse)
	err := c.cc.Invoke(ctx, Verification_CheckEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationServer is the server API for Verification service.
// All implementations must embed UnimplementedVerificationServer
// for forward compatibility
type VerificationServer interface {
	CheckEmail(context.Context, *CheckEmailRequest) (*CheckEmailResponse, error)
	mustEmbedUnimplementedVerificationServer()
}

// UnimplementedVerificationServer must be embedded to have forward compatible implementations.
type UnimplementedVerificationServer struct {
}

func (UnimplementedVerificationServer) CheckEmail(context.Context, *CheckEmailRequest) (*CheckEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmail not implemented")
}
func (UnimplementedVerificationServer) mustEmbedUnimplementedVerificationServer() {}

// UnsafeVerificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerificationServer will
// result in compilation errors.
type UnsafeVerificationServer interface {
	mustEmbedUnimplementedVerificationServer()
}

func RegisterVerificationServer(s grpc.ServiceRegistrar, srv VerificationServer) {
	s.RegisterService(&Verification_ServiceDesc, srv)
}

func _Verification_CheckEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServer).CheckEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verification_CheckEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServer).CheckEmail(ctx, req.(*CheckEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Verification_ServiceDesc is the grpc.ServiceDesc for Verification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Verification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Verification",
	HandlerType: (*VerificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckEmail",
			Handler:    _Verification_CheckEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_verification.proto",
}

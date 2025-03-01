// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0--rc1
// source: payment-service/api/payment.proto

package paymentpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_Charge_FullMethodName               = "/payment.PaymentService/Charge"
	PaymentService_GetTransactionStatus_FullMethodName = "/payment.PaymentService/GetTransactionStatus"
	PaymentService_HealthCheck_FullMethodName          = "/payment.PaymentService/HealthCheck"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	Charge(ctx context.Context, in *ChargeReq, opts ...grpc.CallOption) (*ChargeResp, error)
	GetTransactionStatus(ctx context.Context, in *GetTransactionStatusReq, opts ...grpc.CallOption) (*GetTransactionStatusResp, error)
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Charge(ctx context.Context, in *ChargeReq, opts ...grpc.CallOption) (*ChargeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChargeResp)
	err := c.cc.Invoke(ctx, PaymentService_Charge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetTransactionStatus(ctx context.Context, in *GetTransactionStatusReq, opts ...grpc.CallOption) (*GetTransactionStatusResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionStatusResp)
	err := c.cc.Invoke(ctx, PaymentService_GetTransactionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, PaymentService_HealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
type PaymentServiceServer interface {
	Charge(context.Context, *ChargeReq) (*ChargeResp, error)
	GetTransactionStatus(context.Context, *GetTransactionStatusReq) (*GetTransactionStatusResp, error)
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) Charge(context.Context, *ChargeReq) (*ChargeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Charge not implemented")
}
func (UnimplementedPaymentServiceServer) GetTransactionStatus(context.Context, *GetTransactionStatusReq) (*GetTransactionStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionStatus not implemented")
}
func (UnimplementedPaymentServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_Charge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Charge(ctx, req.(*ChargeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetTransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetTransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetTransactionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetTransactionStatus(ctx, req.(*GetTransactionStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PaymentService_Charge_Handler,
		},
		{
			MethodName: "GetTransactionStatus",
			Handler:    _PaymentService_GetTransactionStatus_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _PaymentService_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment-service/api/payment.proto",
}

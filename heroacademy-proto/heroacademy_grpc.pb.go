// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: heroacademy-proto.proto

package heroacademy_proto

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
	TrainingService_TrainHero_FullMethodName = "/heroacademy-proto.TrainingService/TrainHero"
)

// TrainingServiceClient is the heroacademy-client API for TrainingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainingServiceClient interface {
	TrainHero(ctx context.Context, in *HeroRequest, opts ...grpc.CallOption) (*HeroResponse, error)
}

type trainingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainingServiceClient(cc grpc.ClientConnInterface) TrainingServiceClient {
	return &trainingServiceClient{cc}
}

func (c *trainingServiceClient) TrainHero(ctx context.Context, in *HeroRequest, opts ...grpc.CallOption) (*HeroResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeroResponse)
	err := c.cc.Invoke(ctx, TrainingService_TrainHero_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainingServiceServer is the heroacademy-server API for TrainingService service.
// All implementations must embed UnimplementedTrainingServiceServer
// for forward compatibility.
type TrainingServiceServer interface {
	TrainHero(context.Context, *HeroRequest) (*HeroResponse, error)
	mustEmbedUnimplementedTrainingServiceServer()
}

// UnimplementedTrainingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrainingServiceServer struct{}

func (UnimplementedTrainingServiceServer) TrainHero(context.Context, *HeroRequest) (*HeroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainHero not implemented")
}
func (UnimplementedTrainingServiceServer) mustEmbedUnimplementedTrainingServiceServer() {}
func (UnimplementedTrainingServiceServer) testEmbeddedByValue()                         {}

// UnsafeTrainingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainingServiceServer will
// result in compilation errors.
type UnsafeTrainingServiceServer interface {
	mustEmbedUnimplementedTrainingServiceServer()
}

func RegisterTrainingServiceServer(s grpc.ServiceRegistrar, srv TrainingServiceServer) {
	// If the following call pancis, it indicates UnimplementedTrainingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrainingService_ServiceDesc, srv)
}

func _TrainingService_TrainHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).TrainHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_TrainHero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).TrainHero(ctx, req.(*HeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainingService_ServiceDesc is the grpc.ServiceDesc for TrainingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heroacademy-proto.TrainingService",
	HandlerType: (*TrainingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrainHero",
			Handler:    _TrainingService_TrainHero_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heroacademy-proto.proto",
}

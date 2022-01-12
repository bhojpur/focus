// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// FocusUIClient is the client API for FocusUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FocusUIClient interface {
	// ListEngineSpecs returns a list of Orientation Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (FocusUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type focusUIClient struct {
	cc grpc.ClientConnInterface
}

func NewFocusUIClient(cc grpc.ClientConnInterface) FocusUIClient {
	return &focusUIClient{cc}
}

func (c *focusUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (FocusUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FocusUI_ServiceDesc.Streams[0], "/v1.FocusUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &focusUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FocusUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type focusUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *focusUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *focusUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.FocusUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FocusUIServer is the server API for FocusUI service.
// All implementations must embed UnimplementedFocusUIServer
// for forward compatibility
type FocusUIServer interface {
	// ListEngineSpecs returns a list of Orientation Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, FocusUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedFocusUIServer()
}

// UnimplementedFocusUIServer must be embedded to have forward compatible implementations.
type UnimplementedFocusUIServer struct {
}

func (UnimplementedFocusUIServer) ListEngineSpecs(*ListEngineSpecsRequest, FocusUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedFocusUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedFocusUIServer) mustEmbedUnimplementedFocusUIServer() {}

// UnsafeFocusUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FocusUIServer will
// result in compilation errors.
type UnsafeFocusUIServer interface {
	mustEmbedUnimplementedFocusUIServer()
}

func RegisterFocusUIServer(s grpc.ServiceRegistrar, srv FocusUIServer) {
	s.RegisterService(&FocusUI_ServiceDesc, srv)
}

func _FocusUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FocusUIServer).ListEngineSpecs(m, &focusUIListEngineSpecsServer{stream})
}

type FocusUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type focusUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *focusUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FocusUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FocusUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.FocusUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FocusUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FocusUI_ServiceDesc is the grpc.ServiceDesc for FocusUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FocusUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FocusUI",
	HandlerType: (*FocusUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _FocusUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _FocusUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "focus-ui.proto",
}

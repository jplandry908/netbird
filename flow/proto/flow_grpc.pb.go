// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FlowServiceClient is the client API for FlowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlowServiceClient interface {
	// Client to receiver streams of events and acknowledgements
	Events(ctx context.Context, opts ...grpc.CallOption) (FlowService_EventsClient, error)
}

type flowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowServiceClient(cc grpc.ClientConnInterface) FlowServiceClient {
	return &flowServiceClient{cc}
}

func (c *flowServiceClient) Events(ctx context.Context, opts ...grpc.CallOption) (FlowService_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlowService_ServiceDesc.Streams[0], "/flow.FlowService/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceEventsClient{stream}
	return x, nil
}

type FlowService_EventsClient interface {
	Send(*FlowEvent) error
	Recv() (*FlowEventAck, error)
	grpc.ClientStream
}

type flowServiceEventsClient struct {
	grpc.ClientStream
}

func (x *flowServiceEventsClient) Send(m *FlowEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceEventsClient) Recv() (*FlowEventAck, error) {
	m := new(FlowEventAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlowServiceServer is the server API for FlowService service.
// All implementations must embed UnimplementedFlowServiceServer
// for forward compatibility
type FlowServiceServer interface {
	// Client to receiver streams of events and acknowledgements
	Events(FlowService_EventsServer) error
	mustEmbedUnimplementedFlowServiceServer()
}

// UnimplementedFlowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlowServiceServer struct {
}

func (UnimplementedFlowServiceServer) Events(FlowService_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}
func (UnimplementedFlowServiceServer) mustEmbedUnimplementedFlowServiceServer() {}

// UnsafeFlowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlowServiceServer will
// result in compilation errors.
type UnsafeFlowServiceServer interface {
	mustEmbedUnimplementedFlowServiceServer()
}

func RegisterFlowServiceServer(s grpc.ServiceRegistrar, srv FlowServiceServer) {
	s.RegisterService(&FlowService_ServiceDesc, srv)
}

func _FlowService_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).Events(&flowServiceEventsServer{stream})
}

type FlowService_EventsServer interface {
	Send(*FlowEventAck) error
	Recv() (*FlowEvent, error)
	grpc.ServerStream
}

type flowServiceEventsServer struct {
	grpc.ServerStream
}

func (x *flowServiceEventsServer) Send(m *FlowEventAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceEventsServer) Recv() (*FlowEvent, error) {
	m := new(FlowEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlowService_ServiceDesc is the grpc.ServiceDesc for FlowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.FlowService",
	HandlerType: (*FlowServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _FlowService_Events_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "flow.proto",
}

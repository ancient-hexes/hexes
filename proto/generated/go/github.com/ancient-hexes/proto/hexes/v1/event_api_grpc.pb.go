// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: hexes/v1/event_api.proto

package hexesv1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventAPIClient is the client API for EventAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventAPIClient interface {
}

type eventAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEventAPIClient(cc grpc.ClientConnInterface) EventAPIClient {
	return &eventAPIClient{cc}
}

// EventAPIServer is the server API for EventAPI service.
// All implementations must embed UnimplementedEventAPIServer
// for forward compatibility
type EventAPIServer interface {
	mustEmbedUnimplementedEventAPIServer()
}

// UnimplementedEventAPIServer must be embedded to have forward compatible implementations.
type UnimplementedEventAPIServer struct {
}

func (UnimplementedEventAPIServer) mustEmbedUnimplementedEventAPIServer() {}

// UnsafeEventAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventAPIServer will
// result in compilation errors.
type UnsafeEventAPIServer interface {
	mustEmbedUnimplementedEventAPIServer()
}

func RegisterEventAPIServer(s grpc.ServiceRegistrar, srv EventAPIServer) {
	s.RegisterService(&EventAPI_ServiceDesc, srv)
}

// EventAPI_ServiceDesc is the grpc.ServiceDesc for EventAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hexes.v1.EventAPI",
	HandlerType: (*EventAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "hexes/v1/event_api.proto",
}

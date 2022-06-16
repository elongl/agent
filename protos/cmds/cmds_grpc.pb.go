// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: cmds.proto

package cmds

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

// AgentManagerClient is the client API for AgentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentManagerClient interface {
	RunEchoCommand(ctx context.Context, in *EchoCommandRequest, opts ...grpc.CallOption) (*EchoCommandResponse, error)
	RunShellCommand(ctx context.Context, in *ShellCommandRequest, opts ...grpc.CallOption) (*ShellCommandResponse, error)
	GetAgents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AgentManager_GetAgentsClient, error)
}

type agentManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentManagerClient(cc grpc.ClientConnInterface) AgentManagerClient {
	return &agentManagerClient{cc}
}

func (c *agentManagerClient) RunEchoCommand(ctx context.Context, in *EchoCommandRequest, opts ...grpc.CallOption) (*EchoCommandResponse, error) {
	out := new(EchoCommandResponse)
	err := c.cc.Invoke(ctx, "/AgentManager/RunEchoCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentManagerClient) RunShellCommand(ctx context.Context, in *ShellCommandRequest, opts ...grpc.CallOption) (*ShellCommandResponse, error) {
	out := new(ShellCommandResponse)
	err := c.cc.Invoke(ctx, "/AgentManager/RunShellCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentManagerClient) GetAgents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AgentManager_GetAgentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentManager_ServiceDesc.Streams[0], "/AgentManager/GetAgents", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentManagerGetAgentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentManager_GetAgentsClient interface {
	Recv() (*AgentInfo, error)
	grpc.ClientStream
}

type agentManagerGetAgentsClient struct {
	grpc.ClientStream
}

func (x *agentManagerGetAgentsClient) Recv() (*AgentInfo, error) {
	m := new(AgentInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentManagerServer is the server API for AgentManager service.
// All implementations must embed UnimplementedAgentManagerServer
// for forward compatibility
type AgentManagerServer interface {
	RunEchoCommand(context.Context, *EchoCommandRequest) (*EchoCommandResponse, error)
	RunShellCommand(context.Context, *ShellCommandRequest) (*ShellCommandResponse, error)
	GetAgents(*Empty, AgentManager_GetAgentsServer) error
	mustEmbedUnimplementedAgentManagerServer()
}

// UnimplementedAgentManagerServer must be embedded to have forward compatible implementations.
type UnimplementedAgentManagerServer struct {
}

func (UnimplementedAgentManagerServer) RunEchoCommand(context.Context, *EchoCommandRequest) (*EchoCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunEchoCommand not implemented")
}
func (UnimplementedAgentManagerServer) RunShellCommand(context.Context, *ShellCommandRequest) (*ShellCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunShellCommand not implemented")
}
func (UnimplementedAgentManagerServer) GetAgents(*Empty, AgentManager_GetAgentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAgents not implemented")
}
func (UnimplementedAgentManagerServer) mustEmbedUnimplementedAgentManagerServer() {}

// UnsafeAgentManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentManagerServer will
// result in compilation errors.
type UnsafeAgentManagerServer interface {
	mustEmbedUnimplementedAgentManagerServer()
}

func RegisterAgentManagerServer(s grpc.ServiceRegistrar, srv AgentManagerServer) {
	s.RegisterService(&AgentManager_ServiceDesc, srv)
}

func _AgentManager_RunEchoCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentManagerServer).RunEchoCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgentManager/RunEchoCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentManagerServer).RunEchoCommand(ctx, req.(*EchoCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentManager_RunShellCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentManagerServer).RunShellCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgentManager/RunShellCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentManagerServer).RunShellCommand(ctx, req.(*ShellCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentManager_GetAgents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentManagerServer).GetAgents(m, &agentManagerGetAgentsServer{stream})
}

type AgentManager_GetAgentsServer interface {
	Send(*AgentInfo) error
	grpc.ServerStream
}

type agentManagerGetAgentsServer struct {
	grpc.ServerStream
}

func (x *agentManagerGetAgentsServer) Send(m *AgentInfo) error {
	return x.ServerStream.SendMsg(m)
}

// AgentManager_ServiceDesc is the grpc.ServiceDesc for AgentManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AgentManager",
	HandlerType: (*AgentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunEchoCommand",
			Handler:    _AgentManager_RunEchoCommand_Handler,
		},
		{
			MethodName: "RunShellCommand",
			Handler:    _AgentManager_RunShellCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAgents",
			Handler:       _AgentManager_GetAgents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cmds.proto",
}

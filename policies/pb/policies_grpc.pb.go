// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// PolicyServiceClient is the client API for PolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyServiceClient interface {
	RetrievePolicy(ctx context.Context, in *PolicyByIDReq, opts ...grpc.CallOption) (*PolicyRes, error)
	RetrievePoliciesByGroups(ctx context.Context, in *PoliciesByGroupsReq, opts ...grpc.CallOption) (*PolicyInDSListRes, error)
}

type policyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyServiceClient(cc grpc.ClientConnInterface) PolicyServiceClient {
	return &policyServiceClient{cc}
}

func (c *policyServiceClient) RetrievePolicy(ctx context.Context, in *PolicyByIDReq, opts ...grpc.CallOption) (*PolicyRes, error) {
	out := new(PolicyRes)
	err := c.cc.Invoke(ctx, "/policies.PolicyService/RetrievePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) RetrievePoliciesByGroups(ctx context.Context, in *PoliciesByGroupsReq, opts ...grpc.CallOption) (*PolicyInDSListRes, error) {
	out := new(PolicyInDSListRes)
	err := c.cc.Invoke(ctx, "/policies.PolicyService/RetrievePoliciesByGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServiceServer is the server API for PolicyService service.
// All implementations must embed UnimplementedPolicyServiceServer
// for forward compatibility
type PolicyServiceServer interface {
	RetrievePolicy(context.Context, *PolicyByIDReq) (*PolicyRes, error)
	RetrievePoliciesByGroups(context.Context, *PoliciesByGroupsReq) (*PolicyInDSListRes, error)
	mustEmbedUnimplementedPolicyServiceServer()
}

// UnimplementedPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyServiceServer struct {
}

func (UnimplementedPolicyServiceServer) RetrievePolicy(context.Context, *PolicyByIDReq) (*PolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrievePolicy not implemented")
}
func (UnimplementedPolicyServiceServer) RetrievePoliciesByGroups(context.Context, *PoliciesByGroupsReq) (*PolicyInDSListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrievePoliciesByGroups not implemented")
}
func (UnimplementedPolicyServiceServer) mustEmbedUnimplementedPolicyServiceServer() {}

// UnsafePolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServiceServer will
// result in compilation errors.
type UnsafePolicyServiceServer interface {
	mustEmbedUnimplementedPolicyServiceServer()
}

func RegisterPolicyServiceServer(s grpc.ServiceRegistrar, srv PolicyServiceServer) {
	s.RegisterService(&PolicyService_ServiceDesc, srv)
}

func _PolicyService_RetrievePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).RetrievePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policies.PolicyService/RetrievePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).RetrievePolicy(ctx, req.(*PolicyByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_RetrievePoliciesByGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoliciesByGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).RetrievePoliciesByGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policies.PolicyService/RetrievePoliciesByGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).RetrievePoliciesByGroups(ctx, req.(*PoliciesByGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyService_ServiceDesc is the grpc.ServiceDesc for PolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "policies.PolicyService",
	HandlerType: (*PolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrievePolicy",
			Handler:    _PolicyService_RetrievePolicy_Handler,
		},
		{
			MethodName: "RetrievePoliciesByGroups",
			Handler:    _PolicyService_RetrievePoliciesByGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policies/pb/policies.proto",
}

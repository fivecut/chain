// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: band/feeds/v1beta1/tx.proto

package feedsv1beta1

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
	Msg_VoteSignals_FullMethodName                 = "/band.feeds.v1beta1.Msg/VoteSignals"
	Msg_SubmitSignalPrices_FullMethodName          = "/band.feeds.v1beta1.Msg/SubmitSignalPrices"
	Msg_UpdateReferenceSourceConfig_FullMethodName = "/band.feeds.v1beta1.Msg/UpdateReferenceSourceConfig"
	Msg_UpdateParams_FullMethodName                = "/band.feeds.v1beta1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// VoteSignals is an RPC method to vote signal ids and their powers.
	VoteSignals(ctx context.Context, in *MsgVoteSignals, opts ...grpc.CallOption) (*MsgVoteSignalsResponse, error)
	// SubmitSignalPrices is an RPC method to submit signal prices.
	SubmitSignalPrices(ctx context.Context, in *MsgSubmitSignalPrices, opts ...grpc.CallOption) (*MsgSubmitSignalPricesResponse, error)
	// UpdateReferenceSourceConfig is an RPC method to update reference price source configuration.
	UpdateReferenceSourceConfig(ctx context.Context, in *MsgUpdateReferenceSourceConfig, opts ...grpc.CallOption) (*MsgUpdateReferenceSourceConfigResponse, error)
	// UpdateParams is an RPC method to update parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) VoteSignals(ctx context.Context, in *MsgVoteSignals, opts ...grpc.CallOption) (*MsgVoteSignalsResponse, error) {
	out := new(MsgVoteSignalsResponse)
	err := c.cc.Invoke(ctx, Msg_VoteSignals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitSignalPrices(ctx context.Context, in *MsgSubmitSignalPrices, opts ...grpc.CallOption) (*MsgSubmitSignalPricesResponse, error) {
	out := new(MsgSubmitSignalPricesResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitSignalPrices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateReferenceSourceConfig(ctx context.Context, in *MsgUpdateReferenceSourceConfig, opts ...grpc.CallOption) (*MsgUpdateReferenceSourceConfigResponse, error) {
	out := new(MsgUpdateReferenceSourceConfigResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateReferenceSourceConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// VoteSignals is an RPC method to vote signal ids and their powers.
	VoteSignals(context.Context, *MsgVoteSignals) (*MsgVoteSignalsResponse, error)
	// SubmitSignalPrices is an RPC method to submit signal prices.
	SubmitSignalPrices(context.Context, *MsgSubmitSignalPrices) (*MsgSubmitSignalPricesResponse, error)
	// UpdateReferenceSourceConfig is an RPC method to update reference price source configuration.
	UpdateReferenceSourceConfig(context.Context, *MsgUpdateReferenceSourceConfig) (*MsgUpdateReferenceSourceConfigResponse, error)
	// UpdateParams is an RPC method to update parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) VoteSignals(context.Context, *MsgVoteSignals) (*MsgVoteSignalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteSignals not implemented")
}
func (UnimplementedMsgServer) SubmitSignalPrices(context.Context, *MsgSubmitSignalPrices) (*MsgSubmitSignalPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSignalPrices not implemented")
}
func (UnimplementedMsgServer) UpdateReferenceSourceConfig(context.Context, *MsgUpdateReferenceSourceConfig) (*MsgUpdateReferenceSourceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReferenceSourceConfig not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_VoteSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteSignals)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_VoteSignals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteSignals(ctx, req.(*MsgVoteSignals))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitSignalPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitSignalPrices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitSignalPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitSignalPrices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitSignalPrices(ctx, req.(*MsgSubmitSignalPrices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateReferenceSourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateReferenceSourceConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateReferenceSourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateReferenceSourceConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateReferenceSourceConfig(ctx, req.(*MsgUpdateReferenceSourceConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "band.feeds.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VoteSignals",
			Handler:    _Msg_VoteSignals_Handler,
		},
		{
			MethodName: "SubmitSignalPrices",
			Handler:    _Msg_SubmitSignalPrices_Handler,
		},
		{
			MethodName: "UpdateReferenceSourceConfig",
			Handler:    _Msg_UpdateReferenceSourceConfig_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "band/feeds/v1beta1/tx.proto",
}

//
// The MIT License (MIT)
//
// Copyright (c) 2021 HereweTech Co.LTD
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/list.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	List_InitDB_FullMethodName = "/svc.sicky.list.List/InitDB"
	List_Add_FullMethodName    = "/svc.sicky.list.List/Add"
	List_Get_FullMethodName    = "/svc.sicky.list.List/Get"
	List_Set_FullMethodName    = "/svc.sicky.list.List/Set"
	List_Delete_FullMethodName = "/svc.sicky.list.List/Delete"
	List_List_FullMethodName   = "/svc.sicky.list.List/List"
	List_All_FullMethodName    = "/svc.sicky.list.List/All"
	List_Count_FullMethodName  = "/svc.sicky.list.List/Count"
	List_Purge_FullMethodName  = "/svc.sicky.list.List/Purge"
)

// ListClient is the client API for List service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
	Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
	All(ctx context.Context, in *AllReq, opts ...grpc.CallOption) (*AllResp, error)
	Count(ctx context.Context, in *CountReq, opts ...grpc.CallOption) (*CountResp, error)
	Purge(ctx context.Context, in *PurgeReq, opts ...grpc.CallOption) (*PurgeResp, error)
}

type listClient struct {
	cc grpc.ClientConnInterface
}

func NewListClient(cc grpc.ClientConnInterface) ListClient {
	return &listClient{cc}
}

func (c *listClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, List_InitDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddResp)
	err := c.cc.Invoke(ctx, List_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResp)
	err := c.cc.Invoke(ctx, List_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetResp)
	err := c.cc.Invoke(ctx, List_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, List_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResp)
	err := c.cc.Invoke(ctx, List_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) All(ctx context.Context, in *AllReq, opts ...grpc.CallOption) (*AllResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllResp)
	err := c.cc.Invoke(ctx, List_All_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) Count(ctx context.Context, in *CountReq, opts ...grpc.CallOption) (*CountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountResp)
	err := c.cc.Invoke(ctx, List_Count_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) Purge(ctx context.Context, in *PurgeReq, opts ...grpc.CallOption) (*PurgeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PurgeResp)
	err := c.cc.Invoke(ctx, List_Purge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServer is the server API for List service.
// All implementations must embed UnimplementedListServer
// for forward compatibility.
type ListServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	Add(context.Context, *AddReq) (*AddResp, error)
	Get(context.Context, *GetReq) (*GetResp, error)
	Set(context.Context, *SetReq) (*SetResp, error)
	Delete(context.Context, *DeleteReq) (*DeleteResp, error)
	List(context.Context, *ListReq) (*ListResp, error)
	All(context.Context, *AllReq) (*AllResp, error)
	Count(context.Context, *CountReq) (*CountResp, error)
	Purge(context.Context, *PurgeReq) (*PurgeResp, error)
	mustEmbedUnimplementedListServer()
}

// UnimplementedListServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListServer struct{}

func (UnimplementedListServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedListServer) Add(context.Context, *AddReq) (*AddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedListServer) Get(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedListServer) Set(context.Context, *SetReq) (*SetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedListServer) Delete(context.Context, *DeleteReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedListServer) List(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedListServer) All(context.Context, *AllReq) (*AllResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedListServer) Count(context.Context, *CountReq) (*CountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedListServer) Purge(context.Context, *PurgeReq) (*PurgeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}
func (UnimplementedListServer) mustEmbedUnimplementedListServer() {}
func (UnimplementedListServer) testEmbeddedByValue()              {}

// UnsafeListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListServer will
// result in compilation errors.
type UnsafeListServer interface {
	mustEmbedUnimplementedListServer()
}

func RegisterListServer(s grpc.ServiceRegistrar, srv ListServer) {
	// If the following call pancis, it indicates UnimplementedListServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&List_ServiceDesc, srv)
}

func _List_InitDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).InitDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_InitDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).InitDB(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).Add(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).Set(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).All(ctx, req.(*AllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_Count_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).Count(ctx, req.(*CountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_Purge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).Purge(ctx, req.(*PurgeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// List_ServiceDesc is the grpc.ServiceDesc for List service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var List_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.sicky.list.List",
	HandlerType: (*ListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDB",
			Handler:    _List_InitDB_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _List_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _List_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _List_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _List_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _List_List_Handler,
		},
		{
			MethodName: "All",
			Handler:    _List_All_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _List_Count_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _List_Purge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/list.proto",
}

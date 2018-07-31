// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/admin.proto

package service // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/service"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import admin "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error)
	GetTask(ctx context.Context, in *admin.GetObjectRequest, opts ...grpc.CallOption) (*admin.Task, error)
	ListTaskIds(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error)
	ListTasks(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.TaskList, error)
	CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error)
	GetWorkflow(ctx context.Context, in *admin.GetObjectRequest, opts ...grpc.CallOption) (*admin.Workflow, error)
	ListWorkflowIds(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error)
	ListWorkflows(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.WorkflowList, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error) {
	out := new(admin.TaskCreateResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetTask(ctx context.Context, in *admin.GetObjectRequest, opts ...grpc.CallOption) (*admin.Task, error) {
	out := new(admin.Task)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListTaskIds(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error) {
	out := new(admin.IdentifierList)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/ListTaskIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListTasks(ctx context.Context, in *admin.TaskListRequest, opts ...grpc.CallOption) (*admin.TaskList, error) {
	out := new(admin.TaskList)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error) {
	out := new(admin.WorkflowCreateResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetWorkflow(ctx context.Context, in *admin.GetObjectRequest, opts ...grpc.CallOption) (*admin.Workflow, error) {
	out := new(admin.Workflow)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/GetWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListWorkflowIds(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.IdentifierList, error) {
	out := new(admin.IdentifierList)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/ListWorkflowIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListWorkflows(ctx context.Context, in *admin.WorkflowListRequest, opts ...grpc.CallOption) (*admin.WorkflowList, error) {
	out := new(admin.WorkflowList)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AdminService/ListWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	CreateTask(context.Context, *admin.TaskCreateRequest) (*admin.TaskCreateResponse, error)
	GetTask(context.Context, *admin.GetObjectRequest) (*admin.Task, error)
	ListTaskIds(context.Context, *admin.TaskListRequest) (*admin.IdentifierList, error)
	ListTasks(context.Context, *admin.TaskListRequest) (*admin.TaskList, error)
	CreateWorkflow(context.Context, *admin.WorkflowCreateRequest) (*admin.WorkflowCreateResponse, error)
	GetWorkflow(context.Context, *admin.GetObjectRequest) (*admin.Workflow, error)
	ListWorkflowIds(context.Context, *admin.WorkflowListRequest) (*admin.IdentifierList, error)
	ListWorkflows(context.Context, *admin.WorkflowListRequest) (*admin.WorkflowList, error)
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateTask(ctx, req.(*admin.TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetTask(ctx, req.(*admin.GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListTaskIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListTaskIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/ListTaskIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListTaskIds(ctx, req.(*admin.TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListTasks(ctx, req.(*admin.TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, req.(*admin.WorkflowCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/GetWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetWorkflow(ctx, req.(*admin.GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListWorkflowIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListWorkflowIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/ListWorkflowIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListWorkflowIds(ctx, req.(*admin.WorkflowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AdminService/ListWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListWorkflows(ctx, req.(*admin.WorkflowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _AdminService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _AdminService_GetTask_Handler,
		},
		{
			MethodName: "ListTaskIds",
			Handler:    _AdminService_ListTaskIds_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _AdminService_ListTasks_Handler,
		},
		{
			MethodName: "CreateWorkflow",
			Handler:    _AdminService_CreateWorkflow_Handler,
		},
		{
			MethodName: "GetWorkflow",
			Handler:    _AdminService_GetWorkflow_Handler,
		},
		{
			MethodName: "ListWorkflowIds",
			Handler:    _AdminService_ListWorkflowIds_Handler,
		},
		{
			MethodName: "ListWorkflows",
			Handler:    _AdminService_ListWorkflows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/admin.proto",
}

func init() { proto.RegisterFile("flyteidl/service/admin.proto", fileDescriptor_admin_05d6a593360774a8) }

var fileDescriptor_admin_05d6a593360774a8 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x8b, 0x13, 0x31,
	0x1c, 0xc6, 0x19, 0x0f, 0xbe, 0x64, 0x5d, 0x5f, 0xfe, 0xbe, 0x75, 0xbb, 0xd5, 0xd5, 0x88, 0x1e,
	0x04, 0x1b, 0xdc, 0xaa, 0x48, 0x11, 0x41, 0x3d, 0x94, 0x05, 0x41, 0x50, 0x41, 0xe8, 0x41, 0x49,
	0x3b, 0xff, 0x8e, 0xd9, 0xce, 0x24, 0xb3, 0x93, 0x74, 0x97, 0x52, 0xe6, 0xe2, 0xd1, 0xab, 0xe0,
	0x77, 0xf1, 0x43, 0x78, 0xf2, 0x2b, 0xf8, 0x29, 0x3c, 0xc9, 0x64, 0x92, 0xc1, 0x1d, 0xb7, 0xbb,
	0x3b, 0x9e, 0x5a, 0xf2, 0x3c, 0xf9, 0x3f, 0xbf, 0x4c, 0x1e, 0x42, 0x3a, 0x93, 0x78, 0x6e, 0x50,
	0x84, 0x31, 0xd3, 0x98, 0xed, 0x8a, 0x31, 0x32, 0x1e, 0x26, 0x42, 0x76, 0xd3, 0x4c, 0x19, 0x05,
	0x17, 0xbc, 0xda, 0x75, 0x6a, 0xbb, 0x13, 0x29, 0x15, 0xc5, 0xc8, 0x78, 0x2a, 0x18, 0x97, 0x52,
	0x19, 0x6e, 0x84, 0x92, 0xba, 0xf4, 0xb7, 0xd7, 0xaa, 0x69, 0x76, 0x0a, 0x33, 0x5c, 0x4f, 0x9d,
	0x74, 0xbd, 0x26, 0xed, 0xa9, 0x6c, 0x3a, 0x89, 0xd5, 0x9e, 0x93, 0xd7, 0x6b, 0xf2, 0x58, 0x25,
	0x89, 0x72, 0x18, 0x9b, 0xbf, 0x4f, 0x93, 0xb3, 0xcf, 0x8b, 0xe5, 0xb7, 0x25, 0x05, 0x24, 0x84,
	0xbc, 0xcc, 0x90, 0x1b, 0x7c, 0xc7, 0xf5, 0x14, 0x6e, 0x75, 0x2b, 0xcc, 0x12, 0xbe, 0x58, 0x2d,
	0xf5, 0x37, 0xb8, 0x33, 0x43, 0x6d, 0xda, 0xf4, 0x30, 0x8b, 0x4e, 0x95, 0xd4, 0x48, 0x5b, 0x9f,
	0x7f, 0xfe, 0xfa, 0x7a, 0x02, 0xe8, 0xaa, 0x3d, 0xdc, 0xee, 0x03, 0x4b, 0xaf, 0xfb, 0xc1, 0x3d,
	0xf8, 0x40, 0x4e, 0x0d, 0xd0, 0xd8, 0xac, 0x9b, 0xf5, 0x41, 0x03, 0x34, 0xaf, 0x47, 0xdb, 0x38,
	0x36, 0x3e, 0xea, 0xf2, 0x41, 0x51, 0x74, 0xdd, 0x0e, 0xbf, 0x02, 0x97, 0xf6, 0x0d, 0x67, 0x8b,
	0x59, 0x26, 0x73, 0xf8, 0x12, 0x90, 0x95, 0x57, 0x42, 0xdb, 0x84, 0xad, 0x50, 0xc3, 0xc6, 0x41,
	0x23, 0x0a, 0x83, 0xcf, 0xb8, 0x51, 0x37, 0x6c, 0x85, 0x28, 0x8d, 0x98, 0x08, 0xcc, 0x0a, 0x1b,
	0xed, 0xdb, 0xb4, 0x87, 0xb0, 0xe9, 0xd3, 0xd2, 0x4c, 0x15, 0x8c, 0x6c, 0xe1, 0xfe, 0xe4, 0x2c,
	0x54, 0x09, 0x17, 0x92, 0x2d, 0xca, 0xdf, 0xdc, 0xf2, 0x7c, 0x14, 0xa1, 0x86, 0xef, 0x01, 0x39,
	0xe3, 0x61, 0x8e, 0x81, 0xd2, 0x5a, 0x66, 0xa0, 0x3b, 0x16, 0x62, 0x0a, 0xac, 0x19, 0x84, 0x1e,
	0x3e, 0x81, 0xc7, 0x0d, 0xb7, 0xb0, 0x85, 0xe4, 0x09, 0xe6, 0x90, 0x93, 0x73, 0xe5, 0xa5, 0xbe,
	0x77, 0xed, 0x82, 0x3b, 0x75, 0x3c, 0xaf, 0xec, 0xef, 0xc7, 0xdd, 0xa3, 0x6c, 0xae, 0x23, 0x1d,
	0x7b, 0xa6, 0xab, 0xf4, 0xa2, 0x07, 0xf4, 0x35, 0xb6, 0x3d, 0xd9, 0x26, 0x2b, 0x03, 0x34, 0x55,
	0xf6, 0xd1, 0x5d, 0x69, 0x2d, 0x8b, 0xa5, 0x1b, 0x36, 0x68, 0x0d, 0xae, 0xfd, 0x13, 0xe4, 0x3a,
	0xf3, 0x2d, 0x20, 0xe7, 0x8b, 0xcf, 0xec, 0x77, 0x14, 0xbd, 0xb9, 0xbd, 0x6c, 0x5c, 0x93, 0xee,
	0x3c, 0xb3, 0xc9, 0x8d, 0xee, 0xc0, 0xb3, 0xd9, 0xfe, 0xfc, 0x08, 0xc8, 0xea, 0xdf, 0x60, 0xc7,
	0xc4, 0xea, 0x1c, 0x66, 0xa2, 0x73, 0x0b, 0xa5, 0xa1, 0xd7, 0x1c, 0x4a, 0x0f, 0x9f, 0x42, 0xff,
	0x3f, 0xb6, 0xb9, 0x4e, 0xbd, 0x78, 0x34, 0xec, 0x45, 0xc2, 0x7c, 0x9a, 0x8d, 0xba, 0x63, 0x95,
	0xb0, 0x78, 0x3e, 0x31, 0xac, 0x7a, 0xab, 0x22, 0x94, 0x2c, 0x1d, 0xdd, 0x8f, 0x14, 0xab, 0x3f,
	0xa3, 0xa3, 0x93, 0xf6, 0xe9, 0xea, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x80, 0x07, 0x57, 0x23,
	0x61, 0x05, 0x00, 0x00,
}

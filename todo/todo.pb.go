// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetTodoParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoParams) Reset()         { *m = GetTodoParams{} }
func (m *GetTodoParams) String() string { return proto.CompactTextString(m) }
func (*GetTodoParams) ProtoMessage()    {}
func (*GetTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *GetTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoParams.Unmarshal(m, b)
}
func (m *GetTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoParams.Marshal(b, m, deterministic)
}
func (m *GetTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoParams.Merge(m, src)
}
func (m *GetTodoParams) XXX_Size() int {
	return xxx_messageInfo_GetTodoParams.Size(m)
}
func (m *GetTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoParams proto.InternalMessageInfo

type AddTodoParams struct {
	Task                 string   `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTodoParams) Reset()         { *m = AddTodoParams{} }
func (m *AddTodoParams) String() string { return proto.CompactTextString(m) }
func (*AddTodoParams) ProtoMessage()    {}
func (*AddTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *AddTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTodoParams.Unmarshal(m, b)
}
func (m *AddTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTodoParams.Marshal(b, m, deterministic)
}
func (m *AddTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTodoParams.Merge(m, src)
}
func (m *AddTodoParams) XXX_Size() int {
	return xxx_messageInfo_AddTodoParams.Size(m)
}
func (m *AddTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_AddTodoParams proto.InternalMessageInfo

func (m *AddTodoParams) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type CompleteTodoParams struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteTodoParams) Reset()         { *m = CompleteTodoParams{} }
func (m *CompleteTodoParams) String() string { return proto.CompactTextString(m) }
func (*CompleteTodoParams) ProtoMessage()    {}
func (*CompleteTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *CompleteTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteTodoParams.Unmarshal(m, b)
}
func (m *CompleteTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteTodoParams.Marshal(b, m, deterministic)
}
func (m *CompleteTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteTodoParams.Merge(m, src)
}
func (m *CompleteTodoParams) XXX_Size() int {
	return xxx_messageInfo_CompleteTodoParams.Size(m)
}
func (m *CompleteTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteTodoParams proto.InternalMessageInfo

func (m *CompleteTodoParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoParams struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoParams) Reset()         { *m = DeleteTodoParams{} }
func (m *DeleteTodoParams) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoParams) ProtoMessage()    {}
func (*DeleteTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *DeleteTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoParams.Unmarshal(m, b)
}
func (m *DeleteTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoParams.Marshal(b, m, deterministic)
}
func (m *DeleteTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoParams.Merge(m, src)
}
func (m *DeleteTodoParams) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoParams.Size(m)
}
func (m *DeleteTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoParams proto.InternalMessageInfo

func (m *DeleteTodoParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TodoObject struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Completed            bool     `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	Task                 string   `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoObject) Reset()         { *m = TodoObject{} }
func (m *TodoObject) String() string { return proto.CompactTextString(m) }
func (*TodoObject) ProtoMessage()    {}
func (*TodoObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *TodoObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoObject.Unmarshal(m, b)
}
func (m *TodoObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoObject.Marshal(b, m, deterministic)
}
func (m *TodoObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoObject.Merge(m, src)
}
func (m *TodoObject) XXX_Size() int {
	return xxx_messageInfo_TodoObject.Size(m)
}
func (m *TodoObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoObject.DiscardUnknown(m)
}

var xxx_messageInfo_TodoObject proto.InternalMessageInfo

func (m *TodoObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TodoObject) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *TodoObject) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type TodoResponse struct {
	Todos                []*TodoObject `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TodoResponse) Reset()         { *m = TodoResponse{} }
func (m *TodoResponse) String() string { return proto.CompactTextString(m) }
func (*TodoResponse) ProtoMessage()    {}
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5}
}

func (m *TodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoResponse.Unmarshal(m, b)
}
func (m *TodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoResponse.Marshal(b, m, deterministic)
}
func (m *TodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoResponse.Merge(m, src)
}
func (m *TodoResponse) XXX_Size() int {
	return xxx_messageInfo_TodoResponse.Size(m)
}
func (m *TodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TodoResponse proto.InternalMessageInfo

func (m *TodoResponse) GetTodos() []*TodoObject {
	if m != nil {
		return m.Todos
	}
	return nil
}

type DeleteResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CompleteResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteResponse) Reset()         { *m = CompleteResponse{} }
func (m *CompleteResponse) String() string { return proto.CompactTextString(m) }
func (*CompleteResponse) ProtoMessage()    {}
func (*CompleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{7}
}

func (m *CompleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteResponse.Unmarshal(m, b)
}
func (m *CompleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteResponse.Marshal(b, m, deterministic)
}
func (m *CompleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteResponse.Merge(m, src)
}
func (m *CompleteResponse) XXX_Size() int {
	return xxx_messageInfo_CompleteResponse.Size(m)
}
func (m *CompleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteResponse proto.InternalMessageInfo

func (m *CompleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTodoParams)(nil), "todo.getTodoParams")
	proto.RegisterType((*AddTodoParams)(nil), "todo.addTodoParams")
	proto.RegisterType((*CompleteTodoParams)(nil), "todo.completeTodoParams")
	proto.RegisterType((*DeleteTodoParams)(nil), "todo.deleteTodoParams")
	proto.RegisterType((*TodoObject)(nil), "todo.todoObject")
	proto.RegisterType((*TodoResponse)(nil), "todo.todoResponse")
	proto.RegisterType((*DeleteResponse)(nil), "todo.deleteResponse")
	proto.RegisterType((*CompleteResponse)(nil), "todo.completeResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x0b, 0xad, 0xb6, 0x9d, 0xfe, 0x91, 0x8c, 0xa6, 0x21, 0xc4, 0x03, 0x59, 0x8d, 0x21,
	0xc6, 0xf4, 0x80, 0xd1, 0x93, 0x27, 0x1f, 0x40, 0x0d, 0xfa, 0x02, 0x94, 0x9d, 0x34, 0xa8, 0xb8,
	0x84, 0xdd, 0xf8, 0xf0, 0x9e, 0x0c, 0x2c, 0x74, 0x97, 0xf6, 0xd0, 0xdb, 0xee, 0xcc, 0x37, 0x33,
	0xbf, 0x7c, 0x33, 0x00, 0x4a, 0x70, 0xb1, 0x2e, 0x2b, 0xa1, 0x04, 0x8e, 0xea, 0x37, 0x3b, 0x83,
	0xc5, 0x96, 0xd4, 0x87, 0xe0, 0xe2, 0x2d, 0xad, 0xd2, 0x42, 0xb2, 0x2b, 0x58, 0xa4, 0x9c, 0x9b,
	0x00, 0x22, 0x8c, 0x54, 0x2a, 0xbf, 0x7c, 0x27, 0x74, 0xa2, 0x69, 0xd2, 0xbc, 0xd9, 0x35, 0x60,
	0x26, 0x8a, 0xf2, 0x9b, 0x14, 0x59, 0xca, 0x25, 0xb8, 0x39, 0x6f, 0x75, 0x6e, 0xce, 0x19, 0x03,
	0x8f, 0xd3, 0x11, 0xcd, 0x8b, 0x66, 0x7a, 0xdd, 0x7c, 0x52, 0xa6, 0xf6, 0xb3, 0x78, 0x09, 0xd3,
	0x6e, 0x0e, 0xf7, 0xdd, 0xd0, 0x89, 0x26, 0x89, 0x09, 0xec, 0xc8, 0x86, 0x16, 0xd9, 0x23, 0xcc,
	0xeb, 0x7e, 0x09, 0xc9, 0x52, 0xfc, 0x48, 0xc2, 0x1b, 0x38, 0xa9, 0xff, 0xd2, 0x77, 0xc2, 0x61,
	0x34, 0x8b, 0xbd, 0x75, 0xe3, 0x80, 0x19, 0x99, 0xe8, 0x34, 0xbb, 0x85, 0xa5, 0x66, 0xdd, 0x55,
	0xfa, 0x30, 0x2e, 0x48, 0xca, 0x74, 0x4b, 0x2d, 0x50, 0xf7, 0x65, 0x77, 0xe0, 0x75, 0x10, 0xc7,
	0xd5, 0xf1, 0x9f, 0x03, 0xb3, 0x7a, 0xc6, 0x3b, 0x55, 0xbf, 0x79, 0x46, 0x18, 0xc3, 0xb8, 0x35,
	0x18, 0xcf, 0x35, 0x4d, 0xcf, 0xef, 0xe0, 0x00, 0x91, 0x0d, 0xf0, 0x09, 0xc0, 0x38, 0x89, 0x2b,
	0xad, 0xd8, 0xf7, 0x36, 0xb8, 0xb0, 0xe3, 0x1d, 0x19, 0x1b, 0xe0, 0x33, 0xcc, 0xed, 0x6d, 0xa1,
	0xaf, 0x75, 0x87, 0x1b, 0x0c, 0x56, 0xfd, 0x8c, 0xd5, 0xe3, 0x01, 0x26, 0xed, 0x9d, 0xc8, 0x0e,
	0xbb, 0x77, 0x37, 0x01, 0x1a, 0x6c, 0x53, 0xb6, 0x39, 0x6d, 0x6e, 0xed, 0xfe, 0x3f, 0x00, 0x00,
	0xff, 0xff, 0x7a, 0xa5, 0x12, 0x23, 0x79, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	AddTodo(ctx context.Context, in *AddTodoParams, opts ...grpc.CallOption) (*TodoObject, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoParams, opts ...grpc.CallOption) (*DeleteResponse, error)
	CompleteTodo(ctx context.Context, in *CompleteTodoParams, opts ...grpc.CallOption) (*CompleteResponse, error)
	GetTodos(ctx context.Context, in *GetTodoParams, opts ...grpc.CallOption) (*TodoResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) AddTodo(ctx context.Context, in *AddTodoParams, opts ...grpc.CallOption) (*TodoObject, error) {
	out := new(TodoObject)
	err := c.cc.Invoke(ctx, "/todo.todoService/addTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoParams, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/deleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CompleteTodo(ctx context.Context, in *CompleteTodoParams, opts ...grpc.CallOption) (*CompleteResponse, error) {
	out := new(CompleteResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/completeTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodos(ctx context.Context, in *GetTodoParams, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/getTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	AddTodo(context.Context, *AddTodoParams) (*TodoObject, error)
	DeleteTodo(context.Context, *DeleteTodoParams) (*DeleteResponse, error)
	CompleteTodo(context.Context, *CompleteTodoParams) (*CompleteResponse, error)
	GetTodos(context.Context, *GetTodoParams) (*TodoResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) AddTodo(ctx context.Context, req *AddTodoParams) (*TodoObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoParams) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (*UnimplementedTodoServiceServer) CompleteTodo(ctx context.Context, req *CompleteTodoParams) (*CompleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTodo not implemented")
}
func (*UnimplementedTodoServiceServer) GetTodos(ctx context.Context, req *GetTodoParams) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).AddTodo(ctx, req.(*AddTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CompleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CompleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/CompleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CompleteTodo(ctx, req.(*CompleteTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/GetTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodos(ctx, req.(*GetTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.todoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addTodo",
			Handler:    _TodoService_AddTodo_Handler,
		},
		{
			MethodName: "deleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
		{
			MethodName: "completeTodo",
			Handler:    _TodoService_CompleteTodo_Handler,
		},
		{
			MethodName: "getTodos",
			Handler:    _TodoService_GetTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

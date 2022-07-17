// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: service_book_repo.proto

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

// BookRepoClient is the client API for BookRepo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookRepoClient interface {
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	GetBookByID(ctx context.Context, in *GetBookByIDRequest, opts ...grpc.CallOption) (*GetBookByIDResponse, error)
}

type bookRepoClient struct {
	cc grpc.ClientConnInterface
}

func NewBookRepoClient(cc grpc.ClientConnInterface) BookRepoClient {
	return &bookRepoClient{cc}
}

func (c *bookRepoClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, "/pb.BookRepo/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookRepoClient) GetBookByID(ctx context.Context, in *GetBookByIDRequest, opts ...grpc.CallOption) (*GetBookByIDResponse, error) {
	out := new(GetBookByIDResponse)
	err := c.cc.Invoke(ctx, "/pb.BookRepo/GetBookByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookRepoServer is the server API for BookRepo service.
// All implementations must embed UnimplementedBookRepoServer
// for forward compatibility
type BookRepoServer interface {
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	GetBookByID(context.Context, *GetBookByIDRequest) (*GetBookByIDResponse, error)
	mustEmbedUnimplementedBookRepoServer()
}

// UnimplementedBookRepoServer must be embedded to have forward compatible implementations.
type UnimplementedBookRepoServer struct {
}

func (UnimplementedBookRepoServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookRepoServer) GetBookByID(context.Context, *GetBookByIDRequest) (*GetBookByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookByID not implemented")
}
func (UnimplementedBookRepoServer) mustEmbedUnimplementedBookRepoServer() {}

// UnsafeBookRepoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookRepoServer will
// result in compilation errors.
type UnsafeBookRepoServer interface {
	mustEmbedUnimplementedBookRepoServer()
}

func RegisterBookRepoServer(s grpc.ServiceRegistrar, srv BookRepoServer) {
	s.RegisterService(&BookRepo_ServiceDesc, srv)
}

func _BookRepo_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRepoServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookRepo/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRepoServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookRepo_GetBookByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRepoServer).GetBookByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookRepo/GetBookByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRepoServer).GetBookByID(ctx, req.(*GetBookByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookRepo_ServiceDesc is the grpc.ServiceDesc for BookRepo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookRepo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BookRepo",
	HandlerType: (*BookRepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookRepo_CreateBook_Handler,
		},
		{
			MethodName: "GetBookByID",
			Handler:    _BookRepo_GetBookByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_book_repo.proto",
}

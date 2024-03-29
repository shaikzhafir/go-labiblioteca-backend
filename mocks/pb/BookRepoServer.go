// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"
	pb "go-labiblioteca-backend/pb"

	mock "github.com/stretchr/testify/mock"
)

// BookRepoServer is an autogenerated mock type for the BookRepoServer type
type BookRepoServer struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: _a0, _a1
func (_m *BookRepoServer) CreateBook(_a0 context.Context, _a1 *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.CreateBookResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateBookRequest) *pb.CreateBookResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CreateBookResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateBookRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookByID provides a mock function with given fields: _a0, _a1
func (_m *BookRepoServer) GetBookByID(_a0 context.Context, _a1 *pb.GetBookByIDRequest) (*pb.GetBookByIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetBookByIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetBookByIDRequest) *pb.GetBookByIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetBookByIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetBookByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedBookRepoServer provides a mock function with given fields:
func (_m *BookRepoServer) mustEmbedUnimplementedBookRepoServer() {
	_m.Called()
}

type mockConstructorTestingTNewBookRepoServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookRepoServer creates a new instance of BookRepoServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookRepoServer(t mockConstructorTestingTNewBookRepoServer) *BookRepoServer {
	mock := &BookRepoServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

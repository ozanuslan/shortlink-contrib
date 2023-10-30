// Code generated by mockery v2.36.0. DO NOT EDIT.

package v1

import (
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	mock "github.com/stretchr/testify/mock"
)

// MockWatchServiceServer is an autogenerated mock type for the WatchServiceServer type
type MockWatchServiceServer struct {
	mock.Mock
}

type MockWatchServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWatchServiceServer) EXPECT() *MockWatchServiceServer_Expecter {
	return &MockWatchServiceServer_Expecter{mock: &_m.Mock}
}

// Watch provides a mock function with given fields: _a0, _a1
func (_m *MockWatchServiceServer) Watch(_a0 *v1.WatchRequest, _a1 v1.WatchService_WatchServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.WatchRequest, v1.WatchService_WatchServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWatchServiceServer_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockWatchServiceServer_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - _a0 *v1.WatchRequest
//   - _a1 v1.WatchService_WatchServer
func (_e *MockWatchServiceServer_Expecter) Watch(_a0 interface{}, _a1 interface{}) *MockWatchServiceServer_Watch_Call {
	return &MockWatchServiceServer_Watch_Call{Call: _e.mock.On("Watch", _a0, _a1)}
}

func (_c *MockWatchServiceServer_Watch_Call) Run(run func(_a0 *v1.WatchRequest, _a1 v1.WatchService_WatchServer)) *MockWatchServiceServer_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.WatchRequest), args[1].(v1.WatchService_WatchServer))
	})
	return _c
}

func (_c *MockWatchServiceServer_Watch_Call) Return(_a0 error) *MockWatchServiceServer_Watch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWatchServiceServer_Watch_Call) RunAndReturn(run func(*v1.WatchRequest, v1.WatchService_WatchServer) error) *MockWatchServiceServer_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedWatchServiceServer provides a mock function with given fields:
func (_m *MockWatchServiceServer) mustEmbedUnimplementedWatchServiceServer() {
	_m.Called()
}

// MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedWatchServiceServer'
type MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedWatchServiceServer is a helper method to define mock.On call
func (_e *MockWatchServiceServer_Expecter) mustEmbedUnimplementedWatchServiceServer() *MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	return &MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedWatchServiceServer")}
}

func (_c *MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call) Run(run func()) *MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call) Return() *MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call) RunAndReturn(run func()) *MockWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWatchServiceServer creates a new instance of MockWatchServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWatchServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWatchServiceServer {
	mock := &MockWatchServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
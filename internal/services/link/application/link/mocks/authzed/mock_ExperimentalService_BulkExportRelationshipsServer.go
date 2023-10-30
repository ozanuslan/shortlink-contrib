// Code generated by mockery v2.36.0. DO NOT EDIT.

package v1

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// MockExperimentalService_BulkExportRelationshipsServer is an autogenerated mock type for the ExperimentalService_BulkExportRelationshipsServer type
type MockExperimentalService_BulkExportRelationshipsServer struct {
	mock.Mock
}

type MockExperimentalService_BulkExportRelationshipsServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExperimentalService_BulkExportRelationshipsServer) EXPECT() *MockExperimentalService_BulkExportRelationshipsServer_Expecter {
	return &MockExperimentalService_BulkExportRelationshipsServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockExperimentalService_BulkExportRelationshipsServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockExperimentalService_BulkExportRelationshipsServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockExperimentalService_BulkExportRelationshipsServer_Expecter) Context() *MockExperimentalService_BulkExportRelationshipsServer_Context_Call {
	return &MockExperimentalService_BulkExportRelationshipsServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_Context_Call) Run(run func()) *MockExperimentalService_BulkExportRelationshipsServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_Context_Call) Return(_a0 context.Context) *MockExperimentalService_BulkExportRelationshipsServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_Context_Call) RunAndReturn(run func() context.Context) *MockExperimentalService_BulkExportRelationshipsServer_Context_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockExperimentalService_BulkExportRelationshipsServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockExperimentalService_BulkExportRelationshipsServer_Expecter) RecvMsg(m interface{}) *MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call {
	return &MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call) Run(run func(m interface{})) *MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockExperimentalService_BulkExportRelationshipsServer_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockExperimentalService_BulkExportRelationshipsServer) Send(_a0 *v1.BulkExportRelationshipsResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.BulkExportRelationshipsResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockExperimentalService_BulkExportRelationshipsServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *v1.BulkExportRelationshipsResponse
func (_e *MockExperimentalService_BulkExportRelationshipsServer_Expecter) Send(_a0 interface{}) *MockExperimentalService_BulkExportRelationshipsServer_Send_Call {
	return &MockExperimentalService_BulkExportRelationshipsServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_Send_Call) Run(run func(_a0 *v1.BulkExportRelationshipsResponse)) *MockExperimentalService_BulkExportRelationshipsServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.BulkExportRelationshipsResponse))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_Send_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_Send_Call) RunAndReturn(run func(*v1.BulkExportRelationshipsResponse) error) *MockExperimentalService_BulkExportRelationshipsServer_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockExperimentalService_BulkExportRelationshipsServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockExperimentalService_BulkExportRelationshipsServer_Expecter) SendHeader(_a0 interface{}) *MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call {
	return &MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockExperimentalService_BulkExportRelationshipsServer_SendHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockExperimentalService_BulkExportRelationshipsServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockExperimentalService_BulkExportRelationshipsServer_Expecter) SendMsg(m interface{}) *MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call {
	return &MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call) Run(run func(m interface{})) *MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockExperimentalService_BulkExportRelationshipsServer_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockExperimentalService_BulkExportRelationshipsServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockExperimentalService_BulkExportRelationshipsServer_Expecter) SetHeader(_a0 interface{}) *MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call {
	return &MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockExperimentalService_BulkExportRelationshipsServer_SetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockExperimentalService_BulkExportRelationshipsServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockExperimentalService_BulkExportRelationshipsServer_Expecter) SetTrailer(_a0 interface{}) *MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call {
	return &MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call) Return() *MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call) RunAndReturn(run func(metadata.MD)) *MockExperimentalService_BulkExportRelationshipsServer_SetTrailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExperimentalService_BulkExportRelationshipsServer creates a new instance of MockExperimentalService_BulkExportRelationshipsServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExperimentalService_BulkExportRelationshipsServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExperimentalService_BulkExportRelationshipsServer {
	mock := &MockExperimentalService_BulkExportRelationshipsServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by mockery v2.36.0. DO NOT EDIT.

package v1

import mock "github.com/stretchr/testify/mock"

// MockisBulkCheckPermissionPair_Response is an autogenerated mock type for the isBulkCheckPermissionPair_Response type
type MockisBulkCheckPermissionPair_Response struct {
	mock.Mock
}

type MockisBulkCheckPermissionPair_Response_Expecter struct {
	mock *mock.Mock
}

func (_m *MockisBulkCheckPermissionPair_Response) EXPECT() *MockisBulkCheckPermissionPair_Response_Expecter {
	return &MockisBulkCheckPermissionPair_Response_Expecter{mock: &_m.Mock}
}

// isBulkCheckPermissionPair_Response provides a mock function with given fields:
func (_m *MockisBulkCheckPermissionPair_Response) isBulkCheckPermissionPair_Response() {
	_m.Called()
}

// MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isBulkCheckPermissionPair_Response'
type MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call struct {
	*mock.Call
}

// isBulkCheckPermissionPair_Response is a helper method to define mock.On call
func (_e *MockisBulkCheckPermissionPair_Response_Expecter) isBulkCheckPermissionPair_Response() *MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call {
	return &MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call{Call: _e.mock.On("isBulkCheckPermissionPair_Response")}
}

func (_c *MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call) Run(run func()) *MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call) Return() *MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call) RunAndReturn(run func()) *MockisBulkCheckPermissionPair_Response_isBulkCheckPermissionPair_Response_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockisBulkCheckPermissionPair_Response creates a new instance of MockisBulkCheckPermissionPair_Response. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockisBulkCheckPermissionPair_Response(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockisBulkCheckPermissionPair_Response {
	mock := &MockisBulkCheckPermissionPair_Response{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
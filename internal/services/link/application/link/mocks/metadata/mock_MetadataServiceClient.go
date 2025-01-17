// Code generated by mockery v2.36.0. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
)

// MockMetadataServiceClient is an autogenerated mock type for the MetadataServiceClient type
type MockMetadataServiceClient struct {
	mock.Mock
}

type MockMetadataServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetadataServiceClient) EXPECT() *MockMetadataServiceClient_Expecter {
	return &MockMetadataServiceClient_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *MockMetadataServiceClient) Get(ctx context.Context, in *v1.MetadataServiceGetRequest, opts ...grpc.CallOption) (*v1.MetadataServiceGetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.MetadataServiceGetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceGetRequest, ...grpc.CallOption) (*v1.MetadataServiceGetResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceGetRequest, ...grpc.CallOption) *v1.MetadataServiceGetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MetadataServiceGetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MetadataServiceGetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetadataServiceClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockMetadataServiceClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.MetadataServiceGetRequest
//   - opts ...grpc.CallOption
func (_e *MockMetadataServiceClient_Expecter) Get(ctx interface{}, in interface{}, opts ...interface{}) *MockMetadataServiceClient_Get_Call {
	return &MockMetadataServiceClient_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockMetadataServiceClient_Get_Call) Run(run func(ctx context.Context, in *v1.MetadataServiceGetRequest, opts ...grpc.CallOption)) *MockMetadataServiceClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.MetadataServiceGetRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockMetadataServiceClient_Get_Call) Return(_a0 *v1.MetadataServiceGetResponse, _a1 error) *MockMetadataServiceClient_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetadataServiceClient_Get_Call) RunAndReturn(run func(context.Context, *v1.MetadataServiceGetRequest, ...grpc.CallOption) (*v1.MetadataServiceGetResponse, error)) *MockMetadataServiceClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, in, opts
func (_m *MockMetadataServiceClient) Set(ctx context.Context, in *v1.MetadataServiceSetRequest, opts ...grpc.CallOption) (*v1.MetadataServiceSetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.MetadataServiceSetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceSetRequest, ...grpc.CallOption) (*v1.MetadataServiceSetResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceSetRequest, ...grpc.CallOption) *v1.MetadataServiceSetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MetadataServiceSetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MetadataServiceSetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetadataServiceClient_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockMetadataServiceClient_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.MetadataServiceSetRequest
//   - opts ...grpc.CallOption
func (_e *MockMetadataServiceClient_Expecter) Set(ctx interface{}, in interface{}, opts ...interface{}) *MockMetadataServiceClient_Set_Call {
	return &MockMetadataServiceClient_Set_Call{Call: _e.mock.On("Set",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockMetadataServiceClient_Set_Call) Run(run func(ctx context.Context, in *v1.MetadataServiceSetRequest, opts ...grpc.CallOption)) *MockMetadataServiceClient_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.MetadataServiceSetRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockMetadataServiceClient_Set_Call) Return(_a0 *v1.MetadataServiceSetResponse, _a1 error) *MockMetadataServiceClient_Set_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetadataServiceClient_Set_Call) RunAndReturn(run func(context.Context, *v1.MetadataServiceSetRequest, ...grpc.CallOption) (*v1.MetadataServiceSetResponse, error)) *MockMetadataServiceClient_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetadataServiceClient creates a new instance of MockMetadataServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetadataServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetadataServiceClient {
	mock := &MockMetadataServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

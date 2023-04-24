// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UnitOfWork is an autogenerated mock type for the UnitOfWork type
type UnitOfWork[T interface{}] struct {
	mock.Mock
}

// Commit provides a mock function with given fields: ctx
func (_m *UnitOfWork[T]) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterClean provides a mock function with given fields: in
func (_m *UnitOfWork[T]) RegisterClean(in ...T) error {
	_va := make([]interface{}, len(in))
	for _i := range in {
		_va[_i] = in[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...T) error); ok {
		r0 = rf(in...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterDeleted provides a mock function with given fields: in
func (_m *UnitOfWork[T]) RegisterDeleted(in ...T) error {
	_va := make([]interface{}, len(in))
	for _i := range in {
		_va[_i] = in[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...T) error); ok {
		r0 = rf(in...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterDirty provides a mock function with given fields: in
func (_m *UnitOfWork[T]) RegisterDirty(in ...T) error {
	_va := make([]interface{}, len(in))
	for _i := range in {
		_va[_i] = in[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...T) error); ok {
		r0 = rf(in...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterNew provides a mock function with given fields: in
func (_m *UnitOfWork[T]) RegisterNew(in ...T) error {
	_va := make([]interface{}, len(in))
	for _i := range in {
		_va[_i] = in[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...T) error); ok {
		r0 = rf(in...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields: ctx
func (_m *UnitOfWork[T]) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUnitOfWork interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnitOfWork creates a new instance of UnitOfWork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnitOfWork[T interface{}](t mockConstructorTestingTNewUnitOfWork) *UnitOfWork[T] {
	mock := &UnitOfWork[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

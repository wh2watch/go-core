// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"
	cqrs "core/cqrs"

	mock "github.com/stretchr/testify/mock"
)

// QueryBus is an autogenerated mock type for the QueryBus type
type QueryBus struct {
	mock.Mock
}

type QueryBus_Expecter struct {
	mock *mock.Mock
}

func (_m *QueryBus) EXPECT() *QueryBus_Expecter {
	return &QueryBus_Expecter{mock: &_m.Mock}
}

// Dispatch provides a mock function with given fields: ctx, qry, res
func (_m *QueryBus) Dispatch(ctx context.Context, qry cqrs.Query, res cqrs.Response) error {
	ret := _m.Called(ctx, qry, res)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, cqrs.Query, cqrs.Response) error); ok {
		r0 = rf(ctx, qry, res)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryBus_Dispatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dispatch'
type QueryBus_Dispatch_Call struct {
	*mock.Call
}

// Dispatch is a helper method to define mock.On call
//   - ctx context.Context
//   - qry cqrs.Query
//   - res cqrs.Response
func (_e *QueryBus_Expecter) Dispatch(ctx interface{}, qry interface{}, res interface{}) *QueryBus_Dispatch_Call {
	return &QueryBus_Dispatch_Call{Call: _e.mock.On("Dispatch", ctx, qry, res)}
}

func (_c *QueryBus_Dispatch_Call) Run(run func(ctx context.Context, qry cqrs.Query, res cqrs.Response)) *QueryBus_Dispatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cqrs.Query), args[2].(cqrs.Response))
	})
	return _c
}

func (_c *QueryBus_Dispatch_Call) Return(_a0 error) *QueryBus_Dispatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryBus_Dispatch_Call) RunAndReturn(run func(context.Context, cqrs.Query, cqrs.Response) error) *QueryBus_Dispatch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewQueryBus interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueryBus creates a new instance of QueryBus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueryBus(t mockConstructorTestingTNewQueryBus) *QueryBus {
	mock := &QueryBus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

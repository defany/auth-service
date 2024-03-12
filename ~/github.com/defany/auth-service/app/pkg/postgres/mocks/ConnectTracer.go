// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockpgx

import (
	context "context"

	pgx "github.com/jackc/pgx/v5"
	mock "github.com/stretchr/testify/mock"
)

// MockConnectTracer is an autogenerated mock type for the ConnectTracer type
type MockConnectTracer struct {
	mock.Mock
}

type MockConnectTracer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnectTracer) EXPECT() *MockConnectTracer_Expecter {
	return &MockConnectTracer_Expecter{mock: &_m.Mock}
}

// TraceConnectEnd provides a mock function with given fields: ctx, data
func (_m *MockConnectTracer) TraceConnectEnd(ctx context.Context, data pgx.TraceConnectEndData) {
	_m.Called(ctx, data)
}

// MockConnectTracer_TraceConnectEnd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TraceConnectEnd'
type MockConnectTracer_TraceConnectEnd_Call struct {
	*mock.Call
}

// TraceConnectEnd is a helper method to define mock.On call
//   - ctx context.Context
//   - data pgx.TraceConnectEndData
func (_e *MockConnectTracer_Expecter) TraceConnectEnd(ctx interface{}, data interface{}) *MockConnectTracer_TraceConnectEnd_Call {
	return &MockConnectTracer_TraceConnectEnd_Call{Call: _e.mock.On("TraceConnectEnd", ctx, data)}
}

func (_c *MockConnectTracer_TraceConnectEnd_Call) Run(run func(ctx context.Context, data pgx.TraceConnectEndData)) *MockConnectTracer_TraceConnectEnd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.TraceConnectEndData))
	})
	return _c
}

func (_c *MockConnectTracer_TraceConnectEnd_Call) Return() *MockConnectTracer_TraceConnectEnd_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockConnectTracer_TraceConnectEnd_Call) RunAndReturn(run func(context.Context, pgx.TraceConnectEndData)) *MockConnectTracer_TraceConnectEnd_Call {
	_c.Call.Return(run)
	return _c
}

// TraceConnectStart provides a mock function with given fields: ctx, data
func (_m *MockConnectTracer) TraceConnectStart(ctx context.Context, data pgx.TraceConnectStartData) context.Context {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for TraceConnectStart")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TraceConnectStartData) context.Context); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockConnectTracer_TraceConnectStart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TraceConnectStart'
type MockConnectTracer_TraceConnectStart_Call struct {
	*mock.Call
}

// TraceConnectStart is a helper method to define mock.On call
//   - ctx context.Context
//   - data pgx.TraceConnectStartData
func (_e *MockConnectTracer_Expecter) TraceConnectStart(ctx interface{}, data interface{}) *MockConnectTracer_TraceConnectStart_Call {
	return &MockConnectTracer_TraceConnectStart_Call{Call: _e.mock.On("TraceConnectStart", ctx, data)}
}

func (_c *MockConnectTracer_TraceConnectStart_Call) Run(run func(ctx context.Context, data pgx.TraceConnectStartData)) *MockConnectTracer_TraceConnectStart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.TraceConnectStartData))
	})
	return _c
}

func (_c *MockConnectTracer_TraceConnectStart_Call) Return(_a0 context.Context) *MockConnectTracer_TraceConnectStart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectTracer_TraceConnectStart_Call) RunAndReturn(run func(context.Context, pgx.TraceConnectStartData) context.Context) *MockConnectTracer_TraceConnectStart_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectTracer creates a new instance of MockConnectTracer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectTracer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectTracer {
	mock := &MockConnectTracer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

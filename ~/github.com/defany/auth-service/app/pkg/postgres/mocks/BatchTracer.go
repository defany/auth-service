// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockpgx

import (
	context "context"

	pgx "github.com/jackc/pgx/v5"
	mock "github.com/stretchr/testify/mock"
)

// MockBatchTracer is an autogenerated mock type for the BatchTracer type
type MockBatchTracer struct {
	mock.Mock
}

type MockBatchTracer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBatchTracer) EXPECT() *MockBatchTracer_Expecter {
	return &MockBatchTracer_Expecter{mock: &_m.Mock}
}

// TraceBatchEnd provides a mock function with given fields: ctx, conn, data
func (_m *MockBatchTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
	_m.Called(ctx, conn, data)
}

// MockBatchTracer_TraceBatchEnd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TraceBatchEnd'
type MockBatchTracer_TraceBatchEnd_Call struct {
	*mock.Call
}

// TraceBatchEnd is a helper method to define mock.On call
//   - ctx context.Context
//   - conn *pgx.Conn
//   - data pgx.TraceBatchEndData
func (_e *MockBatchTracer_Expecter) TraceBatchEnd(ctx interface{}, conn interface{}, data interface{}) *MockBatchTracer_TraceBatchEnd_Call {
	return &MockBatchTracer_TraceBatchEnd_Call{Call: _e.mock.On("TraceBatchEnd", ctx, conn, data)}
}

func (_c *MockBatchTracer_TraceBatchEnd_Call) Run(run func(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData)) *MockBatchTracer_TraceBatchEnd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pgx.Conn), args[2].(pgx.TraceBatchEndData))
	})
	return _c
}

func (_c *MockBatchTracer_TraceBatchEnd_Call) Return() *MockBatchTracer_TraceBatchEnd_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBatchTracer_TraceBatchEnd_Call) RunAndReturn(run func(context.Context, *pgx.Conn, pgx.TraceBatchEndData)) *MockBatchTracer_TraceBatchEnd_Call {
	_c.Call.Return(run)
	return _c
}

// TraceBatchQuery provides a mock function with given fields: ctx, conn, data
func (_m *MockBatchTracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
	_m.Called(ctx, conn, data)
}

// MockBatchTracer_TraceBatchQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TraceBatchQuery'
type MockBatchTracer_TraceBatchQuery_Call struct {
	*mock.Call
}

// TraceBatchQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - conn *pgx.Conn
//   - data pgx.TraceBatchQueryData
func (_e *MockBatchTracer_Expecter) TraceBatchQuery(ctx interface{}, conn interface{}, data interface{}) *MockBatchTracer_TraceBatchQuery_Call {
	return &MockBatchTracer_TraceBatchQuery_Call{Call: _e.mock.On("TraceBatchQuery", ctx, conn, data)}
}

func (_c *MockBatchTracer_TraceBatchQuery_Call) Run(run func(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData)) *MockBatchTracer_TraceBatchQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pgx.Conn), args[2].(pgx.TraceBatchQueryData))
	})
	return _c
}

func (_c *MockBatchTracer_TraceBatchQuery_Call) Return() *MockBatchTracer_TraceBatchQuery_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBatchTracer_TraceBatchQuery_Call) RunAndReturn(run func(context.Context, *pgx.Conn, pgx.TraceBatchQueryData)) *MockBatchTracer_TraceBatchQuery_Call {
	_c.Call.Return(run)
	return _c
}

// TraceBatchStart provides a mock function with given fields: ctx, conn, data
func (_m *MockBatchTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	ret := _m.Called(ctx, conn, data)

	if len(ret) == 0 {
		panic("no return value specified for TraceBatchStart")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Conn, pgx.TraceBatchStartData) context.Context); ok {
		r0 = rf(ctx, conn, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockBatchTracer_TraceBatchStart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TraceBatchStart'
type MockBatchTracer_TraceBatchStart_Call struct {
	*mock.Call
}

// TraceBatchStart is a helper method to define mock.On call
//   - ctx context.Context
//   - conn *pgx.Conn
//   - data pgx.TraceBatchStartData
func (_e *MockBatchTracer_Expecter) TraceBatchStart(ctx interface{}, conn interface{}, data interface{}) *MockBatchTracer_TraceBatchStart_Call {
	return &MockBatchTracer_TraceBatchStart_Call{Call: _e.mock.On("TraceBatchStart", ctx, conn, data)}
}

func (_c *MockBatchTracer_TraceBatchStart_Call) Run(run func(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData)) *MockBatchTracer_TraceBatchStart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pgx.Conn), args[2].(pgx.TraceBatchStartData))
	})
	return _c
}

func (_c *MockBatchTracer_TraceBatchStart_Call) Return(_a0 context.Context) *MockBatchTracer_TraceBatchStart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBatchTracer_TraceBatchStart_Call) RunAndReturn(run func(context.Context, *pgx.Conn, pgx.TraceBatchStartData) context.Context) *MockBatchTracer_TraceBatchStart_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBatchTracer creates a new instance of MockBatchTracer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBatchTracer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBatchTracer {
	mock := &MockBatchTracer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockpostgres

import (
	context "context"

	pgconn "github.com/jackc/pgx/v5/pgconn"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	postgres "github.com/defany/auth-service/app/pkg/postgres"
)

// MockPostgres is an autogenerated mock type for the Postgres type
type MockPostgres struct {
	mock.Mock
}

type MockPostgres_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPostgres) EXPECT() *MockPostgres_Expecter {
	return &MockPostgres_Expecter{mock: &_m.Mock}
}

// BeginTx provides a mock function with given fields: ctx, txOptions
func (_m *MockPostgres) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (postgres.Tx, error) {
	ret := _m.Called(ctx, txOptions)

	if len(ret) == 0 {
		panic("no return value specified for BeginTx")
	}

	var r0 postgres.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) (postgres.Tx, error)); ok {
		return rf(ctx, txOptions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) postgres.Tx); ok {
		r0 = rf(ctx, txOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(postgres.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.TxOptions) error); ok {
		r1 = rf(ctx, txOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPostgres_BeginTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginTx'
type MockPostgres_BeginTx_Call struct {
	*mock.Call
}

// BeginTx is a helper method to define mock.On call
//   - ctx context.Context
//   - txOptions pgx.TxOptions
func (_e *MockPostgres_Expecter) BeginTx(ctx interface{}, txOptions interface{}) *MockPostgres_BeginTx_Call {
	return &MockPostgres_BeginTx_Call{Call: _e.mock.On("BeginTx", ctx, txOptions)}
}

func (_c *MockPostgres_BeginTx_Call) Run(run func(ctx context.Context, txOptions pgx.TxOptions)) *MockPostgres_BeginTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.TxOptions))
	})
	return _c
}

func (_c *MockPostgres_BeginTx_Call) Return(_a0 postgres.Tx, _a1 error) *MockPostgres_BeginTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPostgres_BeginTx_Call) RunAndReturn(run func(context.Context, pgx.TxOptions) (postgres.Tx, error)) *MockPostgres_BeginTx_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockPostgres) Close() {
	_m.Called()
}

// MockPostgres_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockPostgres_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockPostgres_Expecter) Close() *MockPostgres_Close_Call {
	return &MockPostgres_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockPostgres_Close_Call) Run(run func()) *MockPostgres_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPostgres_Close_Call) Return() *MockPostgres_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPostgres_Close_Call) RunAndReturn(run func()) *MockPostgres_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: ctx, query, args
func (_m *MockPostgres) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgconn.CommandTag); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Get(0).(pgconn.CommandTag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPostgres_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockPostgres_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockPostgres_Expecter) Exec(ctx interface{}, query interface{}, args ...interface{}) *MockPostgres_Exec_Call {
	return &MockPostgres_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockPostgres_Exec_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockPostgres_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockPostgres_Exec_Call) Return(commandTag pgconn.CommandTag, err error) *MockPostgres_Exec_Call {
	_c.Call.Return(commandTag, err)
	return _c
}

func (_c *MockPostgres_Exec_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)) *MockPostgres_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Pool provides a mock function with given fields:
func (_m *MockPostgres) Pool() *pgxpool.Pool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Pool")
	}

	var r0 *pgxpool.Pool
	if rf, ok := ret.Get(0).(func() *pgxpool.Pool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pgxpool.Pool)
		}
	}

	return r0
}

// MockPostgres_Pool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pool'
type MockPostgres_Pool_Call struct {
	*mock.Call
}

// Pool is a helper method to define mock.On call
func (_e *MockPostgres_Expecter) Pool() *MockPostgres_Pool_Call {
	return &MockPostgres_Pool_Call{Call: _e.mock.On("Pool")}
}

func (_c *MockPostgres_Pool_Call) Run(run func()) *MockPostgres_Pool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPostgres_Pool_Call) Return(_a0 *pgxpool.Pool) *MockPostgres_Pool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostgres_Pool_Call) RunAndReturn(run func() *pgxpool.Pool) *MockPostgres_Pool_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, query, args
func (_m *MockPostgres) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgx.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPostgres_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockPostgres_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockPostgres_Expecter) Query(ctx interface{}, query interface{}, args ...interface{}) *MockPostgres_Query_Call {
	return &MockPostgres_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockPostgres_Query_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockPostgres_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockPostgres_Query_Call) Return(_a0 pgx.Rows, _a1 error) *MockPostgres_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPostgres_Query_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgx.Rows, error)) *MockPostgres_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRow provides a mock function with given fields: ctx, query, args
func (_m *MockPostgres) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRow")
	}

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// MockPostgres_QueryRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRow'
type MockPostgres_QueryRow_Call struct {
	*mock.Call
}

// QueryRow is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockPostgres_Expecter) QueryRow(ctx interface{}, query interface{}, args ...interface{}) *MockPostgres_QueryRow_Call {
	return &MockPostgres_QueryRow_Call{Call: _e.mock.On("QueryRow",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockPostgres_QueryRow_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockPostgres_QueryRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockPostgres_QueryRow_Call) Return(_a0 pgx.Row) *MockPostgres_QueryRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostgres_QueryRow_Call) RunAndReturn(run func(context.Context, string, ...interface{}) pgx.Row) *MockPostgres_QueryRow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPostgres creates a new instance of MockPostgres. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPostgres(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPostgres {
	mock := &MockPostgres{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

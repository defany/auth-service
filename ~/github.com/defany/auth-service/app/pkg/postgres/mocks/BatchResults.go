// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockpgx

import (
	pgconn "github.com/jackc/pgx/v5/pgconn"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"
)

// MockBatchResults is an autogenerated mock type for the BatchResults type
type MockBatchResults struct {
	mock.Mock
}

type MockBatchResults_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBatchResults) EXPECT() *MockBatchResults_Expecter {
	return &MockBatchResults_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockBatchResults) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBatchResults_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockBatchResults_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockBatchResults_Expecter) Close() *MockBatchResults_Close_Call {
	return &MockBatchResults_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockBatchResults_Close_Call) Run(run func()) *MockBatchResults_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBatchResults_Close_Call) Return(_a0 error) *MockBatchResults_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBatchResults_Close_Call) RunAndReturn(run func() error) *MockBatchResults_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields:
func (_m *MockBatchResults) Exec() (pgconn.CommandTag, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func() (pgconn.CommandTag, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() pgconn.CommandTag); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pgconn.CommandTag)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBatchResults_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockBatchResults_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
func (_e *MockBatchResults_Expecter) Exec() *MockBatchResults_Exec_Call {
	return &MockBatchResults_Exec_Call{Call: _e.mock.On("Exec")}
}

func (_c *MockBatchResults_Exec_Call) Run(run func()) *MockBatchResults_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBatchResults_Exec_Call) Return(_a0 pgconn.CommandTag, _a1 error) *MockBatchResults_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBatchResults_Exec_Call) RunAndReturn(run func() (pgconn.CommandTag, error)) *MockBatchResults_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields:
func (_m *MockBatchResults) Query() (pgx.Rows, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func() (pgx.Rows, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() pgx.Rows); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBatchResults_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockBatchResults_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
func (_e *MockBatchResults_Expecter) Query() *MockBatchResults_Query_Call {
	return &MockBatchResults_Query_Call{Call: _e.mock.On("Query")}
}

func (_c *MockBatchResults_Query_Call) Run(run func()) *MockBatchResults_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBatchResults_Query_Call) Return(_a0 pgx.Rows, _a1 error) *MockBatchResults_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBatchResults_Query_Call) RunAndReturn(run func() (pgx.Rows, error)) *MockBatchResults_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRow provides a mock function with given fields:
func (_m *MockBatchResults) QueryRow() pgx.Row {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for QueryRow")
	}

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func() pgx.Row); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// MockBatchResults_QueryRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRow'
type MockBatchResults_QueryRow_Call struct {
	*mock.Call
}

// QueryRow is a helper method to define mock.On call
func (_e *MockBatchResults_Expecter) QueryRow() *MockBatchResults_QueryRow_Call {
	return &MockBatchResults_QueryRow_Call{Call: _e.mock.On("QueryRow")}
}

func (_c *MockBatchResults_QueryRow_Call) Run(run func()) *MockBatchResults_QueryRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBatchResults_QueryRow_Call) Return(_a0 pgx.Row) *MockBatchResults_QueryRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBatchResults_QueryRow_Call) RunAndReturn(run func() pgx.Row) *MockBatchResults_QueryRow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBatchResults creates a new instance of MockBatchResults. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBatchResults(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBatchResults {
	mock := &MockBatchResults{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
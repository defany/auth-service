// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockauthservice

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockJWTConfig is an autogenerated mock type for the JWTConfig type
type MockJWTConfig struct {
	mock.Mock
}

type MockJWTConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockJWTConfig) EXPECT() *MockJWTConfig_Expecter {
	return &MockJWTConfig_Expecter{mock: &_m.Mock}
}

// AccessSecretKey provides a mock function with given fields:
func (_m *MockJWTConfig) AccessSecretKey() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AccessSecretKey")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockJWTConfig_AccessSecretKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AccessSecretKey'
type MockJWTConfig_AccessSecretKey_Call struct {
	*mock.Call
}

// AccessSecretKey is a helper method to define mock.On call
func (_e *MockJWTConfig_Expecter) AccessSecretKey() *MockJWTConfig_AccessSecretKey_Call {
	return &MockJWTConfig_AccessSecretKey_Call{Call: _e.mock.On("AccessSecretKey")}
}

func (_c *MockJWTConfig_AccessSecretKey_Call) Run(run func()) *MockJWTConfig_AccessSecretKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJWTConfig_AccessSecretKey_Call) Return(_a0 []byte) *MockJWTConfig_AccessSecretKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJWTConfig_AccessSecretKey_Call) RunAndReturn(run func() []byte) *MockJWTConfig_AccessSecretKey_Call {
	_c.Call.Return(run)
	return _c
}

// AccessTokenDuration provides a mock function with given fields:
func (_m *MockJWTConfig) AccessTokenDuration() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AccessTokenDuration")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// MockJWTConfig_AccessTokenDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AccessTokenDuration'
type MockJWTConfig_AccessTokenDuration_Call struct {
	*mock.Call
}

// AccessTokenDuration is a helper method to define mock.On call
func (_e *MockJWTConfig_Expecter) AccessTokenDuration() *MockJWTConfig_AccessTokenDuration_Call {
	return &MockJWTConfig_AccessTokenDuration_Call{Call: _e.mock.On("AccessTokenDuration")}
}

func (_c *MockJWTConfig_AccessTokenDuration_Call) Run(run func()) *MockJWTConfig_AccessTokenDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJWTConfig_AccessTokenDuration_Call) Return(_a0 time.Duration) *MockJWTConfig_AccessTokenDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJWTConfig_AccessTokenDuration_Call) RunAndReturn(run func() time.Duration) *MockJWTConfig_AccessTokenDuration_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshSecretKey provides a mock function with given fields:
func (_m *MockJWTConfig) RefreshSecretKey() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RefreshSecretKey")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockJWTConfig_RefreshSecretKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshSecretKey'
type MockJWTConfig_RefreshSecretKey_Call struct {
	*mock.Call
}

// RefreshSecretKey is a helper method to define mock.On call
func (_e *MockJWTConfig_Expecter) RefreshSecretKey() *MockJWTConfig_RefreshSecretKey_Call {
	return &MockJWTConfig_RefreshSecretKey_Call{Call: _e.mock.On("RefreshSecretKey")}
}

func (_c *MockJWTConfig_RefreshSecretKey_Call) Run(run func()) *MockJWTConfig_RefreshSecretKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJWTConfig_RefreshSecretKey_Call) Return(_a0 []byte) *MockJWTConfig_RefreshSecretKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJWTConfig_RefreshSecretKey_Call) RunAndReturn(run func() []byte) *MockJWTConfig_RefreshSecretKey_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshTokenDuration provides a mock function with given fields:
func (_m *MockJWTConfig) RefreshTokenDuration() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RefreshTokenDuration")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// MockJWTConfig_RefreshTokenDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshTokenDuration'
type MockJWTConfig_RefreshTokenDuration_Call struct {
	*mock.Call
}

// RefreshTokenDuration is a helper method to define mock.On call
func (_e *MockJWTConfig_Expecter) RefreshTokenDuration() *MockJWTConfig_RefreshTokenDuration_Call {
	return &MockJWTConfig_RefreshTokenDuration_Call{Call: _e.mock.On("RefreshTokenDuration")}
}

func (_c *MockJWTConfig_RefreshTokenDuration_Call) Run(run func()) *MockJWTConfig_RefreshTokenDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJWTConfig_RefreshTokenDuration_Call) Return(_a0 time.Duration) *MockJWTConfig_RefreshTokenDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJWTConfig_RefreshTokenDuration_Call) RunAndReturn(run func() time.Duration) *MockJWTConfig_RefreshTokenDuration_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockJWTConfig creates a new instance of MockJWTConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockJWTConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockJWTConfig {
	mock := &MockJWTConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockdefserv

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/defany/auth-service/app/internal/model"
)

// MockUserService is an autogenerated mock type for the UserService type
type MockUserService struct {
	mock.Mock
}

type MockUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserService) EXPECT() *MockUserService_Expecter {
	return &MockUserService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, user
func (_m *MockUserService) Create(ctx context.Context, user model.UserCreate) (uint64, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserCreate) (uint64, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UserCreate) uint64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UserCreate) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - user model.UserCreate
func (_e *MockUserService_Expecter) Create(ctx interface{}, user interface{}) *MockUserService_Create_Call {
	return &MockUserService_Create_Call{Call: _e.mock.On("Create", ctx, user)}
}

func (_c *MockUserService_Create_Call) Run(run func(ctx context.Context, user model.UserCreate)) *MockUserService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserCreate))
	})
	return _c
}

func (_c *MockUserService_Create_Call) Return(_a0 uint64, _a1 error) *MockUserService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserService_Create_Call) RunAndReturn(run func(context.Context, model.UserCreate) (uint64, error)) *MockUserService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockUserService) Delete(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUserService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockUserService_Expecter) Delete(ctx interface{}, id interface{}) *MockUserService_Delete_Call {
	return &MockUserService_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockUserService_Delete_Call) Run(run func(ctx context.Context, id uint64)) *MockUserService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockUserService_Delete_Call) Return(_a0 error) *MockUserService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserService_Delete_Call) RunAndReturn(run func(context.Context, uint64) error) *MockUserService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockUserService) Get(ctx context.Context, id uint64) (model.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) model.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockUserService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockUserService_Expecter) Get(ctx interface{}, id interface{}) *MockUserService_Get_Call {
	return &MockUserService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockUserService_Get_Call) Run(run func(ctx context.Context, id uint64)) *MockUserService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockUserService_Get_Call) Return(_a0 model.User, _a1 error) *MockUserService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserService_Get_Call) RunAndReturn(run func(context.Context, uint64) (model.User, error)) *MockUserService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, user
func (_m *MockUserService) Update(ctx context.Context, user model.UserUpdate) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserUpdate) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUserService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - user model.UserUpdate
func (_e *MockUserService_Expecter) Update(ctx interface{}, user interface{}) *MockUserService_Update_Call {
	return &MockUserService_Update_Call{Call: _e.mock.On("Update", ctx, user)}
}

func (_c *MockUserService_Update_Call) Run(run func(ctx context.Context, user model.UserUpdate)) *MockUserService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserUpdate))
	})
	return _c
}

func (_c *MockUserService_Update_Call) Return(_a0 error) *MockUserService_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserService_Update_Call) RunAndReturn(run func(context.Context, model.UserUpdate) error) *MockUserService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserService creates a new instance of MockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserService {
	mock := &MockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

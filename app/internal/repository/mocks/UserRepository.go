// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockrepository

import (
	context "context"

	model "github.com/defany/auth-service/app/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) Create(ctx context.Context, user model.UserCreate) (uint64, error) {
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

// MockUserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - user model.UserCreate
func (_e *MockUserRepository_Expecter) Create(ctx interface{}, user interface{}) *MockUserRepository_Create_Call {
	return &MockUserRepository_Create_Call{Call: _e.mock.On("Create", ctx, user)}
}

func (_c *MockUserRepository_Create_Call) Run(run func(ctx context.Context, user model.UserCreate)) *MockUserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserCreate))
	})
	return _c
}

func (_c *MockUserRepository_Create_Call) Return(_a0 uint64, _a1 error) *MockUserRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_Create_Call) RunAndReturn(run func(context.Context, model.UserCreate) (uint64, error)) *MockUserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) Delete(ctx context.Context, id uint64) error {
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

// MockUserRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUserRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockUserRepository_Expecter) Delete(ctx interface{}, id interface{}) *MockUserRepository_Delete_Call {
	return &MockUserRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockUserRepository_Delete_Call) Run(run func(ctx context.Context, id uint64)) *MockUserRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockUserRepository_Delete_Call) Return(_a0 error) *MockUserRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Delete_Call) RunAndReturn(run func(context.Context, uint64) error) *MockUserRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DoesHaveAccess provides a mock function with given fields: ctx, userRole, endpoint
func (_m *MockUserRepository) DoesHaveAccess(ctx context.Context, userRole string, endpoint string) error {
	ret := _m.Called(ctx, userRole, endpoint)

	if len(ret) == 0 {
		panic("no return value specified for DoesHaveAccess")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userRole, endpoint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_DoesHaveAccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoesHaveAccess'
type MockUserRepository_DoesHaveAccess_Call struct {
	*mock.Call
}

// DoesHaveAccess is a helper method to define mock.On call
//   - ctx context.Context
//   - userRole string
//   - endpoint string
func (_e *MockUserRepository_Expecter) DoesHaveAccess(ctx interface{}, userRole interface{}, endpoint interface{}) *MockUserRepository_DoesHaveAccess_Call {
	return &MockUserRepository_DoesHaveAccess_Call{Call: _e.mock.On("DoesHaveAccess", ctx, userRole, endpoint)}
}

func (_c *MockUserRepository_DoesHaveAccess_Call) Run(run func(ctx context.Context, userRole string, endpoint string)) *MockUserRepository_DoesHaveAccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockUserRepository_DoesHaveAccess_Call) Return(_a0 error) *MockUserRepository_DoesHaveAccess_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_DoesHaveAccess_Call) RunAndReturn(run func(context.Context, string, string) error) *MockUserRepository_DoesHaveAccess_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) Update(ctx context.Context, user model.UserUpdate) error {
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

// MockUserRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUserRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - user model.UserUpdate
func (_e *MockUserRepository_Expecter) Update(ctx interface{}, user interface{}) *MockUserRepository_Update_Call {
	return &MockUserRepository_Update_Call{Call: _e.mock.On("Update", ctx, user)}
}

func (_c *MockUserRepository_Update_Call) Run(run func(ctx context.Context, user model.UserUpdate)) *MockUserRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserUpdate))
	})
	return _c
}

func (_c *MockUserRepository_Update_Call) Return(_a0 error) *MockUserRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Update_Call) RunAndReturn(run func(context.Context, model.UserUpdate) error) *MockUserRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// User provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) User(ctx context.Context, id uint64) (model.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for User")
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

// MockUserRepository_User_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'User'
type MockUserRepository_User_Call struct {
	*mock.Call
}

// User is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockUserRepository_Expecter) User(ctx interface{}, id interface{}) *MockUserRepository_User_Call {
	return &MockUserRepository_User_Call{Call: _e.mock.On("User", ctx, id)}
}

func (_c *MockUserRepository_User_Call) Run(run func(ctx context.Context, id uint64)) *MockUserRepository_User_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockUserRepository_User_Call) Return(_a0 model.User, _a1 error) *MockUserRepository_User_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_User_Call) RunAndReturn(run func(context.Context, uint64) (model.User, error)) *MockUserRepository_User_Call {
	_c.Call.Return(run)
	return _c
}

// UserByNickname provides a mock function with given fields: ctx, nickname
func (_m *MockUserRepository) UserByNickname(ctx context.Context, nickname string) (model.User, error) {
	ret := _m.Called(ctx, nickname)

	if len(ret) == 0 {
		panic("no return value specified for UserByNickname")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, nickname)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, nickname)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nickname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_UserByNickname_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserByNickname'
type MockUserRepository_UserByNickname_Call struct {
	*mock.Call
}

// UserByNickname is a helper method to define mock.On call
//   - ctx context.Context
//   - nickname string
func (_e *MockUserRepository_Expecter) UserByNickname(ctx interface{}, nickname interface{}) *MockUserRepository_UserByNickname_Call {
	return &MockUserRepository_UserByNickname_Call{Call: _e.mock.On("UserByNickname", ctx, nickname)}
}

func (_c *MockUserRepository_UserByNickname_Call) Run(run func(ctx context.Context, nickname string)) *MockUserRepository_UserByNickname_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_UserByNickname_Call) Return(_a0 model.User, _a1 error) *MockUserRepository_UserByNickname_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_UserByNickname_Call) RunAndReturn(run func(context.Context, string) (model.User, error)) *MockUserRepository_UserByNickname_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

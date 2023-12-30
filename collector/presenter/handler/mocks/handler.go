// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields:
func (_m *Handler) Handle() {
	_m.Called()
}

// Handler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type Handler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
func (_e *Handler_Expecter) Handle() *Handler_Handle_Call {
	return &Handler_Handle_Call{Call: _e.mock.On("Handle")}
}

func (_c *Handler_Handle_Call) Run(run func()) *Handler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_Handle_Call) Return() *Handler_Handle_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_Handle_Call) RunAndReturn(run func()) *Handler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

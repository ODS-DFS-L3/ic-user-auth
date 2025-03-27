// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// IOperatorHandler is an autogenerated mock type for the IOperatorHandler type
type IOperatorHandler struct {
	mock.Mock
}

// GetOperator provides a mock function with given fields: c
func (_m *IOperatorHandler) GetOperator(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetOperator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutOperator provides a mock function with given fields: c
func (_m *IOperatorHandler) PutOperator(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for PutOperator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIOperatorHandler creates a new instance of IOperatorHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIOperatorHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IOperatorHandler {
	mock := &IOperatorHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

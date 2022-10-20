// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/Sosial-Media-App/sosialta/features/contents/domain"
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// AddContent provides a mock function with given fields: newContent
func (_m *Services) AddContent(newContent domain.Core) (domain.Core, error) {
	ret := _m.Called(newContent)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newContent)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContent provides a mock function with given fields: id
func (_m *Services) DeleteContent(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtractToken provides a mock function with given fields: c
func (_m *Services) ExtractToken(c echo.Context) uint {
	ret := _m.Called(c)

	var r0 uint
	if rf, ok := ret.Get(0).(func(echo.Context) uint); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// GetContent provides a mock function with given fields:
func (_m *Services) GetContent() ([]domain.Core, error) {
	ret := _m.Called()

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func() []domain.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContentDetail provides a mock function with given fields: id
func (_m *Services) GetContentDetail(id uint) (domain.Core, error) {
	ret := _m.Called(id)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(uint) domain.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateContent provides a mock function with given fields: updateData
func (_m *Services) UpdateContent(updateData domain.Core) (domain.Core, error) {
	ret := _m.Called(updateData)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(updateData)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServices(t mockConstructorTestingTNewServices) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/Sosial-Media-App/sosialta/features/comments/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id_content
func (_m *Repository) Get(id_content uint) ([]domain.Core, error) {
	ret := _m.Called(id_content)

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func(uint) []domain.Core); ok {
		r0 = rf(id_content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id_content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newComment
func (_m *Repository) Insert(newComment domain.Core) (domain.Core, error) {
	ret := _m.Called(newComment)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newComment)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: newComment
func (_m *Repository) Update(newComment domain.Core) (domain.Core, error) {
	ret := _m.Called(newComment)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newComment)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

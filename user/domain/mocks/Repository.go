// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	userDomain "github.com/gozzafadillah/user/domain"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// BanUser provides a mock function with given fields: id, domain
func (_m *Repository) BanUser(id int, domain userDomain.Users) (userDomain.Users, error) {
	ret := _m.Called(id, domain)

	var r0 userDomain.Users
	if rf, ok := ret.Get(0).(func(int, userDomain.Users) userDomain.Users); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Get(0).(userDomain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, userDomain.Users) error); ok {
		r1 = rf(id, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id int) (userDomain.Users, error) {
	ret := _m.Called(id)

	var r0 userDomain.Users
	if rf, ok := ret.Get(0).(func(int) userDomain.Users); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(userDomain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *Repository) GetByUsername(username string) (int, error) {
	ret := _m.Called(username)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsernamePassword provides a mock function with given fields: username, password
func (_m *Repository) GetUsernamePassword(username string, password string) (userDomain.Users, error) {
	ret := _m.Called(username, password)

	var r0 userDomain.Users
	if rf, ok := ret.Get(0).(func(string, string) userDomain.Users); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(userDomain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: domain
func (_m *Repository) Save(domain userDomain.Users) (int, error) {
	ret := _m.Called(domain)

	var r0 int
	if rf, ok := ret.Get(0).(func(userDomain.Users) int); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(userDomain.Users) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t testing.TB) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

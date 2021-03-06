// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	userDomain "github.com/gozzafadillah/user/domain"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// BanUser provides a mock function with given fields: username
func (_m *Service) BanUser(username string) (userDomain.Users, error) {
	ret := _m.Called(username)

	var r0 userDomain.Users
	if rf, ok := ret.Get(0).(func(string) userDomain.Users); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(userDomain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetId provides a mock function with given fields: id
func (_m *Service) GetId(id int) (userDomain.Users, error) {
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

// GetUsername provides a mock function with given fields: username
func (_m *Service) GetUsername(username string) (userDomain.Users, error) {
	ret := _m.Called(username)

	var r0 userDomain.Users
	if rf, ok := ret.Get(0).(func(string) userDomain.Users); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(userDomain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertData provides a mock function with given fields: domain
func (_m *Service) InsertData(domain userDomain.Users) (userDomain.Users, error) {
	ret := _m.Called(domain)

	var r0 userDomain.Users
	if rf, ok := ret.Get(0).(func(userDomain.Users) userDomain.Users); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(userDomain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(userDomain.Users) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: username, password
func (_m *Service) Login(username string, password string) (string, error) {
	ret := _m.Called(username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t testing.TB) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

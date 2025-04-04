// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	authentication "authenticator-backend/domain/model/authentication"

	mock "github.com/stretchr/testify/mock"

	repository "authenticator-backend/domain/repository"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// ListAPIKeyOperators provides a mock function with given fields: param
func (_m *AuthRepository) ListAPIKeyOperators(param repository.APIKeyOperatorsParam) (authentication.APIKeyOperators, error) {
	ret := _m.Called(param)

	if len(ret) == 0 {
		panic("no return value specified for ListAPIKeyOperators")
	}

	var r0 authentication.APIKeyOperators
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.APIKeyOperatorsParam) (authentication.APIKeyOperators, error)); ok {
		return rf(param)
	}
	if rf, ok := ret.Get(0).(func(repository.APIKeyOperatorsParam) authentication.APIKeyOperators); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authentication.APIKeyOperators)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.APIKeyOperatorsParam) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAPIKeys provides a mock function with given fields: param
func (_m *AuthRepository) ListAPIKeys(param repository.APIKeysParam) (authentication.APIKeys, error) {
	ret := _m.Called(param)

	if len(ret) == 0 {
		panic("no return value specified for ListAPIKeys")
	}

	var r0 authentication.APIKeys
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.APIKeysParam) (authentication.APIKeys, error)); ok {
		return rf(param)
	}
	if rf, ok := ret.Get(0).(func(repository.APIKeysParam) authentication.APIKeys); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authentication.APIKeys)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.APIKeysParam) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCidrs provides a mock function with given fields: param
func (_m *AuthRepository) ListCidrs(param repository.APIKeyCidrsParam) (authentication.Cidrs, error) {
	ret := _m.Called(param)

	if len(ret) == 0 {
		panic("no return value specified for ListCidrs")
	}

	var r0 authentication.Cidrs
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.APIKeyCidrsParam) (authentication.Cidrs, error)); ok {
		return rf(param)
	}
	if rf, ok := ret.Get(0).(func(repository.APIKeyCidrsParam) authentication.Cidrs); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authentication.Cidrs)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.APIKeyCidrsParam) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

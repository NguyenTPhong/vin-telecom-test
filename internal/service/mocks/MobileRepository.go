// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "vinbigdata/internal/repository/entity"

	mock "github.com/stretchr/testify/mock"

	model "vinbigdata/internal/delivery/http/model"
)

// MobileRepository is an autogenerated mock type for the MobileRepository type
type MobileRepository struct {
	mock.Mock
}

// GetUserCallSummary provides a mock function with given fields: userName
func (_m *MobileRepository) GetUserCallSummary(userName string) (*model.BillingData, error) {
	ret := _m.Called(userName)

	var r0 *model.BillingData
	if rf, ok := ret.Get(0).(func(string) *model.BillingData); ok {
		r0 = rf(userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BillingData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUserCall provides a mock function with given fields: userCall
func (_m *MobileRepository) SaveUserCall(userCall *entity.UserCalls) error {
	ret := _m.Called(userCall)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.UserCalls) error); ok {
		r0 = rf(userCall)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMobileRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMobileRepository creates a new instance of MobileRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMobileRepository(t mockConstructorTestingTNewMobileRepository) *MobileRepository {
	mock := &MobileRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

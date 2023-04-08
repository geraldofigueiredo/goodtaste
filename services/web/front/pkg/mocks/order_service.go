// Code generated by mockery v2.23.2. DO NOT EDIT.

package mocks

import (
	entity "github.com/geraldofigueiredo/goodtaste/services/web/front/entity"
	mock "github.com/stretchr/testify/mock"
)

// OrderService is an autogenerated mock type for the OrderService type
type OrderService struct {
	mock.Mock
}

// GetOrder provides a mock function with given fields: id
func (_m *OrderService) GetOrder(id string) (*entity.Order, error) {
	ret := _m.Called(id)

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Order, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Order); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrder provides a mock function with given fields:
func (_m *OrderService) NewOrder() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewOrderService interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderService creates a new instance of OrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderService(t mockConstructorTestingTNewOrderService) *OrderService {
	mock := &OrderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
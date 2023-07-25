// Code generated by mockery v2.16.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/Amtelco-Software/ari/v6"
	mock "github.com/stretchr/testify/mock"
)

// Sender is an autogenerated mock type for the Sender type
type Sender struct {
	mock.Mock
}

// Send provides a mock function with given fields: e
func (_m *Sender) Send(e ari.Event) {
	_m.Called(e)
}

type mockConstructorTestingTNewSender interface {
	mock.TestingT
	Cleanup(func())
}

// NewSender creates a new instance of Sender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSender(t mockConstructorTestingTNewSender) *Sender {
	mock := &Sender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

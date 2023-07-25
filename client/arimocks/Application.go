// Code generated by mockery v2.16.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/Amtelco-Software/ari/v6"
	mock "github.com/stretchr/testify/mock"
)

// Application is an autogenerated mock type for the Application type
type Application struct {
	mock.Mock
}

// Data provides a mock function with given fields: key
func (_m *Application) Data(key *ari.Key) (*ari.ApplicationData, error) {
	ret := _m.Called(key)

	var r0 *ari.ApplicationData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.ApplicationData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ApplicationData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key
func (_m *Application) Get(key *ari.Key) *ari.ApplicationHandle {
	ret := _m.Called(key)

	var r0 *ari.ApplicationHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.ApplicationHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ApplicationHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields: _a0
func (_m *Application) List(_a0 *ari.Key) ([]*ari.Key, error) {
	ret := _m.Called(_a0)

	var r0 []*ari.Key
	if rf, ok := ret.Get(0).(func(*ari.Key) []*ari.Key); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ari.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: key, eventSource
func (_m *Application) Subscribe(key *ari.Key, eventSource string) error {
	ret := _m.Called(key, eventSource)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string) error); ok {
		r0 = rf(key, eventSource)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: key, eventSource
func (_m *Application) Unsubscribe(key *ari.Key, eventSource string) error {
	ret := _m.Called(key, eventSource)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string) error); ok {
		r0 = rf(key, eventSource)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewApplication interface {
	mock.TestingT
	Cleanup(func())
}

// NewApplication creates a new instance of Application. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApplication(t mockConstructorTestingTNewApplication) *Application {
	mock := &Application{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

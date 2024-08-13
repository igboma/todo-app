// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Task is an autogenerated mock type for the Task type
type Task struct {
	mock.Mock
}

// ID provides a mock function with given fields:
func (_m *Task) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Task) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PerformTask provides a mock function with given fields:
func (_m *Task) PerformTask() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PerformTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTask creates a new instance of Task. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTask(t interface {
	mock.TestingT
	Cleanup(func())
}) *Task {
	mock := &Task{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

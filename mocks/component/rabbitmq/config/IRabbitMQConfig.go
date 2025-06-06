// Code generated by mockery v2.23.0. DO NOT EDIT.

package mock_config

import mock "github.com/stretchr/testify/mock"

// IRabbitMQConfig is an autogenerated mock type for the IRabbitMQConfig type
type IRabbitMQConfig struct {
	mock.Mock
}

// GetDelay provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetDelay() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetExchangeDirect provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetExchangeDirect() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetExchangeDlx provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetExchangeDlx() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetExchangeJunk provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetExchangeJunk() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetExchangeKind provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetExchangeKind() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLimit provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetLimit() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetQueueDirect provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetQueueDirect() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetQueueDlq provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetQueueDlq() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetQueueJunk provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetQueueJunk() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRoutingKeyDirect provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetRoutingKeyDirect() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRoutingKeyDlx provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetRoutingKeyDlx() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRoutingKeyJunk provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetRoutingKeyJunk() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTtl provides a mock function with given fields:
func (_m *IRabbitMQConfig) GetTtl() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// IsEnabled provides a mock function with given fields:
func (_m *IRabbitMQConfig) IsEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewIRabbitMQConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRabbitMQConfig creates a new instance of IRabbitMQConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRabbitMQConfig(t mockConstructorTestingTNewIRabbitMQConfig) *IRabbitMQConfig {
	mock := &IRabbitMQConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

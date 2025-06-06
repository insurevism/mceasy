// Code generated by mockery v2.23.0. DO NOT EDIT.

package mock_producer

import (
	config "mceasy/internal/component/rabbitmq/config"

	mock "github.com/stretchr/testify/mock"
)

// Producer is an autogenerated mock type for the Producer type
type Producer struct {
	mock.Mock
}

// SendToDirect provides a mock function with given fields: _a0, message
func (_m *Producer) SendToDirect(_a0 config.IRabbitMQConfig, message []byte) (bool, error) {
	ret := _m.Called(_a0, message)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) (bool, error)); ok {
		return rf(_a0, message)
	}
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) bool); ok {
		r0 = rf(_a0, message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(config.IRabbitMQConfig, []byte) error); ok {
		r1 = rf(_a0, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendToDirectJsonMarshaled provides a mock function with given fields: _a0, message
func (_m *Producer) SendToDirectJsonMarshaled(_a0 config.IRabbitMQConfig, message interface{}) (bool, error) {
	ret := _m.Called(_a0, message)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, interface{}) (bool, error)); ok {
		return rf(_a0, message)
	}
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, interface{}) bool); ok {
		r0 = rf(_a0, message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(config.IRabbitMQConfig, interface{}) error); ok {
		r1 = rf(_a0, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendToDirectJsonWithIncrementDelay provides a mock function with given fields: _a0, message, incrementDelay
func (_m *Producer) SendToDirectJsonWithIncrementDelay(_a0 config.IRabbitMQConfig, message interface{}, incrementDelay int64) (bool, error) {
	ret := _m.Called(_a0, message, incrementDelay)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, interface{}, int64) (bool, error)); ok {
		return rf(_a0, message, incrementDelay)
	}
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, interface{}, int64) bool); ok {
		r0 = rf(_a0, message, incrementDelay)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(config.IRabbitMQConfig, interface{}, int64) error); ok {
		r1 = rf(_a0, message, incrementDelay)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendToJunk provides a mock function with given fields: _a0, message
func (_m *Producer) SendToJunk(_a0 config.IRabbitMQConfig, message []byte) (bool, error) {
	ret := _m.Called(_a0, message)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) (bool, error)); ok {
		return rf(_a0, message)
	}
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) bool); ok {
		r0 = rf(_a0, message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(config.IRabbitMQConfig, []byte) error); ok {
		r1 = rf(_a0, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProducer interface {
	mock.TestingT
	Cleanup(func())
}

// NewProducer creates a new instance of Producer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProducer(t mockConstructorTestingTNewProducer) *Producer {
	mock := &Producer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

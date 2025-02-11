// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	 sarama "github.com/IBM/sarama"
)

// SaramaConsumerGroupSession is an autogenerated mock type for the SaramaConsumerGroupSession type
type SaramaConsumerGroupSession struct {
	mock.Mock
}

// Claims provides a mock function with given fields:
func (_m *SaramaConsumerGroupSession) Claims() map[string][]int32 {
	ret := _m.Called()

	var r0 map[string][]int32
	if rf, ok := ret.Get(0).(func() map[string][]int32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]int32)
		}
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *SaramaConsumerGroupSession) Commit() {
	_m.Called()
}

// Context provides a mock function with given fields:
func (_m *SaramaConsumerGroupSession) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// GenerationID provides a mock function with given fields:
func (_m *SaramaConsumerGroupSession) GenerationID() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// MarkMessage provides a mock function with given fields: msg, metadata
func (_m *SaramaConsumerGroupSession) MarkMessage(msg *sarama.ConsumerMessage, metadata string) {
	_m.Called(msg, metadata)
}

// MarkOffset provides a mock function with given fields: topic, partition, offset, metadata
func (_m *SaramaConsumerGroupSession) MarkOffset(topic string, partition int32, offset int64, metadata string) {
	_m.Called(topic, partition, offset, metadata)
}

// MemberID provides a mock function with given fields:
func (_m *SaramaConsumerGroupSession) MemberID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResetOffset provides a mock function with given fields: topic, partition, offset, metadata
func (_m *SaramaConsumerGroupSession) ResetOffset(topic string, partition int32, offset int64, metadata string) {
	_m.Called(topic, partition, offset, metadata)
}

package mocks

import (
	context "context"

	"github.com/IBM/sarama"
	mock "github.com/stretchr/testify/mock"
	// sarama "github.com/Shopify/sarama"
)

// SaramaConsumerGroup is an autogenerated mock type for the SaramaConsumerGroup type
type SaramaConsumerGroup struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *SaramaConsumerGroup) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consume provides a mock function with given fields: ctx, topics, handler
func (_m *SaramaConsumerGroup) Consume(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) error {
	ret := _m.Called(ctx, topics, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, sarama.ConsumerGroupHandler) error); ok {
		r0 = rf(ctx, topics, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Errors provides a mock function with given fields:
func (_m *SaramaConsumerGroup) Errors() <-chan error {
	ret := _m.Called()

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func() <-chan error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}

// Pause provides a mock function with given fields:
func (_m *SaramaConsumerGroup) Pause(partitions map[string][]int32) {
	_m.Called(partitions)
}

func (_m *SaramaConsumerGroup) PauseAll() {
	_m.Called()
}

func (_m *SaramaConsumerGroup) Resume(partitions map[string][]int32) {
	_m.Called(partitions)
}

func (_m *SaramaConsumerGroup) ResumeAll() {
	_m.Called()
}

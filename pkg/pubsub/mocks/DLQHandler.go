// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	pubsub "github.com/sidaurukdedi/go-boiler/pkg/pubsub"
)

// DLQHandler is an autogenerated mock type for the DLQHandler type
type DLQHandler struct {
	mock.Mock
}

// Send provides a mock function with given fields: ctx, dlqMessage
func (_m *DLQHandler) Send(ctx context.Context, dlqMessage *pubsub.DeadLetterQueueMessage) error {
	ret := _m.Called(ctx, dlqMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pubsub.DeadLetterQueueMessage) error); ok {
		r0 = rf(ctx, dlqMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

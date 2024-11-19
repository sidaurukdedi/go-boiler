package pubsub_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/sidaurukdedi/go-boiler/pkg/pubsub"
	"github.com/sidaurukdedi/go-boiler/pkg/pubsub/mocks"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"go.elastic.co/apm"
)

func TestNewRedisSubscriberAdapter_Success(t *testing.T) {

	mockpubsub := &mocks.RedisPubSub{}
	mockhandle := &mocks.EventHandler{}
	subscriber := pubsub.NewRedisSubscriberAdapter(apm.DefaultTracer, logrus.New(), mockpubsub, mockhandle)
	messagemock := &redis.Message{
		Channel: "__keyevent@0__:expired",
		Payload: "payload test",
	}
	mockpubsub.On("ReceiveMessage", mock.Anything).Return(messagemock, nil)
	mockpubsub.On("Close").Return(nil)
	mockhandle.On("Handle", mock.Anything, mock.Anything).Return(nil)

	subscriber.Subscribe()
	<-time.After(time.Millisecond * 10)
	subscriber.Close()

	mockhandle.AssertExpectations(t)
	mockpubsub.AssertExpectations(t)

}
func TestNewRedisSubscriberAdapter_Error(t *testing.T) {

	mockpubsub := &mocks.RedisPubSub{}
	mockhandle := &mocks.EventHandler{}
	subscriber := pubsub.NewRedisSubscriberAdapter(apm.DefaultTracer, logrus.New(), mockpubsub, mockhandle)

	mockpubsub.On("ReceiveMessage", mock.Anything).Return(nil, redis.ErrClosed)
	mockpubsub.On("Close").Return(nil)

	subscriber.Subscribe()
	<-time.After(time.Millisecond * 10)
	subscriber.Close()

	mockpubsub.AssertExpectations(t)
}

func TestNewRedisSubscriberAdapter_ErrorHandler(t *testing.T) {
	mockpubsub := &mocks.RedisPubSub{}

	subscriber := pubsub.NewRedisSubscriberAdapter(apm.DefaultTracer, logrus.New(), mockpubsub, nil)
	messagemock := &redis.Message{
		Channel: "__keyevent@0__:expired",
		Payload: "payload test",
	}
	mockpubsub.On("ReceiveMessage", mock.Anything).Return(messagemock, nil)
	mockpubsub.On("Close").Return(nil)

	subscriber.Subscribe()
	<-time.After(time.Millisecond * 10)
	subscriber.Close()

	mockpubsub.AssertExpectations(t)
}
func TestNewRedisSubscriberAdapter_ErrorHandleMessage(t *testing.T) {

	mockpubsub := &mocks.RedisPubSub{}
	mockhandle := &mocks.EventHandler{}
	subscriber := pubsub.NewRedisSubscriberAdapter(apm.DefaultTracer, logrus.New(), mockpubsub, mockhandle)
	messagemock := &redis.Message{
		Channel: "__keyevent@0__:expired",
		Payload: "payload test",
	}
	mockpubsub.On("ReceiveMessage", mock.Anything).Return(messagemock, nil)
	mockpubsub.On("Close").Return(nil)
	mockhandle.On("Handle", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))

	subscriber.Subscribe()
	<-time.After(time.Millisecond * 10)
	subscriber.Close()

	mockhandle.AssertExpectations(t)
	mockpubsub.AssertExpectations(t)

}

func TestNewRedisSubscriberAdapter_ClosingError(t *testing.T) {
	mockpubsub := &mocks.RedisPubSub{}
	mockhandle := &mocks.EventHandler{}
	subscriber := pubsub.NewRedisSubscriberAdapter(apm.DefaultTracer, logrus.New(), mockpubsub, mockhandle)
	messagemock := &redis.Message{
		Channel: "__keyevent@0__:expired",
		Payload: "payload test",
	}
	mockpubsub.On("ReceiveMessage", mock.Anything).Return(messagemock, nil)
	mockpubsub.On("Close").Return(redis.ErrClosed)
	mockhandle.On("Handle", mock.Anything, mock.Anything).Return(nil)

	subscriber.Subscribe()
	<-time.After(time.Millisecond * 10)
	subscriber.Close()

	mockhandle.AssertExpectations(t)
	mockpubsub.AssertExpectations(t)
}

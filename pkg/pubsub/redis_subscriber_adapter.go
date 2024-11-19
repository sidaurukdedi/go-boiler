package pubsub

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
)

// RedisPubSub is a collection of behavior of a Redis PubSub
type RedisPubSub interface {
	ReceiveMessage(ctx context.Context) (*redis.Message, error)
	Close() error
}

type RedisSubscriberAdapter struct {
	tracer  *apm.Tracer
	logger  *logrus.Logger
	pubsub  RedisPubSub
	handler EventHandler
}

// NewRedisSubscriberAdapter is a constructor.
func NewRedisSubscriberAdapter(tracer *apm.Tracer, logger *logrus.Logger, pubsub RedisPubSub, handler EventHandler) Subscriber {
	return &RedisSubscriberAdapter{
		tracer:  tracer,
		logger:  logger,
		pubsub:  pubsub,
		handler: handler,
	}
}

func (rs *RedisSubscriberAdapter) Subscribe() {
	go func() {
		for {
			msg, err := rs.pubsub.ReceiveMessage(context.Background())
			if err != nil {
				rs.logger.Errorf("[redis] : %s", err.Error())
				return
			}
			txName := fmt.Sprintf("On Event: %s", msg.Channel)
			txType := "Redis Subscriber"
			txSuccess := "Success"

			tx := rs.tracer.StartTransaction(txName, txType)

			ctx := apm.ContextWithTransaction(context.Background(), tx)

			if rs.handler == nil {
				rs.logger.WithFields(logrus.Fields{
					"UnImplementEventHandler": "Event Handler is null",
				}).Info(msg.Payload)
				tx.Result = txSuccess
				tx.End()
				continue
			}

			if err := rs.handler.Handle(ctx, msg); err != nil {
				tx.Result = err.Error()
			}
			tx.End()
		}
	}()

}

// Close will stop the redis consumer
func (rs *RedisSubscriberAdapter) Close() (err error) {
	if err = rs.pubsub.Close(); err != nil {
		rs.logger.Errorf("[redis] %s", err.Error())
		return
	}
	rs.logger.Info("[redis] Subscribe is gracefully shut down.")
	return
}

package pubsub_test

import (
	"context"
	"testing"

	"github.com/sidaurukdedi/go-boiler/pkg/pubsub"

	// "github.com/Shopify/sarama/mocks"
	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
	"github.com/sirupsen/logrus"
)

func TestSaramaKafkaProducer(t *testing.T) {
	t.Run("when the message is successfully sent and delivered", func(t *testing.T) {
		saramaProducer := mocks.NewAsyncProducer(t, nil)
		saramaProducer.ExpectInputAndSucceed()

		publisher := pubsub.NewSaramaKafkaProducerAdapter(logrus.New(), &pubsub.SaramaKafkaProducerAdapterConfig{
			AsyncProducer: saramaProducer,
		})

		headers := pubsub.MessageHeaders{}
		headers.Add("test", "header")

		publisher.Send(context.TODO(), "test-topic", "test-key", headers, []byte("Hola"))
		publisher.Close()
	})

	t.Run("when the connection is timeout", func(t *testing.T) {
		saramaProducer := mocks.NewAsyncProducer(t, nil)
		saramaProducer.ExpectInputAndFail(sarama.ErrRequestTimedOut)

		publisher := pubsub.NewSaramaKafkaProducerAdapter(logrus.New(), &pubsub.SaramaKafkaProducerAdapterConfig{
			AsyncProducer: saramaProducer,
		})

		headers := pubsub.MessageHeaders{}
		headers.Add("test", "header")

		publisher.Send(context.TODO(), "test-topic", "test-key", headers, []byte("Hola"))
		publisher.Close()
	})

	t.Run("when the message channel is already closed", func(t *testing.T) {
		saramaProducer := mocks.NewAsyncProducer(t, nil)

		publisher := pubsub.NewSaramaKafkaProducerAdapter(logrus.New(), &pubsub.SaramaKafkaProducerAdapterConfig{
			AsyncProducer: saramaProducer,
		})

		headers := pubsub.MessageHeaders{}
		headers.Add("test", "header")

		publisher.Close()
		publisher.Send(context.TODO(), "test-topic", "test-key", headers, []byte("Hola"))
	})
}

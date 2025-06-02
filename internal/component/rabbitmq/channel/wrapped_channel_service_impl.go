package channel

import (
	"context"
	"hokusai/configs/rabbitmq/connection"

	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type WrappedChannelServiceImpl struct {
	connection *connection.RabbitMQConnection
}

func NewWrappedChannel(connection *connection.RabbitMQConnection) *WrappedChannelServiceImpl {
	return &WrappedChannelServiceImpl{connection: connection}
}

func (wc *WrappedChannelServiceImpl) PublishMessage(exchange, key string, msg amqp.Publishing) error {
	log.Debug("Performing additional actions before produce...")
	return wc.connection.GetChannelProduce().PublishWithContext(context.Background(), exchange, key, false, false, msg)
}

func (wc *WrappedChannelServiceImpl) ConsumeMessage(queue string) (<-chan amqp.Delivery, error) {
	log.Debug("Performing additional actions before consume...")

	// Set QoS for the consumer channel - prefetch count from config or default to 100
	prefetchCount := viper.GetInt("rabbitmq.consumer.prefetch_count")
	if prefetchCount == 0 {
		prefetchCount = 100 // Default value if not configured
	}

	// Set QoS/prefetch for the channel
	if err := wc.connection.GetChannelConsume().Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size (0 means no specific size limit)
		false,         // global (false means applied to this channel only)
	); err != nil {
		log.Errorf("Failed to set QoS: %v", err)
		return nil, err
	}

	return wc.connection.GetChannelConsume().Consume(queue, "", false, false, false, false, nil)
}

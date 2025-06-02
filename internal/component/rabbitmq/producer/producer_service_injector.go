//go:build wireinject
// +build wireinject

package producer

import (
	"hokusai/configs/rabbitmq/connection"
	"hokusai/internal/component/rabbitmq/channel"

	"github.com/google/wire"
)

var provider = wire.NewSet(
	NewProducerService,
	channel.NewWrappedChannel,

	wire.Bind(new(channel.WrappedChannelService), new(*channel.WrappedChannelServiceImpl)),
	wire.Bind(new(Producer), new(*ProducerServiceImpl)),
)

func InitializedProducer(connection *connection.RabbitMQConnection) *ProducerServiceImpl {
	wire.Build(provider)
	return nil
}

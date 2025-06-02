package registry

import (
	"mceasy/configs/rabbitmq/connection"
	"mceasy/ent"

	"github.com/go-redis/redis/v8"
)

type ConsumerRegistry struct {
	client *ent.Client
	conn   *connection.RabbitMQConnection
	redis  *redis.Client
}

func NewConsumerRegistry(client *ent.Client, redis *redis.Client, conn *connection.RabbitMQConnection) *ConsumerRegistry {
	return &ConsumerRegistry{client: client, conn: conn, redis: redis}
}

func (f *ConsumerRegistry) Register() {
	// Consumer registration can be added here when needed
	// For now, we don't have any consumers for the attendance management system

	// Example of how to register a consumer:
	// consumer := inbound.NewExampleConsumer(dbClient)
	// config.RegisterConsumer(ch, "queue_name", consumer.Consume)
}

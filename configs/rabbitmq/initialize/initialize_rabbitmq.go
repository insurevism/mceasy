package initialize

import (
	"hokusai/configs/rabbitmq/connection"
	"hokusai/configs/rabbitmq/recovery"
	"hokusai/ent"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
)

func RabbitMQInitialize(client *ent.Client, redis *redis.Client) *connection.RabbitMQConnection {

	newRabbitMQ := connection.NewRabbitMQ()
	_, err := newRabbitMQ.Connection()
	if err != nil {
		log.Errorf("Error closing RabbitMQConnection connection:", err)
	}

	rabbitConf := newRabbitMQ.GetConfig()
	if err != nil {
		log.Errorf("Error closing RabbitMQConnection connection: %v", err)
	}

	//for recovery reconnection RabbitMQ:
	go recovery.RabbitMQRecovery(client, redis, rabbitConf)

	return rabbitConf
}

package inbound

import "mceasy/internal/component/rabbitmq/config"

type Inbound interface {
	GetMessage(cfg config.IRabbitMQConfig) (bool, error)
}

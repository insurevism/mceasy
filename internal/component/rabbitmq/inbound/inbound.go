package inbound

import "hokusai/internal/component/rabbitmq/config"

type Inbound interface {
	GetMessage(cfg config.IRabbitMQConfig) (bool, error)
}

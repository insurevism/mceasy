package producer

import "mceasy/internal/component/rabbitmq/config"

type Producer interface {
	SendToDirect(producer config.IRabbitMQConfig, message []byte) (bool, error)
	SendToDirectJsonMarshaled(producer config.IRabbitMQConfig, message any) (bool, error)
	SendToDirectJsonWithIncrementDelay(producer config.IRabbitMQConfig, message any, incrementDelay int64) (bool, error)
	SendToJunk(producer config.IRabbitMQConfig, message []byte) (bool, error)
}

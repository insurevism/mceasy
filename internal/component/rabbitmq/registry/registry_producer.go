package registry

import (
	"mceasy/configs/rabbitmq/connection"
	"mceasy/internal/component/rabbitmq/config"
)

type ProducerRegistry struct {
	conn *connection.RabbitMQConnection
}

func NewProducerRegistry(conn *connection.RabbitMQConnection) *ProducerRegistry {
	return &ProducerRegistry{conn: conn}
}

func (f *ProducerRegistry) Register() {

	mqConfigs := []config.IRabbitMQConfig{

		config.NewRabbitMQConfigTickAggregrator(),
		config.NewRabbitMQConfigTickDBProcessor(),
		config.NewRabbitMQConfigTickV2Processor(),
	}

	// run registry:
	f.execute(mqConfigs)
}

package registry

import (
	"mceasy/internal/component/rabbitmq/config"

	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (f *ProducerRegistry) execute(mqConfigs []config.IRabbitMQConfig) {
	for _, mqConfig := range mqConfigs {
		if mqConfig.IsEnabled() {
			// Declare Exchange:
			err := f.buildExchange(mqConfig)
			if err != nil {
				log.Fatalf("Failed to declare an exchange: %v", err)
			}

			// Declare Queue:
			err = f.buildQueue(mqConfig)
			if err != nil {
				log.Fatalf("Failed to declare a queue: %v", err)
			}

			// Binding a Queue to an Exchange with a Routing Key:
			err = f.buildBind(mqConfig)
			if err != nil {
				log.Fatalf("Failed to bind the queue: %v", err)
			}

			log.Infof("`%s` exchange and `%s` queue has been successfully declared", mqConfig.GetExchangeDirect(), mqConfig.GetQueueDirect())
		}
	}
}

func (f *ProducerRegistry) buildExchange(producer config.IRabbitMQConfig) error {
	// Declare Exchange direct:
	err := f.conn.GetChannelProduce().ExchangeDeclare(
		producer.GetExchangeDirect(),
		"x-delayed-message",
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-delayed-type": producer.GetExchangeKind(),
		},
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange direct: %v", err)
	}

	// Declare Exchange DLX:
	err = f.conn.GetChannelProduce().ExchangeDeclare(
		producer.GetExchangeDlx(),
		producer.GetExchangeKind(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange direct: %v", err)
	}

	// Declare Exchange Junk:
	err = f.conn.GetChannelProduce().ExchangeDeclare(
		producer.GetExchangeJunk(),
		producer.GetExchangeKind(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange junk: %v", err)
	}

	return err
}

func (f *ProducerRegistry) buildQueue(producer config.IRabbitMQConfig) error {
	// Declare Queue direct:
	_, err := f.conn.GetChannelProduce().QueueDeclare(
		producer.GetQueueDirect(),
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    producer.GetExchangeDlx(),
			"x-dead-letter-routing-key": producer.GetRoutingKeyDlx(),
		},
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue direct: %v", err)
		return err
	}

	// Declare Queue DLQ:
	_, err = f.conn.GetChannelProduce().QueueDeclare(
		producer.GetQueueDlq(),
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    producer.GetExchangeDirect(),
			"x-dead-letter-routing-key": producer.GetRoutingKeyDirect(),
			"x-message-ttl":             producer.GetTtl(),
		},
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue dlq: %v", err)
		return err
	}

	// Declare Queue Junk:
	_, err = f.conn.GetChannelProduce().QueueDeclare(
		producer.GetQueueJunk(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue junk: %v", err)
		return err
	}

	return err
}

func (f *ProducerRegistry) buildBind(producer config.IRabbitMQConfig) error {
	// Binding a Queue to an Exchange with a Routing Key Direct:
	err := f.conn.GetChannelProduce().QueueBind(
		producer.GetQueueDirect(),
		producer.GetRoutingKeyDirect(),
		producer.GetExchangeDirect(),
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue direct: %v", err)
		return err
	}

	// Binding a Queue to an Exchange with a Routing Key Dlq:
	err = f.conn.GetChannelProduce().QueueBind(
		producer.GetQueueDlq(),
		producer.GetRoutingKeyDlx(),
		producer.GetExchangeDlx(),
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue dlq: %v", err)
		return err
	}

	// Binding a Queue to an Exchange with a Routing Key Junk:
	err = f.conn.GetChannelProduce().QueueBind(
		producer.GetQueueJunk(),
		producer.GetRoutingKeyJunk(),
		producer.GetExchangeJunk(),
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue junk: %v", err)
		return err
	}

	return err
}

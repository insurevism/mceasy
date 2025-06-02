package producer

import (
	"encoding/json"
	"mceasy/internal/component/rabbitmq/channel"
	"mceasy/internal/component/rabbitmq/config"
	"mceasy/internal/component/rabbitmq/errors"
	"mceasy/internal/component/rabbitmq/utils"

	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ProducerServiceImpl struct {
	ch channel.WrappedChannelService
}

func NewProducerService(ch channel.WrappedChannelService) *ProducerServiceImpl {
	return &ProducerServiceImpl{ch: ch}
}

func (p *ProducerServiceImpl) SendToDirect(producer config.IRabbitMQConfig, message []byte) (bool, error) {
	if !p.canPublishMessage(producer) {
		return false, errors.NewProducerError("can't publish a message to direct queue: %s; isEnabled: %t", producer.GetQueueDirect(), producer.IsEnabled())
	}

	// Send messages to Exchange:
	err := p.ch.PublishMessage(
		producer.GetExchangeDirect(),
		producer.GetRoutingKeyDirect(),
		amqp.Publishing{
			ContentType:  utils.GetContentType(),
			DeliveryMode: amqp.Persistent,
			Body:         message,
			Headers: amqp.Table{
				"x-delay": producer.GetDelay(),
			},
		},
	)
	if err != nil {
		log.Errorf("Failed to publish a message direct: %v", err)
		return false, err
	}

	// log.Infof("[%s] Sent a message direct: %s\n", producer.GetExchangeDirect(), message)
	return true, err
}

func (p *ProducerServiceImpl) SendToDirectJsonMarshaled(producer config.IRabbitMQConfig, message any) (bool, error) {

	payload, err := json.Marshal(message)
	if err != nil {
		log.Errorf("failed coverting data to json : %v", err)
		return false, err
	}

	return p.SendToDirect(producer, payload)
}

func (p *ProducerServiceImpl) SendToDirectJsonWithIncrementDelay(producer config.IRabbitMQConfig, message any, incrementDelay int64) (bool, error) {
	if !p.canPublishMessage(producer) {
		return false, errors.NewProducerError("can't publish a delayed message to direct queue: %s; isEnabled: %t", producer.GetQueueDirect(), producer.IsEnabled())
	}

	payload, err := json.Marshal(message)
	if err != nil {
		log.Errorf("failed coverting data to json : %v", err)
		return false, err
	}

	// Send messages to Exchange:
	err = p.ch.PublishMessage(
		producer.GetExchangeDirect(),
		producer.GetRoutingKeyDirect(),
		amqp.Publishing{
			ContentType:  utils.GetContentType(),
			DeliveryMode: amqp.Persistent,
			Body:         payload,
			Headers: amqp.Table{
				"x-delay": producer.GetDelay() * incrementDelay,
			},
		},
	)
	if err != nil {
		log.Errorf("Failed to publish a message direct with increment delay: %v", err)
		return false, err
	}

	// log.Infof("[%s] Sent a message direct with increment delay: %s\n", producer.GetExchangeDirect(), message)
	return true, err
}

func (p *ProducerServiceImpl) SendToJunk(producer config.IRabbitMQConfig, message []byte) (bool, error) {
	if !p.canPublishMessage(producer) {
		return false, errors.NewProducerError("can't publish a message to junk queue: %s; isEnabled: %t", producer.GetQueueDirect(), producer.IsEnabled())
	}

	// Send messages to Exchange:
	err := p.ch.PublishMessage(
		producer.GetExchangeJunk(),
		producer.GetRoutingKeyJunk(),
		amqp.Publishing{
			ContentType:  utils.GetContentType(),
			DeliveryMode: amqp.Persistent,
			Body:         message,
		},
	)
	if err != nil {
		log.Errorf("Failed to publish a message junk: %v", err)
		return false, err
	}

	log.Infof("[%s] Sent a message junk: %s\n", producer.GetExchangeJunk(), message)
	return true, err
}

func (*ProducerServiceImpl) canPublishMessage(producerConfig config.IRabbitMQConfig) bool {
	return producerConfig.IsEnabled()
}

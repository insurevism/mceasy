package inbound

import (
	"context"
	"mceasy/configs/rabbitmq/connection"
	"mceasy/internal/component/rabbitmq/channel"
	"mceasy/internal/component/rabbitmq/config"
	"mceasy/internal/component/rabbitmq/consumer"
	"mceasy/internal/component/rabbitmq/producer"
	"mceasy/internal/component/rabbitmq/utils"

	"runtime"

	"github.com/labstack/gommon/log"
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type Retriable[Data any] struct {
	channel  channel.WrappedChannelService
	producer producer.Producer
	consumer consumer.Consumer[Data]
}

func NewRetriable[Data any](
	connection *connection.RabbitMQConnection,
	consumer consumer.Consumer[Data],
) *Retriable[Data] {

	wrapped := channel.NewWrappedChannel(connection)

	return &Retriable[Data]{
		channel:  wrapped,
		producer: producer.NewProducerService(wrapped),
		consumer: consumer,
	}
}

func (t *Retriable[Data]) sendToJunk(cfg config.IRabbitMQConfig, message amqp091.Delivery) {
	if !cfg.IsEnabled() {
		return
	}

	ok, err := t.producer.SendToJunk(cfg, message.Body)

	if err != nil {
		log.Warnf("Failed to publish a message junk: error %v", err)
		return
	}

	if !ok {
		log.Warn("Failed to publish a message junk: returning false")
	}
}

func (t *Retriable[Data]) GetMessage(cfg config.IRabbitMQConfig) (bool, error) {
	if !cfg.IsEnabled() {
		return false, nil
	}

	queue := cfg.GetQueueDirect()
	message, errConsume := t.channel.ConsumeMessage(queue)

	if errConsume != nil {
		log.Errorf("Failed to consume messages: %v", errConsume)
		return false, errConsume
	}

	// Get worker count from config or default to number of CPUs
	workerCount := viper.GetInt("rabbitmq.consumer.worker_count")
	if workerCount == 0 {
		workerCount = runtime.NumCPU() // Default to number of available CPUs
	}

	// Create a channel for distributing work
	tasks := make(chan amqp091.Delivery, 1000) // Buffer for message handling

	// Start worker pool
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			for msg := range tasks {
				t.processMessage(cfg, msg, workerID)
			}
		}(i)
	}

	// Collect messages from RabbitMQ and distribute to workers
	go func() {
		for msg := range message {
			// Send to worker pool
			tasks <- msg
		}
		// Close tasks channel if RabbitMQ connection closes
		close(tasks)
	}()

	log.Infof("ready to consume incoming messages from `%s` queue with %d workers", cfg.GetQueueDirect(), workerCount)
	return true, errConsume
}

// Process individual messages
func (t *Retriable[Data]) processMessage(cfg config.IRabbitMQConfig, msg amqp091.Delivery, workerID int) {
	count := utils.CheckLimitRetry(msg)

	// Process message here:
	data, err := t.consumer.ParseData(msg.Body)
	if err != nil {
		log.Errorf("Failed to parse data [%s]: %v", cfg.GetQueueDirect(), err)

		err = msg.Ack(false)
		if err != nil {
			log.Warnf("Failed acknowledge message [%s]: %v", cfg.GetQueueDirect(), err)
		}

		t.sendToJunk(cfg, msg)
		return
	}

	err = t.consumer.Handle(context.Background(), data)
	if err != nil {
		log.Warnf("Failed to process service [%s]: %v", cfg.GetQueueDirect(), err)
	} else {
		err = msg.Ack(false)
	}

	if err != nil {
		isHasExceeded := utils.IsHasExceeded(cfg.GetLimit(), count, msg)
		if isHasExceeded {
			t.sendToJunk(cfg, msg)
		}
	}
}

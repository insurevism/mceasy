package connection

import (
	"fmt"
	"mceasy/configs/credential"

	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConnection struct {
	name           string
	conn           *amqp.Connection
	channelProduce *amqp.Channel
	channelConsume *amqp.Channel
	err            chan error
}

func NewRabbitMQ() *RabbitMQConnection {
	c := &RabbitMQConnection{
		err: make(chan error),
	}
	return c
}

func (c *RabbitMQConnection) Connection() (*amqp.Connection, error) {

	config := fmt.Sprintf("amqp://%s:%s@%s:%s",
		credential.GetString("rabbitmq.configs.username"),
		credential.GetString("rabbitmq.configs.password"),
		credential.GetString("rabbitmq.configs.host"),
		credential.GetString("rabbitmq.configs.port"))

	log.Infof("Connecting to RabbitMQ: %s", config)

	var err error
	c.conn, err = amqp.Dial(config)
	if err != nil {
		log.Errorf("Failed to connect to RabbitMQConnection: %v", err)
		return nil, err
	}

	c.channelProduce, err = c.conn.Channel()
	if err != nil {
		log.Errorf("Failed to open a producer channel: %v", err)
	}

	c.channelConsume, err = c.conn.Channel()
	if err != nil {
		log.Errorf("Failed to open a consumer channel: %v", err)
	}

	return c.conn, nil
}

func (c *RabbitMQConnection) ChannelRabbitMQ(conn *amqp.Connection) *amqp.Channel {
	var err error
	c.channelProduce, err = conn.Channel()
	if err != nil {
		log.Errorf("Failed to open a channel: %v", err)
	}

	return c.channelProduce
}

func (c *RabbitMQConnection) Reconnect() error {
	if _, err := c.Connection(); err != nil {
		return err
	}

	return nil
}

func (c *RabbitMQConnection) GetConfig() *RabbitMQConnection {
	return c
}

func (c *RabbitMQConnection) GetConnection() *amqp.Connection {
	return c.conn
}

func (c *RabbitMQConnection) GetChannelProduce() *amqp.Channel {
	return c.channelProduce
}

func (c *RabbitMQConnection) GetChannelConsume() *amqp.Channel {
	return c.channelConsume
}

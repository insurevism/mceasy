package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewRabbitMQConfigTickDBProcessor() IRabbitMQConfig {

	prefix := fmt.Sprintf("mceasy-%s-", viper.GetString("application.name"))

	return &RabbitMQConfig{
		Enabled: viper.GetBool("rabbitmq.mceasy.tickDbProcessor.enabled"),

		ExchangeKind: viper.GetString("rabbitmq.mceasy.tickDbProcessor.exchangeKind"),

		ExchangeDirect:   prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.exchangeDirect"),
		QueueDirect:      prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.queueDirect"),
		RoutingKeyDirect: prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.routingKeyDirect"),

		ExchangeDlx:   prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.exchangeDlx"),
		QueueDlq:      prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.queueDlq"),
		RoutingKeyDlx: prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.routingKeyDlx"),

		ExchangeJunk:   prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.exchangeJunk"),
		QueueJunk:      prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.queueJunk"),
		RoutingKeyJunk: prefix + viper.GetString("rabbitmq.mceasy.tickDbProcessor.routingKeyJunk"),

		Ttl:   viper.GetInt64("rabbitmq.mceasy.tickDbProcessor.ttl"),
		Delay: viper.GetInt64("rabbitmq.mceasy.tickDbProcessor.delay"),
		Limit: viper.GetInt64("rabbitmq.mceasy.tickDbProcessor.limit"),
	}
}

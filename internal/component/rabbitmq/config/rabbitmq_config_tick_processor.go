package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewRabbitMQConfigTickV2Processor() IRabbitMQConfig {

	prefix := fmt.Sprintf("mainhaus-%s-", viper.GetString("application.name"))

	return &RabbitMQConfig{
		Enabled: viper.GetBool("rabbitmq.mceasy.tickProcessor.enabled"),

		ExchangeKind: viper.GetString("rabbitmq.mceasy.tickProcessor.exchangeKind"),

		ExchangeDirect:   prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.exchangeDirect"),
		QueueDirect:      prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.queueDirect"),
		RoutingKeyDirect: prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.routingKeyDirect"),

		ExchangeDlx:   prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.exchangeDlx"),
		QueueDlq:      prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.queueDlq"),
		RoutingKeyDlx: prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.routingKeyDlx"),

		ExchangeJunk:   prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.exchangeJunk"),
		QueueJunk:      prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.queueJunk"),
		RoutingKeyJunk: prefix + viper.GetString("rabbitmq.mceasy.tickProcessor.routingKeyJunk"),

		Ttl:   viper.GetInt64("rabbitmq.mceasy.tickProcessor.ttl"),
		Delay: viper.GetInt64("rabbitmq.mceasy.tickProcessor.delay"),
		Limit: viper.GetInt64("rabbitmq.mceasy.tickProcessor.limit"),
	}
}

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewRabbitMQConfigTickV2Processor() IRabbitMQConfig {

	prefix := fmt.Sprintf("mainhaus-%s-", viper.GetString("application.name"))

	return &RabbitMQConfig{
		Enabled: viper.GetBool("rabbitmq.hokusai.tickProcessor.enabled"),

		ExchangeKind: viper.GetString("rabbitmq.hokusai.tickProcessor.exchangeKind"),

		ExchangeDirect:   prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.exchangeDirect"),
		QueueDirect:      prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.queueDirect"),
		RoutingKeyDirect: prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.routingKeyDirect"),

		ExchangeDlx:   prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.exchangeDlx"),
		QueueDlq:      prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.queueDlq"),
		RoutingKeyDlx: prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.routingKeyDlx"),

		ExchangeJunk:   prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.exchangeJunk"),
		QueueJunk:      prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.queueJunk"),
		RoutingKeyJunk: prefix + viper.GetString("rabbitmq.hokusai.tickProcessor.routingKeyJunk"),

		Ttl:   viper.GetInt64("rabbitmq.hokusai.tickProcessor.ttl"),
		Delay: viper.GetInt64("rabbitmq.hokusai.tickProcessor.delay"),
		Limit: viper.GetInt64("rabbitmq.hokusai.tickProcessor.limit"),
	}
}

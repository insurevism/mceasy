package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewRabbitMQConfigTickDBProcessor() IRabbitMQConfig {

	prefix := fmt.Sprintf("mainhaus-%s-", viper.GetString("application.name"))

	return &RabbitMQConfig{
		Enabled: viper.GetBool("rabbitmq.hokusai.tickDbProcessor.enabled"),

		ExchangeKind: viper.GetString("rabbitmq.hokusai.tickDbProcessor.exchangeKind"),

		ExchangeDirect:   prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.exchangeDirect"),
		QueueDirect:      prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.queueDirect"),
		RoutingKeyDirect: prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.routingKeyDirect"),

		ExchangeDlx:   prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.exchangeDlx"),
		QueueDlq:      prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.queueDlq"),
		RoutingKeyDlx: prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.routingKeyDlx"),

		ExchangeJunk:   prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.exchangeJunk"),
		QueueJunk:      prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.queueJunk"),
		RoutingKeyJunk: prefix + viper.GetString("rabbitmq.hokusai.tickDbProcessor.routingKeyJunk"),

		Ttl:   viper.GetInt64("rabbitmq.hokusai.tickDbProcessor.ttl"),
		Delay: viper.GetInt64("rabbitmq.hokusai.tickDbProcessor.delay"),
		Limit: viper.GetInt64("rabbitmq.hokusai.tickDbProcessor.limit"),
	}
}

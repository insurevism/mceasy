package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewRabbitMQConfigTickAggregrator() IRabbitMQConfig {

	prefix := fmt.Sprintf("mainhaus-%s-", viper.GetString("application.name"))

	return &RabbitMQConfig{
		Enabled: viper.GetBool("rabbitmq.hokusai.tickAggregrator.enabled"),

		ExchangeKind: viper.GetString("rabbitmq.hokusai.tickAggregrator.exchangeKind"),

		ExchangeDirect:   prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.exchangeDirect"),
		QueueDirect:      prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.queueDirect"),
		RoutingKeyDirect: prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.routingKeyDirect"),

		ExchangeDlx:   prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.exchangeDlx"),
		QueueDlq:      prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.queueDlq"),
		RoutingKeyDlx: prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.routingKeyDlx"),

		ExchangeJunk:   prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.exchangeJunk"),
		QueueJunk:      prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.queueJunk"),
		RoutingKeyJunk: prefix + viper.GetString("rabbitmq.hokusai.tickAggregrator.routingKeyJunk"),

		Ttl:   viper.GetInt64("rabbitmq.hokusai.tickAggregrator.ttl"),
		Delay: viper.GetInt64("rabbitmq.hokusai.tickAggregrator.delay"),
		Limit: viper.GetInt64("rabbitmq.hokusai.tickAggregrator.limit"),
	}
}

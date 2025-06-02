package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewRabbitMQConfigTickAggregrator() IRabbitMQConfig {

	prefix := fmt.Sprintf("mainhaus-%s-", viper.GetString("application.name"))

	return &RabbitMQConfig{
		Enabled: viper.GetBool("rabbitmq.mceasy.tickAggregrator.enabled"),

		ExchangeKind: viper.GetString("rabbitmq.mceasy.tickAggregrator.exchangeKind"),

		ExchangeDirect:   prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.exchangeDirect"),
		QueueDirect:      prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.queueDirect"),
		RoutingKeyDirect: prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.routingKeyDirect"),

		ExchangeDlx:   prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.exchangeDlx"),
		QueueDlq:      prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.queueDlq"),
		RoutingKeyDlx: prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.routingKeyDlx"),

		ExchangeJunk:   prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.exchangeJunk"),
		QueueJunk:      prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.queueJunk"),
		RoutingKeyJunk: prefix + viper.GetString("rabbitmq.mceasy.tickAggregrator.routingKeyJunk"),

		Ttl:   viper.GetInt64("rabbitmq.mceasy.tickAggregrator.ttl"),
		Delay: viper.GetInt64("rabbitmq.mceasy.tickAggregrator.delay"),
		Limit: viper.GetInt64("rabbitmq.mceasy.tickAggregrator.limit"),
	}
}

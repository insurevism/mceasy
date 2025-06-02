package recovery

import (
	"mceasy/configs/rabbitmq/connection"
	"mceasy/ent"
	"mceasy/internal/component/rabbitmq/registry"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func RabbitMQRecovery(client *ent.Client, redis *redis.Client, rabbitConf *connection.RabbitMQConnection) {
	go func() {
		for {
			reason, ok := <-rabbitConf.GetConnection().NotifyClose(make(chan *amqp.Error))
			if !ok {
				log.Errorf("connection closed")
				break
			}
			log.Errorf("connection closed, reason: %v", reason)

			for {
				//time sleep for waiting connection up:
				timeRecovery := viper.GetInt("rabbitmq.configs.recovery")
				time.Sleep(time.Duration(timeRecovery) * time.Second)

				err := rabbitConf.Reconnect()
				if err == nil {
					log.Infof("reconnect rabbitmq success")

					//rabbitmq registry exchange, queue, dlq and other:
					registerMq := registry.NewProducerRegistry(rabbitConf)
					registerMq.Register()

					//rabbitmq registry consumer:
					registerConsumer := registry.NewConsumerRegistry(client, redis, rabbitConf)
					registerConsumer.Register()

					break
				}

				log.Errorf("reconnect rabbitmq failed, err: %v", err)
			}

		}
	}()
}

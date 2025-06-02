package registry

import (
	"hokusai/configs/rabbitmq/connection"
	"hokusai/ent"

	"hokusai/internal/applications/tick"
	tickDto "hokusai/internal/applications/tick/dto"

	tickAggregationDto "hokusai/internal/applications/tick_aggregation/dto"
	"hokusai/internal/applications/tick_v2"

	"hokusai/internal/component/rabbitmq/config"
	inbound "hokusai/internal/component/rabbitmq/inbound"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
)

type ConsumerRegistry struct {
	client *ent.Client
	conn   *connection.RabbitMQConnection
	redis  *redis.Client
}

func NewConsumerRegistry(client *ent.Client, redis *redis.Client, conn *connection.RabbitMQConnection) *ConsumerRegistry {
	return &ConsumerRegistry{client: client, conn: conn, redis: redis}
}

func (f *ConsumerRegistry) Register() {

	//init testing inbound:
	{
		//inbound := example.InitializedExampleInbound(f.client, f.conn)
		//_, err := inbound.GetMessage(config.NewRabbitMQConfigExample())
		//if err != nil {
		//	log.Fatalf("Failed to consume messages: %v", err)
		//}
	}

	//init other consumer here...

	{
		cnf := config.NewRabbitMQConfigTickAggregrator()
		consumer := tick.InitializedTickAggregrator(f.client, f.redis, f.conn)
		_, err := inbound.NewRetriable[*tickDto.ForexDataRequest](f.conn, consumer).GetMessage(cnf)
		if err != nil {
			log.Fatalf("Failed to consume  + Sync Integration messages: %v", err)
		}
	}

	{
		cnf := config.NewRabbitMQConfigTickDBProcessor()
		consumer := tick.InitializedTickV4DbProcessor(f.client, f.redis, f.conn)
		_, err := inbound.NewRetriable[*tickDto.ForexDataRequest](f.conn, consumer).GetMessage(cnf)
		if err != nil {
			log.Fatalf("Failed to consume  + Sync Integration messages: %v", err)
		}
	}

	{
		cnf := config.NewRabbitMQConfigTickV2Processor()
		consumer := tick_v2.InitializedTickV2Processor(f.client, f.redis, f.conn)
		_, err := inbound.NewRetriable[*tickAggregationDto.AggregationResultDTO](f.conn, consumer).GetMessage(cnf)
		if err != nil {
			log.Fatalf("Failed to consume  + Sync Integration messages: %v", err)
		}
	}

}

//go:generate wire
//go:build wireinject
// +build wireinject

package salary

import (
	"mceasy/ent"
	"mceasy/internal/applications/salary/repository"
	"mceasy/internal/applications/salary/service"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerSalary = wire.NewSet(
	repository.NewSalaryRepository,
	transaction.NewTrx,
	service.NewSalaryService,
	cache.NewCache,

	wire.Bind(new(repository.SalaryRepository), new(*repository.SalaryRepositoryImpl)),
	wire.Bind(new(transaction.Trx), new(*transaction.TrxImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
	wire.Bind(new(service.SalaryService), new(*service.SalaryServiceImpl)),
)

func InitializedSalaryService(dbClient *ent.Client, redisClient *redis.Client) *service.SalaryServiceImpl {
	wire.Build(providerSalary)
	return nil
}

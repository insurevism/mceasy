//go:build wireinject
// +build wireinject

package employee

import (
	"mceasy/ent"
	"mceasy/internal/applications/employee/repository"
	"mceasy/internal/applications/employee/service"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerEmployee = wire.NewSet(
	repository.NewEmployeeRepository,
	transaction.NewTrx,
	service.NewEmployeeService,
	cache.NewCache,

	wire.Bind(new(repository.EmployeeRepository), new(*repository.EmployeeRepositoryImpl)),
	wire.Bind(new(transaction.Trx), new(*transaction.TrxImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
	wire.Bind(new(service.EmployeeService), new(*service.EmployeeServiceImpl)),
)

func InitializedEmployeeService(dbClient *ent.Client, redisClient *redis.Client) *service.EmployeeServiceImpl {
	wire.Build(providerEmployee)
	return nil
}

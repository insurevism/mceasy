//go:build wireinject
// +build wireinject

package health

import (
	"mceasy/ent"
	"mceasy/internal/applications/health/repository"
	"mceasy/internal/applications/health/service"
	"mceasy/internal/component/cache"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerSetHealth = wire.NewSet(
	repository.NewHealthRepository,
	service.NewHealthService,
	cache.NewCache,
	wire.Bind(new(repository.HealthRepository), new(*repository.HealthRepositoryImpl)),
	wire.Bind(new(service.HealthService), new(*service.HealthServiceImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
)

func InitializeHealthService(dbClient *ent.Client, cacheClient *redis.Client) *service.HealthServiceImpl {
	wire.Build(providerSetHealth)
	return nil
}

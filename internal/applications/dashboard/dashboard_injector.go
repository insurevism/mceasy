//go:generate wire
//go:build wireinject
// +build wireinject

package dashboard

import (
	"mceasy/ent"
	"mceasy/internal/applications/dashboard/controller"
	"mceasy/internal/applications/dashboard/repository"
	"mceasy/internal/applications/dashboard/service"
	"mceasy/internal/component/cache"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerDashboard = wire.NewSet(
	repository.NewDashboardRepository,
	service.NewDashboardService,
	cache.NewCache,

	wire.Bind(new(repository.DashboardRepository), new(*repository.DashboardRepositoryImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
	wire.Bind(new(service.DashboardService), new(*service.DashboardServiceImpl)),
)

func InitializedDashboardService(dbClient *ent.Client, redisClient *redis.Client) service.DashboardService {
	wire.Build(providerDashboard)
	return nil
}

func InitializedDashboardController(dbClient *ent.Client, redisClient *redis.Client) *controller.DashboardController {
	wire.Build(
		providerDashboard,
		controller.NewDashboardController,
	)
	return nil
}

//go:build wireinject
// +build wireinject

package user

import (
	"mceasy/ent"
	"mceasy/internal/applications/user/repository"
	"mceasy/internal/applications/user/service"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerUser = wire.NewSet(
	repository.NewUserRepository,
	transaction.NewTrx,
	service.NewUserService,
	cache.NewCache,

	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Bind(new(transaction.Trx), new(*transaction.TrxImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

func InitializedUserService(dbClient *ent.Client, redisClient *redis.Client) *service.UserServiceImpl {
	wire.Build(providerUser)
	return nil
}

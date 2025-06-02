//go:build wireinject
// +build wireinject

package auth

import (
	"mceasy/ent"
	"mceasy/internal/applications/auth/service"
	passwordhasher "mceasy/internal/applications/auth/utils/password_hasher"
	userRepository "mceasy/internal/applications/user/repository"
	userService "mceasy/internal/applications/user/service"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerSetAuth = wire.NewSet(
	// services
	service.NewAuthService,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),

	// user service dependencies
	userService.NewUserService,
	wire.Bind(new(userService.UserService), new(*userService.UserServiceImpl)),

	// user repository
	userRepository.NewUserRepository,
	wire.Bind(new(userRepository.UserRepository), new(*userRepository.UserRepositoryImpl)),

	// transaction and cache
	transaction.NewTrx,
	wire.Bind(new(transaction.Trx), new(*transaction.TrxImpl)),
	cache.NewCache,
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),

	// password hasher
	passwordhasher.NewBcryptHasher,
	wire.Bind(new(passwordhasher.PasswordHasher), new(*passwordhasher.BcryptHasher)),
)

func InitializedAuthService(dbClient *ent.Client, redis *redis.Client) *service.AuthServiceImpl {
	wire.Build(providerSetAuth)
	return nil
}

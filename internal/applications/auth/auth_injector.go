//go:build wireinject
// +build wireinject

package auth

import (
	"mceasy/ent"
	"mceasy/internal/applications/auth/service"
	passwordhasher "mceasy/internal/applications/auth/utils/password_hasher"
	clientcredentialdb "mceasy/internal/applications/auth_client_credential/repository/db"
	clientsession "mceasy/internal/applications/auth_client_session"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerSetAuth = wire.NewSet(
	// services
	service.NewAuthService,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),

	// repositories
	clientcredentialdb.NewClientCredentialRepository,
	wire.Bind(new(clientcredentialdb.ClientCredentialRepository), new(*clientcredentialdb.ClientCredentialRepositoryImpl)),

	// etc
	passwordhasher.NewBcryptHasher,
	wire.Bind(new(passwordhasher.PasswordHasher), new(*passwordhasher.BcryptHasher)),
)

func InitializedAuthService(dbClient *ent.Client, redis *redis.Client) *service.AuthServiceImpl {
	panic(wire.Build(providerSetAuth, clientsession.ProviderSetClientSession))
}

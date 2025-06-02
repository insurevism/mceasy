//go:generate wire
//go:build wireinject
// +build wireinject

package attendance

import (
	"mceasy/ent"
	"mceasy/internal/applications/attendance/repository"
	"mceasy/internal/applications/attendance/service"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var providerAttendance = wire.NewSet(
	repository.NewAttendanceRepository,
	transaction.NewTrx,
	service.NewAttendanceService,
	cache.NewCache,

	wire.Bind(new(repository.AttendanceRepository), new(*repository.AttendanceRepositoryImpl)),
	wire.Bind(new(transaction.Trx), new(*transaction.TrxImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
	wire.Bind(new(service.AttendanceService), new(*service.AttendanceServiceImpl)),
)

func InitializedAttendanceService(dbClient *ent.Client, redisClient *redis.Client) *service.AttendanceServiceImpl {
	wire.Build(providerAttendance)
	return nil
}
 
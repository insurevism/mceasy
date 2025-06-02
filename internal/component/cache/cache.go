package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Ping(ctx context.Context) error
	Create(ctx context.Context, key string, data interface{}, expiration time.Duration) (bool, error)
	Get(ctx context.Context, key string, data interface{}) (interface{}, error)
	Delete(ctx context.Context, key string) (bool, error)
	GetOrCreate(ctx context.Context, key string, expiration time.Duration, factory func() (interface{}, error), data interface{}) (interface{}, error)
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	CreateWithNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	UpdateWithOptimisticLock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	CreateRawString(ctx context.Context, key string, value string, expiration time.Duration) (bool, error)
	GetRawString(ctx context.Context, key string) (string, error)
	RawUpdateWithOptimisticLock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	GetRedisClient() *redis.Client
}

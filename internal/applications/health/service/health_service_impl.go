package service

import (
	"context"
	"mceasy/internal/applications/health/repository"
	"mceasy/internal/component/cache"
)

type HealthServiceImpl struct {
	repository repository.HealthRepository
	cache      cache.Cache
}

func NewHealthService(repository repository.HealthRepository, cache cache.Cache) *HealthServiceImpl {
	return &HealthServiceImpl{
		repository: repository,
		cache:      cache,
	}
}

func (s *HealthServiceImpl) Health(ctx context.Context, message string) (map[string]string, error) {
	messageService := message + "hello from service layer "
	result, errRepo := s.repository.Health(ctx, messageService)

	errCache := s.cache.Ping(ctx)
	if errCache != nil {
		result["cache_status"] = "DOWN"
		result["cache_name"] = "redis"
	} else {
		result["cache_status"] = "UP"
		result["cache_name"] = "redis"
	}

	//database will have first err priority
	if errRepo != nil {
		return result, errRepo
	}

	if errCache != nil {
		return result, errCache
	}

	return result, nil
}

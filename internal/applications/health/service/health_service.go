package service

import "context"

type HealthService interface {
	Health(ctx context.Context, message string) (map[string]string, error)
}

package service

import (
	"context"
	"mceasy/internal/applications/auth/dto"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.ClientCredential, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.ClientSession, error)
	ValidateClientKeyMiddleware() echo.MiddlewareFunc
}

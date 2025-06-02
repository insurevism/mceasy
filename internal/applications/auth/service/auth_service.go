package service

import (
	"context"
	"mceasy/ent"
	"mceasy/internal/applications/auth/dto"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*ent.ClientCredential, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*ent.ClientSession, error)
	ValidateClientKeyMiddleware() echo.MiddlewareFunc
}

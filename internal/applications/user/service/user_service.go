package service

import (
	"context"
	"hokusai/ent"
	"hokusai/internal/applications/user/dto"
)

type (
	UserService interface {
		Create(ctx context.Context, request *dto.UserRequest) (*ent.User, string, error)
		Update(ctx context.Context, id uint64, request *dto.UserRequest) (*ent.User, error)
		Delete(ctx context.Context, id uint64) (*ent.User, error)
		GetById(ctx context.Context, id uint64) (*ent.User, error)
		GetAll(ctx context.Context) ([]*ent.User, error)
		Login(ctx context.Context, request *dto.UserLoginRequest) (*ent.User, string, error)
	}
)

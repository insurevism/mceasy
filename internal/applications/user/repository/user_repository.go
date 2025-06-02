package repository

import (
	"context"
	"mceasy/ent"
)

type UserRepository interface {
	CreateTx(ctx context.Context, txClient *ent.Client, newUser ent.User) (*ent.User, error)
	UpdateTx(ctx context.Context, txClient *ent.Client, updateUser *ent.User) (*ent.User, error)
	ValidateExistenceByEmailAndUname(ctx context.Context, email string, username string) (*ent.User, bool, bool, error)
	Delete(ctx context.Context, id uint64) (*ent.User, error)
	SoftDelete(ctx context.Context, id uint64) (*ent.User, error)
	GetById(ctx context.Context, id uint64) (*ent.User, error)
	GetAll(ctx context.Context) ([]*ent.User, error)
	Login(ctx context.Context, email string) (*ent.User, error)
}

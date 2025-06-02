package repository

import (
	"context"
	"errors"
	"mceasy/ent"
	"mceasy/ent/user"
	"mceasy/middleware"
	"time"

	"github.com/labstack/gommon/log"
)

type UserRepositoryImpl struct {
	client *ent.Client //non transactional client
}

func NewUserRepository(client *ent.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{client: client}
}

func (r *UserRepositoryImpl) CreateTx(ctx context.Context, txClient *ent.Client, newUser ent.User) (*ent.User, error) {
	//txClient is transactional client that handled in service layer for post rollback logic
	response, err := txClient.User.Create().
		SetFullname(newUser.Fullname).
		SetUsername(newUser.Username).
		SetEmail(newUser.Email).
		SetPassword(newUser.Password).
		SetAvatar(newUser.Avatar).
		Save(ctx)

	if err != nil {
		middleware.SendDiscordNotificationError(middleware.WebhookChannelSourceERRORDATABASETRX, err)
		log.Error(err)
		return nil, err
	}

	return response, nil
}

func (r *UserRepositoryImpl) UpdateTx(ctx context.Context, txClient *ent.Client, updateUser *ent.User) (*ent.User, error) {

	affected, err := txClient.User.Update().
		Where(user.ID(updateUser.ID)).
		SetFullname(updateUser.Fullname).
		SetUsername(updateUser.Username).
		SetEmail(updateUser.Email).
		SetPassword(updateUser.Password).
		SetAvatar(updateUser.Avatar).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if affected < 1 {
		log.Errorf("ID %s no records were updated in database", updateUser.ID)
		return nil, errors.New("no records were updated in database")
	}

	updated, err := txClient.User.Get(ctx, updateUser.ID)
	if err != nil {
		middleware.SendDiscordNotificationError(middleware.WebhookChannelSourceERRORDATABASETRX, err)
		return nil, err
	}

	return updated, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id uint64) (*ent.User, error) {
	err := r.client.User.DeleteOneID(id).Exec(ctx)

	if err != nil {
		middleware.SendDiscordNotificationError(middleware.WebhookChannelSourceERRORDATABASETRX, err)
		return nil, err
	}

	return nil, nil
}

func (r *UserRepositoryImpl) SoftDelete(ctx context.Context, id uint64) (*ent.User, error) {
	deleted, err := r.client.User.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)

	if err != nil {
		middleware.SendDiscordNotificationError(middleware.WebhookChannelSourceERRORDATABASETRX, err)
		return nil, err
	}

	return deleted, nil
}

func (r *UserRepositoryImpl) GetById(ctx context.Context, id uint64) (*ent.User, error) {
	data, err := r.client.User.Query().
		Where(user.And(
			user.ID(id),
			user.DeletedAtIsNil(),
		)).
		WithUserAccounts(func(q *ent.UserAccountQuery) {
			q.WithAccountConfig()
		}).
		Only(ctx)

	if err != nil {
		middleware.SendDiscordNotificationError(middleware.WebhookChannelSourceERRORDATABASETRX, err)
		return nil, err
	}

	return data, nil
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]*ent.User, error) {
	data, err := r.client.User.Query().
		Where(user.DeletedAtIsNil()).
		All(ctx)

	if err != nil {
		middleware.SendDiscordNotificationError(middleware.WebhookChannelSourceERRORDATABASETRX, err)
		return nil, err
	}

	return data, nil
}

func (r *UserRepositoryImpl) ValidateExistenceByEmailAndUname(ctx context.Context, email string, username string) (*ent.User, bool, bool, error) {
	user, err := r.client.User.Query().
		Where(
			user.Or(
				user.EmailEQ(email),
				user.UsernameEQ(username),
			),
			user.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		return nil, false, false, err
	}

	// Check which field matched
	isEmailMatched := user.Email == email
	isUsernameMatched := user.Username == username

	return user, isEmailMatched, isUsernameMatched, nil
}

func (r *UserRepositoryImpl) Login(ctx context.Context, email string) (*ent.User, error) {
	user, err := r.client.User.Query().
		Where(
			user.EmailEQ(email),
			user.DeletedAtIsNil(),
		).
		WithUserAccounts(func(q *ent.UserAccountQuery) {
			q.WithAccountConfig()
		}).
		Only(ctx)

	if err != nil {
		middleware.SendDiscordNotificationError(middleware.WebhookChannelSourceERRORDATABASETRX, err)
		return nil, err
	}

	return user, nil
}

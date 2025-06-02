package service

import (
	"context"
	"fmt"
	"mceasy/ent"
	"mceasy/exceptions"
	"mceasy/internal/applications/user/dto"
	userRepository "mceasy/internal/applications/user/repository"
	caching "mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"
	"mceasy/internal/vars"
	"mceasy/middleware"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository userRepository.UserRepository
	transaction    transaction.Trx
	cache          caching.Cache
}

func NewUserService(
	userRepository userRepository.UserRepository,
	transaction transaction.Trx,
	cache caching.Cache,
) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
		transaction:    transaction,
		cache:          cache,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, request *dto.UserRequest) (*ent.User, string, error) {
	// Validate email and username existence:
	_, emailExists, usernameExists, err := s.userRepository.ValidateExistenceByEmailAndUname(ctx, request.Email, request.Username)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	if emailExists {
		return nil, "", fmt.Errorf("%s is already taken!", request.Email)
	}
	if usernameExists {
		return nil, "", fmt.Errorf("%s is already taken!", request.Username)
	}

	var userNew = &ent.User{}
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {
		// Create user object:
		userRequest := ent.User{
			Fullname: request.Fullname,
			Username: request.Username,
			Email:    request.Email,
			Password: request.Password,
			Avatar:   request.Avatar,
		}

		// Save user:
		userResult, err := s.userRepository.CreateTx(ctx, tx.Client(), userRequest)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
		}

		userNew = userResult

		return nil
	}); err != nil {
		// Add rollback logic here
		log.Errorf("do rollback from transactional database operation. ERROR: %v", err)
		return nil, "", err
	}

	// Generate JWT token for the newly created user
	token, err := middleware.JWTTokenGenerator(userNew)
	if err != nil {
		log.Errorf("Failed to generate JWT token: %v", err)
		return nil, "", err
	}

	// Return user object along with the token
	return userNew, token, nil
}
func (s *UserServiceImpl) Update(ctx context.Context, id uint64, request *dto.UserRequest) (*ent.User, error) {
	var userUpdated = &ent.User{}
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {

		userExisting, err := s.userRepository.GetById(ctx, id)
		if userExisting == nil || err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
		}

		userExisting.Fullname = request.Fullname
		userExisting.Username = request.Username
		userExisting.Email = request.Email
		userExisting.Password = request.Password
		userExisting.Avatar = request.Avatar

		//update user:
		userResult, err := s.userRepository.UpdateTx(ctx, tx.Client(), userExisting)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
		}

		//set value to userUpdated for return value:
		userUpdated = userResult

		//create cache, don't throw exception if failed:
		_, _ = s.cache.Create(ctx, CacheKeyUserWithId(id), userUpdated, vars.GetTtlShortPeriod())

		return nil

	}); err != nil {
		//add rollback logic here
		log.Error("do rollback from transactional database operation")
		return nil, err
	}

	return userUpdated, nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id uint64) (*ent.User, error) {
	data, err := s.userRepository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataDeleteFailed, err)
	}

	_, err = s.cache.Delete(ctx, CacheKeyUserWithId(id))
	if err != nil {
		return data, nil
	}

	return data, nil
}

func (s *UserServiceImpl) GetById(ctx context.Context, id uint64) (*ent.User, error) {
	result, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	_, err = s.cache.Create(ctx, CacheKeyUserWithId(id), result, vars.GetTtlShortPeriod())
	if err != nil {
		return result, nil
	}

	return result, nil
}

func (s *UserServiceImpl) GetAll(ctx context.Context) ([]*ent.User, error) {
	userCache, err := s.cache.Get(ctx, CacheKeyUsers(), &[]*ent.User{})
	if userCache != nil {
		userResult := append([]*ent.User(nil), *userCache.(*[]*ent.User)...)
		return userResult, err
	}

	result, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	_, err = s.cache.Create(ctx, CacheKeyUsers(), &result, vars.GetTtlShortPeriod())
	if err != nil {
		return result, nil
	}

	return result, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, request *dto.UserLoginRequest) (*ent.User, string, error) {
	user, err := s.userRepository.Login(ctx, request.Email)

	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, "", fmt.Errorf("credential does not match")
	}

	// Generate JWT token for the user
	token, err := middleware.JWTTokenGenerator(user)
	if err != nil {
		log.Errorf("Failed to generate JWT token: %v", err)
		return nil, "", err
	}

	return user, token, nil

}

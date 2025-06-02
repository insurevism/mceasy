package service

import (
	"context"
	"errors"
	"mceasy/exceptions"
	"mceasy/internal/applications/auth/constant"
	"mceasy/internal/applications/auth/dto"
	autherr "mceasy/internal/applications/auth/errors"
	passwordhasher "mceasy/internal/applications/auth/utils/password_hasher"
	userDto "mceasy/internal/applications/user/dto"
	userService "mceasy/internal/applications/user/service"
	"mceasy/internal/helper/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

var _ AuthService = (*AuthServiceImpl)(nil)

type AuthServiceImpl struct {
	passwordHasher passwordhasher.PasswordHasher
	userService    userService.UserService
}

func NewAuthService(
	passwordHasher passwordhasher.PasswordHasher,
	userService userService.UserService,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		passwordHasher: passwordHasher,
		userService:    userService,
	}
}

var ErrUsernameIsRegisterd = exceptions.NewBusinessLogicError(
	exceptions.DataCreateFailed,
	errors.New("username is already registered"),
)

var ErrUsernameIsNotFound = exceptions.NewBusinessLogicError(
	exceptions.DataNotFound,
	errors.New("username is not registered"),
)

var ErrCredentialDoesNotMatch = exceptions.NewBusinessLogicError(
	exceptions.InvalidValue,
	errors.New("username and password does not match"),
)

var ErrClientKeyIsNotProvided = autherr.NewClientAuthError("client key is not provided")
var ErrClientUnauthorized = autherr.NewClientAuthError("unauthorized client")

// Register - creates a new user in the database
func (s *AuthServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.ClientCredential, error) {
	// Hash the password
	hashedPassword, err := s.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	// Create user request
	userReq := &userDto.UserRequest{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Avatar:   req.Avatar,
	}

	// Create user using user service
	user, _, err := s.userService.Create(ctx, userReq)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	// Return client credential with the created user info
	return &dto.ClientCredential{
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.Fullname,
	}, nil
}

// Login - validates user credentials
func (s *AuthServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.ClientSession, error) {
	// Use user service to authenticate
	loginReq := &userDto.UserLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	user, token, err := s.userService.Login(ctx, loginReq)
	if err != nil {
		return nil, ErrCredentialDoesNotMatch
	}

	// Return client session with the token
	return &dto.ClientSession{
		Username:  user.Username,
		Email:     user.Email,
		Fullname:  user.Fullname,
		ClientKey: token,
	}, nil
}

// ValidateClientKeyMiddleware - simplified middleware (for now just checks if key exists)
func (s *AuthServiceImpl) ValidateClientKeyMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			clientKey := req.Header.Get(constant.ClientKeyHeader)
			if clientKey == "" {
				return response.Error(c, http.StatusUnauthorized, ErrClientKeyIsNotProvided, nil)
			}

			// For now, accept any non-empty client key
			// In a real implementation, you would validate against a session store
			if clientKey == "" {
				return response.Error(c, http.StatusUnauthorized, ErrClientUnauthorized, nil)
			}

			return next(c)
		}
	}
}

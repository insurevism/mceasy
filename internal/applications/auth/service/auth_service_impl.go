package service

import (
	"context"
	"errors"
	"mceasy/ent"
	"mceasy/exceptions"
	"mceasy/internal/applications/auth/constant"
	"mceasy/internal/applications/auth/dto"
	autherr "mceasy/internal/applications/auth/errors"
	passwordhasher "mceasy/internal/applications/auth/utils/password_hasher"
	clientcredentialdto "mceasy/internal/applications/auth_client_credential/dto"
	clientcredentialdb "mceasy/internal/applications/auth_client_credential/repository/db"
	clientsessionsvc "mceasy/internal/applications/auth_client_session/service"
	"mceasy/internal/helper/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

var _ AuthService = (*AuthServiceImpl)(nil)

type AuthServiceImpl struct {
	clientCredentialRepo clientcredentialdb.ClientCredentialRepository
	clientSessionService clientsessionsvc.ClientSessionService
	passwordHasher       passwordhasher.PasswordHasher
}

func NewAuthService(
	clientCredentialRepo clientcredentialdb.ClientCredentialRepository,
	clientSessionService clientsessionsvc.ClientSessionService,
	passwordHasher passwordhasher.PasswordHasher,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		clientCredentialRepo: clientCredentialRepo,
		clientSessionService: clientSessionService,
		passwordHasher:       passwordHasher,
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

// Register registers a new client with the provided username and password.
// It checks if the username is already registered, hashes the password, and creates a new client credential.
// Returns the created client credential on success, or an error if the registration fails.
func (s *AuthServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*ent.ClientCredential, error) {
	// check if the username is already reigstered or not
	_, err := s.clientCredentialRepo.GetOne(ctx, &clientcredentialdto.ClientCredentialOption{
		Condition: &clientcredentialdto.ClientCredentialCondition{
			Username:       req.Username,
			DeletedAtIsNil: true,
		},
	})

	isRegistered, err := s.isClientRegistered(err)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	if isRegistered {
		return nil, ErrUsernameIsRegisterd
	}

	// hash the password
	hashedPassword, err := s.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	newUser := &ent.ClientCredential{
		Username: req.Username,
		Password: hashedPassword,
	}

	created, err := s.clientCredentialRepo.Create(ctx, newUser)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	return created, nil
}

// Login authenticates a user based on the provided login request.
// It checks if the username is registered, compares the password, and either returns a client session or an error.
func (s *AuthServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*ent.ClientSession, error) {
	// check if the username is already reigstered or not
	clientCredential, err := s.clientCredentialRepo.GetOne(ctx, &clientcredentialdto.ClientCredentialOption{
		Condition: &clientcredentialdto.ClientCredentialCondition{
			Username:       req.Username,
			DeletedAtIsNil: true,
		},
	})

	isRegistered, err := s.isClientRegistered(err)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	if !isRegistered {
		return nil, ErrUsernameIsNotFound
	}

	// compare the password
	isCredentialMatches, err := s.passwordHasher.Compare(req.Password, clientCredential.Password)
	if err != nil {
		return nil, err
	}

	if !isCredentialMatches {
		return nil, ErrCredentialDoesNotMatch
	}

	// get existing or create new session
	return s.clientSessionService.GetOrCreate(ctx, clientCredential)
}

// ValidateClientKeyMiddleware returns an Echo middleware function that checks the client key in the request header.
// If the client key is not provided or invalid, it returns an unauthorized error response.
// It validates the client key using the client session service.
func (s *AuthServiceImpl) ValidateClientKeyMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			clientKey := req.Header.Get(constant.ClientKeyHeader)
			if clientKey == "" {
				return response.Error(c, http.StatusUnauthorized, ErrClientKeyIsNotProvided, nil)
			}

			isValidSession, err := s.clientSessionService.IsValidClientKey(req.Context(), clientKey)
			if err != nil {
				return response.Error(c, http.StatusUnauthorized, err, nil)
			}

			if !isValidSession {
				return response.Error(c, http.StatusUnauthorized, ErrClientUnauthorized, nil)
			}

			return next(c)
		}
	}
}

func (*AuthServiceImpl) isClientRegistered(err error) (bool, error) {
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

package dto

type (
	UserRequest struct {
		Fullname string `json:"fullname" validate:"required"`
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
		Avatar   string `json:"avatar" validate:"required"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

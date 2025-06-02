package dto

import "time"

type (
	UserResponse struct {
		ID        int64     `json:"id" validate:"required"`
		Fullname  string    `json:"fullname" validate:"required"`
		Username  string    `json:"username" validate:"required"`
		Email     string    `json:"email" validate:"required"`
		Avatar    string    `json:"avatar" validate:"required"`
		Phone     string    `json:"phone" validate:"omitempty"`
		CreatedAt time.Time `json:"created_at" validate:"required"`
		Token     string    `json:"token" validate:"required"`
	}

	UserUpdateResponse struct {
		ID         int64     `json:"id" validate:"required"`
		Fullname   string    `json:"fullname" validate:"required"`
		Username   string    `json:"username" validate:"required"`
		Email      string    `json:"email" validate:"required"`
		Avatar     string    `json:"avatar" validate:"required"`
		Phone      string    `json:"phone" validate:"omitempty"`
		ModifiedAt time.Time `json:"modified_at" validate:"required"`
	}

	UserDeleteResponse struct {
		IsDeleted bool      `json:"is_deleted" validate:"required"`
		DeletedAt time.Time `json:"deleted_at" validate:"required"`
	}

	UserLoginResponse struct {
		ID       int64  `json:"id" validate:"required"`
		Fullname string `json:"fullname" validate:"required"`
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Avatar   string `json:"avatar" validate:"required"`
		Phone    string `json:"phone" validate:"omitempty"`
		Token    string `json:"token" validate:"required"`
	}
)

package dto

type RegisterRequest struct {
	Username string `form:"username" validate:"required"`
	// Password with minimum length of 8 chars.
	Password string `form:"password" validate:"required,min=8"`
}

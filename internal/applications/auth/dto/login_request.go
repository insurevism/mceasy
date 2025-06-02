package dto

type LoginRequest struct {
	Email string `form:"email" validate:"required"`
	// Password with minimum length of 8 chars.
	Password string `form:"password" validate:"required,min=8"`
}

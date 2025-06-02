package dto

type RegisterRequest struct {
	Fullname string `form:"fullname" json:"fullname" validate:"required"`
	Username string `form:"username" json:"username" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	// Password with minimum length of 8 chars.
	Password string `form:"password" json:"password" validate:"required,min=8"`
	Avatar   string `form:"avatar" json:"avatar"`
}

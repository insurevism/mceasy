package dto

type RegisterSuccessResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

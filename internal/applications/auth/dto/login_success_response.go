package dto

type LoginSuccessResponse struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	ClientKey string `json:"clientKey"`
}

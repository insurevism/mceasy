package dto

type LoginSuccessResponse struct {
	ClientID  uint64 `json:"clientId"`
	ClientKey string `json:"clientKey"`
}

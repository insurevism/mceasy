package http

import (
	"hokusai/middleware"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func ClientHttpLazada() *resty.Client {
	clientApi := resty.New()

	timeout := viper.GetInt("http.configs.lazada.timeout")
	retryCount := viper.GetInt("http.configs.lazada.retryCount")
	retryWaitTime := viper.GetInt("http.configs.lazada.waitTime")
	retryMaxWaitTime := viper.GetInt("http.configs.lazada.maxWaitTime")

	clientApi.SetTimeout(time.Duration(timeout) * time.Millisecond).
		SetRetryCount(retryCount).                                               // Maximum number of retries
		SetRetryWaitTime(time.Duration(retryWaitTime) * time.Millisecond).       // Time to wait between retries
		SetRetryMaxWaitTime(time.Duration(retryMaxWaitTime) * time.Millisecond). // Maximum time to wait between retries
		AddRetryCondition(func(response *resty.Response, err error) bool {
			switch response.StatusCode() {
			case http.StatusOK:
				return false
			case http.StatusForbidden: // unauthorized, invalid token
				return false
			default:
				return true
			}
		}).
		OnRequestLog(middleware.LogRequest("[lazada]")).
		OnResponseLog(middleware.LogResponse("[lazada]"))

	return clientApi
}

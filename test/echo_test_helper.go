package test

import (
	"hokusai/configs/validator"
	"testing"

	"github.com/labstack/echo/v4"
)

func InitEchoTest(*testing.T) *echo.Echo {
	e := echo.New()
	validator.SetupValidator(e)
	validator.SetupGlobalHttpUnhandleErrors(e)

	return e
}

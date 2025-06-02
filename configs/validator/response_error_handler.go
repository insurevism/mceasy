package validator

import (
	"fmt"
	"mceasy/internal/helper/response"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SetupGlobalHttpUnhandleErrors(e *echo.Echo) {
	e.HTTPErrorHandler = GlobalUnHandleErrors()
	log.Infof("initialized GlobalUnHandleErrors : success")
}

func GlobalUnHandleErrors() func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		if errs, ok := err.(validator.ValidationErrors); ok {
			errorMap := make(map[string]string)
			for _, e := range errs {

				constructErrMsg := fmt.Sprintf("Struct Field=%s is %s actual value is=%s", e.Field(), e.ActualTag(), e.Param())
				// e is of type validator.FieldError
				// You can access fields like:
				// e.Field()
				// e.Tag()
				// e.Value()
				// e.Param()

				//_, msg := MapperErrorCode(e)
				errorMap[e.Field()] = constructErrMsg
			}

			_ = response.Base(ctx, http.StatusBadRequest, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errorMap, nil)
			return
		}

		errBusinessCode, msg := MapperErrorCode(err)
		_ = response.Base(ctx, http.StatusInternalServerError, errBusinessCode, msg, nil, nil)
		return
	}
}

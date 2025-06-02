package validator

import (
	"errors"
	"fmt"
	"mceasy/exceptions"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// MapperErrorCode all mapper goes here
func MapperErrorCode(err error) (errBusinessCode int, errMessage string) {

	if errors.Is(err, exceptions.TargetBusinessLogicError) {
		errorCode := err.(*exceptions.BusinessLogicError).ErrorCode
		errorLogic := exceptions.BusinessLogicReason(errorCode)

		return errorLogic.ErrCode, errorLogic.Message
	}

	errorMessage := err.Error()
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", err.Field())
			case "email":
				errorMessage = fmt.Sprintf("%s is not valid email", err.Field())
			case "gte":
				errorMessage = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
			case "lte":
				errorMessage = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
			case "password":
				errorMessage = fmt.Sprintf("%s %s", err.Field(), err.Value())
			}
			break
		}

		return http.StatusBadRequest, errorMessage
	}

	return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)

}

package errors

var _ error = (*ClientAuthError)(nil)

type ClientAuthError struct {
	Message string
}

func NewClientAuthError(message string) *ClientAuthError {
	return &ClientAuthError{Message: message}
}

func (e *ClientAuthError) Error() string {
	return e.Message
}

package errors

import "fmt"

var _ error = (*ProducerError)(nil)

type ProducerError struct {
	msg string
}

func NewProducerError(reason string, args ...any) *ProducerError {
	return &ProducerError{
		msg: fmt.Sprintf(reason, args...),
	}
}

func (pe *ProducerError) Error() string {
	return pe.msg
}

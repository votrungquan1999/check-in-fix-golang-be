package utils

import "net/http"

type CommonErrors uint

const (
	ErrorInternal CommonErrors = iota
	ErrorBadRequest
	ErrorEntityNotFound
)

var mapErrorHTTPCode = map[CommonErrors]int{
	ErrorInternal:       http.StatusInternalServerError,
	ErrorBadRequest:     http.StatusBadRequest,
	ErrorEntityNotFound: http.StatusNotFound,
}

func (c CommonErrors) New(message string) error {
	return &CustomError{
		Code:    mapErrorHTTPCode[c],
		Message: message,
	}
}

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return c.Message
}
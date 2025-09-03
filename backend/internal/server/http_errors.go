package server

import "time"

const (
	ErrInvalidPayload = "ERR_INVALID_PAYLOAD"
	ErrInternal       = "ERR_INTERNAL_SERVER_ERROR"
)

type httpError struct {
	Timestamp  time.Time `json:"timestamp"`
	Error      string    `json:"error"`
	Message    any       `json:"message"`
	StatusCode int       `json:"status_code"`
}

func NewHttpError(statusCode int, message any, err string) httpError {
	return httpError{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
		Timestamp:  time.Now(),
	}
}

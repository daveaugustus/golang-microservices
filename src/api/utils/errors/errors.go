package errors

import "net/http"

type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
	Aerror   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.Astatus
}

func (e *apiError) Message() string {
	return e.Amessage
}

func (e *apiError) Error() string {
	return e.Aerror
}

func NewApiError(statusCode int, message string) APIError {
	return &apiError{
		Astatus:  statusCode,
		Amessage: message,
	}
}

func NewInternalServerError(message string) APIError {
	return &apiError{
		Astatus:  http.StatusInternalServerError,
		Amessage: message,
	}
}

func NewNotFoundError(message string) APIError {
	return &apiError{
		Astatus:  http.StatusNotFound,
		Amessage: message,
	}
}

func NewBadRequestError(message string) APIError {
	return &apiError{
		Astatus:  http.StatusBadRequest,
		Amessage: message,
	}
}

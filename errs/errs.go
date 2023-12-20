package errs

import "net/http"

type AppError struct {
	Code      int
	Messesage string
}

func (e AppError) Error() string {
	return e.Messesage
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:      http.StatusNotFound,
		Messesage: message,
	}
}

func NewUnexpectedError(message string) error {
	return AppError{
		Code:      http.StatusInternalServerError,
		Messesage: message,
	}
}

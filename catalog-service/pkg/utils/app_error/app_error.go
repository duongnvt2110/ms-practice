package apperror

import "net/http"

type AppError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	err        error
}

func (e AppError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return e.Message
}

func (e AppError) Wrap(err error) AppError {
	e.err = err
	return e
}

func Err(status int, msg string) AppError {
	return AppError{
		StatusCode: status,
		Message:    msg,
	}
}

var (
	ErrBadRequest     = Err(http.StatusBadRequest, "bad_request")
	ErrInternalServer = Err(http.StatusInternalServerError, "internal_error")
	ErrNotFound       = Err(http.StatusNotFound, "not_found")
)

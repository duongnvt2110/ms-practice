package errors

import (
	"errors"
	"fmt"
)

type AppError interface {
	error
	Catch(err error) AppError
	Origin(err error) error
	Wrap(cause error) AppError
	Wrapf(cause error, format string, a ...interface{}) AppError
	GetHttpCode() int
	GetErrCode() string
}

type appError struct {
	err       error
	httpCode  int
	errorCode string
	cause     error
	stack     *Stack
}

func NewAppError(errCode string, httpCode int, errMsg string) AppError {
	return &appError{
		err:       errors.New(errMsg),
		httpCode:  httpCode,
		errorCode: errCode,
	}
}

func (e *appError) Origin(err error) error {
	for {
		e := errors.Unwrap(err)
		if e == nil {
			return err
		}
		err = e
	}
}

func (e *appError) Error() string {
	return e.err.Error()
}

func (e *appError) Cause() error {
	return e.cause
}

func (e *appError) Catch(err error) AppError {
	if err == nil {
		return nil
	}
	if aerr, ok := err.(AppError); ok {
		return aerr
	}
	return ErrInternalServer.Wrap(err)
}

func (e *appError) Wrap(cause error) AppError {
	return e.withStack().setCause(cause)
}

func (e *appError) Wrapf(cause error, format string, a ...interface{}) AppError {
	return e.withStack().setCause(cause).setMessage(format, a...)
}

func (e *appError) GetHttpCode() int {
	return e.httpCode
}

func (e *appError) GetErrCode() string {
	return e.errorCode
}

// Private function
func (e *appError) setCause(err error) *appError {
	e.cause = err
	return e
}

func (e *appError) setMessage(format string, a ...interface{}) *appError {
	a = append([]interface{}{e.err}, a...)
	e.err = fmt.Errorf("%w: "+format, a...)
	return e
}

func (e *appError) withStack() *appError {
	err := *e
	err.stack = StackTrace(2)
	return &err
}

package errorsx

import (
	"errors"
	"fmt"
)

type AppError interface {
	error
	Catch(err error) AppError
	Origin(err error) error
	PublicMessage() string
	Wrap(cause error) AppError
	Unwrap() error
	Wrapf(cause error, format string, a ...interface{}) AppError
	GetHttpCode() int
	GetErrCode() string
	GetStacks() string
}

type appError struct {
	msg       string
	httpCode  int
	errorCode string
	cause     error
	detail    string
	stack     *Stack
}

type stackError struct {
	err   error
	stack *Stack
}

func NewAppError(errCode string, httpCode int, errMsg string) AppError {
	return &appError{
		msg:       errMsg,
		httpCode:  httpCode,
		errorCode: errCode,
		stack:     StackTrace(2),
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
	return e.msg
}

func (e *appError) Cause() string {
	if e.cause == nil {
		return ""
	}
	return e.cause.Error()
}

func (e *appError) Catch(err error) AppError {
	if err == nil {
		return nil
	}
	var aerr AppError
	if errors.As(err, &aerr) {
		return aerr
	}
	return ErrInternalServerError.Wrap(err)
}

func (e *appError) Unwrap() error {
	return e.cause
}

func (e *appError) Wrap(cause error) AppError {
	return e.withStack().setCause(cause)
}

func (e *appError) Wrapf(cause error, format string, a ...interface{}) AppError {
	return e.withStack().setCause(cause).setDetail(format, a...)
}

func (e *appError) GetHttpCode() int {
	return e.httpCode
}

func (e *appError) GetErrCode() string {
	return e.errorCode
}

func (e *appError) GetStacks() string {
	if e.stack == nil {
		return ""
	}
	return fmt.Sprintf("%+v", e.stack)
}

func (e *appError) PublicMessage() string {
	return e.msg
}

func (e *appError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprint(s, e.msg)
			if e.detail != "" {
				fmt.Fprintf(s, ": %s", e.detail)
			}
			if e.stack != nil {
				fmt.Fprintf(s, "\n%+v", e.stack)
			}
			if e.cause != nil {
				fmt.Fprintf(s, "\nCaused by: %+v", e.cause)
			}
			return
		}
		fallthrough
	case 's':
		fmt.Fprint(s, e.msg)
	}
}

// Private function
func (e *appError) setCause(err error) *appError {
	e.cause = err
	return e
}

func (e *appError) setDetail(format string, a ...interface{}) *appError {
	e.detail = fmt.Sprintf(format, a...)
	return e
}

func (e *appError) withStack() *appError {
	err := *e
	err.stack = StackTrace(2)
	return &err
}

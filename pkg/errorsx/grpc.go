package errorsx

import (
	"errors"
	"net/http"
	"strconv"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToStatus(err error) error {
	app, ok := err.(*appError)
	if !ok {
		return status.Error(codes.Internal, "internal error")
	}

	httpCode := mapHTTPToGRPC(app.httpCode)
	st := status.New(httpCode, app.msg)
	ei := &errdetails.ErrorInfo{
		Reason: app.errorCode,
		Metadata: map[string]string{
			"msg":       app.msg,
			"http_code": strconv.Itoa(app.httpCode),
			"err_code":  app.errorCode,
			"detail":    app.detail,
			"cause":     app.Cause(),
		},
	}
	st2, err := st.WithDetails(ei)
	if err != nil {
		return st.Err()
	}
	return st2.Err()
}

func FromStatus(err error) (error, bool) {
	if err == nil {
		return nil, false
	}
	st, ok := status.FromError(err)
	if !ok {
		return ErrInternalServerError, false
	}

	var ei *errdetails.ErrorInfo
	for _, d := range st.Details() {
		if x, ok := d.(*errdetails.ErrorInfo); ok {
			ei = x
			break
		}
	}

	if ei != nil {
		errCode := ei.Metadata["err_code"]

		detail := ei.Metadata["detail"]
		cause := ei.Metadata["cause"]
		msg := ei.Metadata["msg"]

		httpCode := http.StatusInternalServerError
		if s := ei.Metadata["http_code"]; s != "" {
			if v, err := strconv.Atoi(s); err == nil && v >= 100 && v <= 599 {
				httpCode = v
			}
		}

		base := NewAppError(errCode, httpCode, msg)

		if cause != "" {
			c := errors.New(cause)
			if detail != "" {
				return base.Wrapf(c, "%s", detail), true
			}
			return base.Wrap(c), true
		}

		return base, true
	}

	return ErrUnknownError, true
}

func mapHTTPToGRPC(httpStatus int) codes.Code {
	switch httpStatus {
	case 400:
		return codes.InvalidArgument
	case 401:
		return codes.Unauthenticated
	case 403:
		return codes.PermissionDenied
	case 404:
		return codes.NotFound
	case 409:
		return codes.AlreadyExists
	case 429:
		return codes.ResourceExhausted
	default:
		if httpStatus >= 500 {
			return codes.Internal
		}
		return codes.Unknown
	}
}

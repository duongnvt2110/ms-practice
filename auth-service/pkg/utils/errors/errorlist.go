package errors

var (
	ErrInternalServer = NewAppError("500000", 500, "internal server error")
	ErrUnauthorized   = NewAppError("403000", 403, "unauthorized")
)

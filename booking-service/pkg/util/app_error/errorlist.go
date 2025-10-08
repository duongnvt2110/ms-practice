package apperror

var (
	ErrInternalServer = NewAppError("500000", 500, "internal server error")
	ErrBadRequest     = NewAppError("400000", 400, "bad request")
	ErrNotFound       = NewAppError("404000", 404, "not found")
)

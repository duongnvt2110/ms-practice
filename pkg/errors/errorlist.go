package errors

var (
	ErrInternalServerError = NewAppError("500000", 500, "internal server error")
	ErrUnauthorized        = NewAppError("403000", 403, "unauthorized")
	ErrUnauthenticated     = NewAppError("401000", 401, "unauthenticated")
	ErrBadRequest          = NewAppError("400000", 400, "bad request")
	ErrNotFound            = NewAppError("404000", 404, "not found")
	ErrConflict            = NewAppError("409000", 409, "conflict")
)

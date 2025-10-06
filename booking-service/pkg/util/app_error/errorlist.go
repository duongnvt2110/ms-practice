package apperror

var (
	ErrInternalServer = NewAppError("500000", 500, "internal server error")
)

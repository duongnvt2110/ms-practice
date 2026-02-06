package apperr

import "ms-practice/pkg/errorsx"

var (
	ErrUserAlreadyExists    = errorsx.NewAppError("409001", 409, "user already exists")
	ErrInvalidCredentials   = errorsx.NewAppError("401001", 401, "invalid email or password")
	ErrTokenRequired        = errorsx.NewAppError("400001", 400, "token required")
	ErrRefreshTokenRequired = errorsx.NewAppError("400002", 400, "refresh token is required")
	ErrInvalidRefreshToken  = errorsx.NewAppError("401002", 401, "invalid refresh token")
	ErrRefreshTokenExpired  = errorsx.NewAppError("401003", 401, "refresh token expired")
	ErrInvalidToken         = errorsx.NewAppError("401004", 401, "invalid token")
	ErrTokenExpired         = errorsx.NewAppError("401005", 401, "token expired")
	ErrUserNotFound         = errorsx.NewAppError("404001", 404, "user not found")
)

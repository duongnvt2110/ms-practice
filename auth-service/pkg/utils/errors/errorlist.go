package errors

import sharederrors "ms-practice/pkg/errors"

var (
	ErrUserAlreadyExists    = sharederrors.NewAppError("409001", 409, "user already exists")
	ErrInvalidCredentials   = sharederrors.NewAppError("401001", 401, "invalid email or password")
	ErrTokenRequired        = sharederrors.NewAppError("400001", 400, "token required")
	ErrRefreshTokenRequired = sharederrors.NewAppError("400002", 400, "refresh token is required")
	ErrInvalidRefreshToken  = sharederrors.NewAppError("401002", 401, "invalid refresh token")
	ErrRefreshTokenExpired  = sharederrors.NewAppError("401003", 401, "refresh token expired")
	ErrInvalidToken         = sharederrors.NewAppError("401004", 401, "invalid token")
	ErrTokenExpired         = sharederrors.NewAppError("401005", 401, "token expired")
	ErrUserNotFound         = sharederrors.NewAppError("404001", 404, "user not found")
)

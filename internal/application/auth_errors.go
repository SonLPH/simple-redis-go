package application

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrDBInternal     = errors.New("internal server")
	ErrAuth           = errors.New("invalid pass or email")
	ErrSessionInvalid = errors.New("invalid session, please login again")
	ErrPing           = errors.New("ping failed")
	ErrRateLimit      = errors.New("rate limit exceeded")
)

package errors

import "errors"

var (
	ErrOtpTokenExpired    = errors.New("otp confirm token expired")
	ErrOtpUnavailable     = errors.New("unavailable otp confirm token")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInternalServer     = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrSysUserAlreadyExists = errors.New("sysuser already exists")
	ErrRoleNotFound         = errors.New("role not found")
	ErrUserCreationFailed   = errors.New("failed to create user")
)

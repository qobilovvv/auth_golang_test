package errors

import "errors"

var (
	ErrOtpTokenExpired       = errors.New("otp confirm token expired")
	ErrOtpUnavailable        = errors.New("unavailable otp confirm token")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrInvalidEmail          = errors.New("invalid email format")
	ErrInternalServer        = errors.New("internal server error")
	ErrUserCreationFailed    = errors.New("failed to create user")
)

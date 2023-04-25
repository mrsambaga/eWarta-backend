package httperror

import (
	"errors"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrinvalidEmailFormat     = errors.New("invalid email format")
	ErrInvalidPasswordLength  = errors.New("password length must be 8 or more")
	ErrFailedCreateToken      = errors.New("failed to create token")
	ErrGenerateHash           = errors.New("failed to generate hash")
	ErrUserNotFound           = errors.New("user not found")
	ErrFailedGetUserByEmail   = errors.New("failed to get user by email")
	ErrInvalidEmailPassword   = errors.New("invalid email or password")
)

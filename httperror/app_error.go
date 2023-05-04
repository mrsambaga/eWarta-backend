package httperror

import (
	"errors"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrInvalidEmailFormat     = errors.New("invalid email format")
	ErrInvalidPasswordLength  = errors.New("password must be at least 8 characters")
	ErrFailedCreateToken      = errors.New("failed to create token")
	ErrGenerateHash           = errors.New("failed to generate hash")
	ErrUserNotFound           = errors.New("user not found")
	ErrFailedGetUserByEmail   = errors.New("failed to get user by email")
	ErrInvalidEmailPassword   = errors.New("invalid email or password")
	ErrInvalidReferral        = errors.New("invalid referral")
	ErrCreatePost             = errors.New("failed to create post")
	ErrFindNews               = errors.New("failed to find news")
	ErrNewsNotFound           = errors.New("news not found")
	ErrFailedConvertId        = errors.New("failed to convert id to integer")
	ErrInvalidRole            = errors.New("invalid role")
	ErrDeleteNews             = errors.New("failed to delete news")
)

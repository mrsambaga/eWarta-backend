package handler

import (
	"stage01-project-backend/usecase"
)

type Handler struct {
	userUsecase usecase.UsersUsecase
}

type Config struct {
	UserUsecase usecase.UsersUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase: cfg.UserUsecase,
	}
}

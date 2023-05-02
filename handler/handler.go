package handler

import (
	"stage01-project-backend/usecase"
)

type Handler struct {
	userUsecase usecase.UsersUsecase
	postUsecase usecase.PostsUsecase
}

type Config struct {
	UserUsecase usecase.UsersUsecase
	PostUsecase usecase.PostsUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase: cfg.UserUsecase,
		postUsecase: cfg.PostUsecase,
	}
}

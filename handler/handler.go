package handler

import (
	"stage01-project-backend/usecase"
)

type Handler struct {
	userUsecase        usecase.UsersUsecase
	postUsecase        usecase.PostsUsecase
	transactionUsecase usecase.TransactionUsecase
}

type Config struct {
	UserUsecase        usecase.UsersUsecase
	PostUsecase        usecase.PostsUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase:        cfg.UserUsecase,
		postUsecase:        cfg.PostUsecase,
		transactionUsecase: cfg.TransactionUsecase,
	}
}

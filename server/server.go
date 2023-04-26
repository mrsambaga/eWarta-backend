package server

import (
	"log"
	"stage01-project-backend/db"
	"stage01-project-backend/repository"
	"stage01-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.Get(),
	})
	userUsecase := usecase.NewUsersUsecase(&usecase.UsersUsecaseConfig{
		UsersRepository: userRepo,
	})

	return NewRouter(&RouterConfig{
		UserUsecase: userUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
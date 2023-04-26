package server

import (
	"log"
	"net/http"
	"stage01-project-backend/handler"
	"stage01-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase usecase.UsersUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.New()

	h := handler.New(&handler.Config{
		UserUsecase: cfg.UserUsecase,
	})

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	log.Fatal(http.ListenAndServe(":8000", router))
	return router
}

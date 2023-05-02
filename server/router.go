package server

import (
	"log"
	"net/http"
	"stage01-project-backend/handler"
	"stage01-project-backend/middleware"
	"stage01-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase usecase.UsersUsecase
	PostUsecase usecase.PostsUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.New(&handler.Config{
		UserUsecase: cfg.UserUsecase,
		PostUsecase: cfg.PostUsecase,
	})

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
	router.GET("/news", middleware.AuthorizeJWT, h.FindAllNews)

	log.Fatal(http.ListenAndServe(":8000", router))
	return router
}

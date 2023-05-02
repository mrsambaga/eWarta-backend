package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FindAllNews(c *gin.Context) {
	posts, err := h.postUsecase.FindAllNews()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success Find Posts",
		"data":    posts,
	})
}

func (h *Handler) FindAllNewsHighlight(c *gin.Context) {
	postsHighlight, err := h.postUsecase.FindAllNewsHighlight()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success Find Posts Highlight",
		"data":    postsHighlight,
	})
}

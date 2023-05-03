package handler

import (
	"net/http"
	"stage01-project-backend/constant"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FindAllNews(c *gin.Context) {
	params := &constant.Params{
		Category: c.Query("category"),
		NewsType: c.Query("type"),
		Date:     c.Query("date"),
	}
	posts, err := h.postUsecase.FindAllNews(params)
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
	params := &constant.Params{
		Title:    c.Query("title"),
		Category: c.Query("category"),
		NewsType: c.Query("type"),
		Date:     c.Query("date"),
	}

	postsHighlight, err := h.postUsecase.FindAllNewsHighlight(params)
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

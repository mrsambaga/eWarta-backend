package handler

import (
	"errors"
	"net/http"
	"stage01-project-backend/constant"
	"stage01-project-backend/dto"
	"stage01-project-backend/httperror"
	"stage01-project-backend/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) FindAllNews(c *gin.Context) {
	params := &constant.Params{
		Title:    c.Query("title"),
		Category: c.Query("category"),
		NewsType: c.Query("type"),
		Date:     c.Query("date"),
	}
	posts, err := h.postUsecase.FindAllNews(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
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

func (h *Handler) FindNewsDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": httperror.ErrFailedConvertId,
			"data":    nil,
		})
		return
	}

	postsHighlight, err := h.postUsecase.FindNewsDetail(uint64(idInt))
	if err != nil {
		if errors.Is(err, httperror.ErrNewsNotFound) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    "SUCCESS_ACCESSED",
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
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

func (h *Handler) SoftDeletePost(c *gin.Context) {
	deletedPost := &dto.DeletePostDTO{}

	if err := c.ShouldBindJSON(&deletedPost); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := h.postUsecase.SoftDeleteNews(deletedPost)
	if err != nil {
		if errors.Is(err, httperror.ErrNewsNotFound) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    "SUCCESS_ACCESSED",
				"message": "Post not found",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success Delete Post",
		"data":    nil,
	})
}

func (h *Handler) CreateNews(c *gin.Context) {
	newPostDTO := &dto.NewPostRequestDTO{}
	var validate *validator.Validate = validator.New()

	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Failed to parse multipart form data",
			"data":    nil,
		})
		return
	}

	c.ShouldBind(newPostDTO)

	if newPostDTO.Image.Filename == "" || newPostDTO.Image.Size == 0 || newPostDTO.Image.Header.Get("Content-Type") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Image field is required",
			"data":    nil,
		})
		return
	}

	if newPostDTO.Image.Header.Get("Content-Type") != "image/jpeg" && newPostDTO.Image.Header.Get("Content-Type") != "image/png" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Uploaded file is not an image (JPEG or PNG)",
			"data":    nil,
		})
		return
	}

	err = validate.Struct(newPostDTO)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		var errMsg []string
		for _, fieldError := range validationError {
			errMsg = append(errMsg, util.GetErrorMsg(fieldError))
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": errMsg,
			"data":    nil,
		})
		return
	}

	err = h.postUsecase.CreateNews(newPostDTO)
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidCategory) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid category",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidType) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid type",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success Create New Post",
		"data":    nil,
	})
}

func (h *Handler) EditNews(c *gin.Context) {
	editedPostDTO := &dto.EditPostRequestDTO{}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": httperror.ErrFailedConvertId,
			"data":    nil,
		})
		return
	}

	err = c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Failed to parse multipart form data",
			"data":    nil,
		})
		return
	}

	c.ShouldBind(editedPostDTO)

	err = h.postUsecase.EditNews(editedPostDTO, uint64(idInt))
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidCategory) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid category",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidType) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid type",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success Edit Post",
		"data":    nil,
	})
}

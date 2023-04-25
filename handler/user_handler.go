package handler

import (
	"errors"
	"net/http"
	"stage01-project-backend/dto"
	"stage01-project-backend/httperror"
	"stage01-project-backend/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) Register(c *gin.Context) {
	var newUser *dto.RegisterRequestDTO

	if err := c.ShouldBindJSON(&newUser); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errMsg := make([]util.ErrorMsg, len(ve))
			for i, fe := range ve {
				errMsg[i] = util.ErrorMsg{Message: util.GetErrorMsg(fe)}
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": errMsg,
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err,
			"data":    nil,
		})
		return

	}

	err := h.userUsecase.Register(newUser)
	if err != nil {
		if errors.Is(err, httperror.ErrEmailAlreadyRegistered) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"error":   "SUCCESS_CREATED",
				"message": "Email already registered !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrCreateUser) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"error":   "SUCCESS_CREATED",
				"message": "Failed to create user !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrGenerateHash) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"error":   "SUCCESS_CREATED",
				"message": "Failed to generate hash password !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidEmailFormat) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "SUCCESS_CREATED",
				"message": "Invalid email, please enter this format : 'xxx@xxx.xxx'",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Register successful !",
		"data":    nil,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var loginUserDTO *dto.LoginRequestDTO

	if err := c.ShouldBindJSON(&loginUserDTO); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errMsg := make([]util.ErrorMsg, len(ve))
			for i, fe := range ve {
				errMsg[i] = util.ErrorMsg{Message: util.GetErrorMsg(fe)}
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": errMsg,
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err,
			"data":    nil,
		})
		return
	}

	token, err := h.userUsecase.Login(loginUserDTO)
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidEmailPassword) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Wrong email or password !",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Login successful !",
		"data":    token,
	})
}
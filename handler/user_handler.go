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
	var validate *validator.Validate = validator.New()

	c.ShouldBindJSON(&newUser)
	err := validate.Struct(newUser)
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

	err = h.userUsecase.Register(newUser)
	if err != nil {
		if errors.Is(err, httperror.ErrEmailAlreadyRegistered) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    "ERROR_CREATED",
				"message": "Email already registered !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrCreateUser) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    "ERROR_CREATED",
				"message": "Failed to create user !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrGenerateHash) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    "SUCCESS_CREATED",
				"message": "Failed to generate hash password !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidEmailFormat) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "SUCCESS_CREATED",
				"message": "Invalid email, please enter this format : 'xxx@xxx.xxx'",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidPasswordLength) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "SUCCESS_CREATED",
				"message": "Password length must be 8 or more",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidReferral) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "SUCCESS_CREATED",
				"message": "Invalid referral code",
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
		"message": "Register successful !",
		"data":    nil,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var loginUserDTO *dto.LoginRequestDTO
	var validate *validator.Validate = validator.New()

	c.ShouldBindJSON(&loginUserDTO)
	err := validate.Struct(loginUserDTO)
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

	token, err := h.userUsecase.Login(loginUserDTO)
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidEmailPassword) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   "BAD_REQUEST",
				"message": "Wrong email or password !",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":   "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Login successful !",
		"data":    token,
	})
}

func (h *Handler) GetProfile(c *gin.Context) {
	loggedUserId := c.GetInt("id")

	user, err := h.userUsecase.GetProfile(loggedUserId)
	if err != nil {
		if errors.Is(err, httperror.ErrUserNotFound) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":   "SUCCESS_CREATED",
				"message": "User not found !",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":   "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success get user !",
		"data":    user,
	})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	loggedUserId := c.GetInt("id")
	var editedUser *dto.EditUserRequestDTO
	var validate *validator.Validate = validator.New()

	c.ShouldBindJSON(&editedUser)
	if editedUser.Email != "" {
		err := validate.Var(editedUser.Email, "email")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid email format",
				"data":    nil,
			})
			return
		}
	}

	if editedUser.Phone != "" {
		err := validate.Var(editedUser.Phone, "e164")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid phone number format",
				"data":    nil,
			})
			return
		}
	}

	err := h.userUsecase.UpdateUser(editedUser, loggedUserId)
	if err != nil {
		if errors.Is(err, httperror.ErrEmailAlreadyRegistered) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":   "SUCCESS_CREATED",
				"message": "Email already registered !",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":   "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success edit user !",
		"data":    nil,
	})
}

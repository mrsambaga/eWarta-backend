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

func (h *Handler) CreateNewTransaction(c *gin.Context) {
	userId := c.GetInt("id")
	var validate *validator.Validate = validator.New()
	newTransactionDTO := &dto.TransactionRequestDTO{}

	c.ShouldBindJSON(newTransactionDTO)
	err := validate.Struct(newTransactionDTO)
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

	transaction, err := h.transactionUsecase.CreateNewTransaction(newTransactionDTO, uint64(userId))
	if err != nil {
		if errors.Is(err, httperror.ErrCreateInvoice) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "SUCCESS_ACCESSED",
				"message": "Failed to create invoice",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrCreateTransaction) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "SUCCESS_ACCESSED",
				"message": "Failed to create transaction !",
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
		"message": "Success create transaction !",
		"data":    transaction,
	})
}

func (h *Handler) FindAllUserTransactions(c *gin.Context) {
	userId := c.GetInt("id")

	transactions, err := h.transactionUsecase.FindUserTransactions(userId)
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
		"message": "Success Find User Transactions",
		"data":    transactions,
	})
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	var validate *validator.Validate = validator.New()
	editTransactionDTO := &dto.EditTransactionRequestDTO{}

	c.ShouldBindJSON(editTransactionDTO)
	err := validate.Struct(editTransactionDTO)
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

	err = h.transactionUsecase.UpdateTransaction(editTransactionDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success Update Transactions",
		"data":    nil,
	})
}

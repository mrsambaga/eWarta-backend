package util

import "github.com/go-playground/validator/v10"

type ErrorMsg struct {
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " field is required"
	case "eqfield":
		return fe.Field() + " must be equal to " + fe.Param()
	}
	return "Unknown error"
}

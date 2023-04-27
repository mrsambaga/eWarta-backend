package util

import "github.com/go-playground/validator/v10"

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " field is required"
	case "eqfield":
		return fe.Field() + " must be equal to " + fe.Param()
	case "email":
		return "Invalid email format : please enter xxx@xxx.xxx format"
	case "min":
		return fe.Field() + " must be " + fe.Param() + " characters or more"
	case "e164":
		return "Invalid phone number format, example : +62081283905547"
	}
	return fe.Error()
}

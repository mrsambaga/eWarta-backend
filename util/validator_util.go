package util

import (
	"github.com/go-playground/validator/v10"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " field is required"
	case "eqfield":
		return "Password & Password Confirm is not equal"
	case "email":
		return "Invalid email format : please enter xxx@xxx.xxx format"
	case "min":
		return fe.Field() + " must be " + fe.Param() + " characters or more"
	case "e164":
		return "Invalid phone number format, example : +6281283905547"
	case "eq=user|eq=admin":
		return "Invalid role !"
	}
	return fe.Error()
}

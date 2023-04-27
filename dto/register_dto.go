package dto

type RegisterRequestDTO struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
	Phone           string    `json:"phone" validate:"required,e164"`
	Address         string `json:"address" validate:"required"`
	Role            string `json:"role"`
}

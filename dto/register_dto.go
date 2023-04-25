package dto

type RegisterRequestDTO struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required, eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Address         string `json:"address" binding:"required"`
	Role            string `json:"role"`
}
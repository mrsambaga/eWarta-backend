package dto

import "time"

type TransactionRequestDTO struct {
	Status         string    `json:"status" validate:"required"`
	Total          float64   `json:"total" validate:"required"`
	PaymentDate    time.Time `json:"paymentDate"`
	VoucherId      uint64    `json:"voucherId"`
	SubscriptionId uint64    `json:"subscriptionId" validate:"required"`
}

type TransactionResponseDTO struct {
	TransactionId uint64  `json:"transactionId"`
	Status        string  `json:"status"`
	Total         float64 `json:"total"`
	Subscription  string  `json:"subscription"`
}

type EditTransactionRequestDTO struct {
	TransactionId uint64  `json:"transactionId" validate:"required"`
	Total         float64 `json:"total" validate:"required"`
}

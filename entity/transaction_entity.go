package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id             uint64 `gorm:"PrimaryKey"`
	SubscriptionId uint64
	UserId         uint64
	VoucherId      *uint64
	Status         string
	Total          float64
	PaymentDate    time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Subscription   Subscription
}

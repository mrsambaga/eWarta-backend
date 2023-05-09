package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id             uint64 `gorm:"PrimaryKey"`
	InvoiceId      uint64
	SubscriptionId uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

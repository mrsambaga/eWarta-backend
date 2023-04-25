package entity

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	SubscriptionId uint64 `gorm:"PrimaryKey"`
	Name           string
	Price          float64
	Quota          int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

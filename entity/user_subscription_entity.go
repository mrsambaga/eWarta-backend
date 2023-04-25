package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserSubscription struct {
	Id             uint64 `gorm:"PrimaryKey"`
	UserId         int
	SubscriptionId int
	DateStart      time.Time
	DateEnd        time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

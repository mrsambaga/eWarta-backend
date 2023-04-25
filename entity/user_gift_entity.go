package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserGift struct {
	Id         uint64 `gorm:"PrimaryKey"`
	UserId     uint64
	GiftId     uint64
	TrackingId uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	Id          uint64 `gorm:"PrimaryKey"`
	UserId      uint64
	VoucherId   *uint64
	Status      string
	Total       float64
	PaymentDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

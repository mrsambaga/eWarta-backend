package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserVoucher struct {
	Id          uint64 `gorm:"PrimaryKey"`
	UserId      uint64
	VoucherId   uint64
	DateExpired time.Time
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	Id          uint64 `gorm:"PrimaryKey"`
	Name        string
	Status      string
	VoucherCode string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

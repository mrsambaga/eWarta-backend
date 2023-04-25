package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    uint64 `gorm:"PrimaryKey"`
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	Role      string
	PhoneNo   string
	Quota     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

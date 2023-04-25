package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        uint64 `gorm:"PrimaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

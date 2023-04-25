package entity

import (
	"time"

	"gorm.io/gorm"
)

type Gift struct {
	Id            uint64 `gorm:"PrimaryKey"`
	Status        string
	DateGenerated time.Time
	DateEstimated time.Time
	Name          string
	Description   string
	GeneratedBy   string
	Total         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserPost struct {
	Id        uint64 `gorm:"PrimaryKey"`
	UserId    uint64
	PostId    uint64
	Date      time.Time
	Liked     int
	Shared    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

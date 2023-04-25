package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	PostId      uint64 `gorm:"PrimaryKey"`
	CategoryId  uint64
	Title       string
	Slug        string
	SummaryDesc string
	Content     string
	Type        string
	ImgUrl      string
	AuthorName  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

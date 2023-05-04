package dto

import (
	"time"

	"gorm.io/gorm"
)

type PostDetail struct {
	Title       string `json:"title"`
	SummaryDesc string `json:"summary_desc"`
	ImgUrl      string `json:"img_url"`
	Content     string `json:"content"`
	Author      string `json:"author"`
}

type PostDTO struct {
	PostId      uint64         `json:"postId"`
	Title       string         `json:"title"`
	SummaryDesc string         `json:"summaryDesc"`
	ImgUrl      string         `json:"imgUrl"`
	Content     string         `json:"content"`
	Author      string         `json:"author"`
	Slug        string         `json:"slug"`
	TypeId      uint64         `json:"typeId"`
	CategoryId  uint64         `json:"categoryId:"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

type DeletePostDTO struct {
	PostId uint64 `json:"postId"`
}

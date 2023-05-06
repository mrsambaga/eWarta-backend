package dto

import (
	"mime/multipart"
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

type NewPostRequestDTO struct {
	Title       string               `form:"title" validate:"required"`
	SummaryDesc string               `form:"summaryDesc" validate:"required"`
	Image       multipart.FileHeader `form:"image" validate:"required,file"`
	Content     string               `form:"content" validate:"required"`
	Author      string               `form:"author" validate:"required"`
	Slug        string               `form:"slug" validate:"required"`
	TypeId      uint64               `form:"typeId" validate:"required"`
	CategoryId  uint64               `form:"categoryId" validate:"required"`
}

type DeletePostDTO struct {
	PostId uint64 `json:"postId"`
}

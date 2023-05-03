package repository

import (
	"fmt"
	"stage01-project-backend/constant"
	"stage01-project-backend/entity"
	"stage01-project-backend/httperror"

	"gorm.io/gorm"
)

type PostsRepository interface {
	CreatePost(newPost *entity.Post) error
	GetPosts(*constant.Params) ([]*entity.Post, error)
}

type postsRepositoryImp struct {
	db *gorm.DB
}

type PostsRConfig struct {
	DB *gorm.DB
}

func NewPostRepository(cfg *PostsRConfig) PostsRepository {
	return &postsRepositoryImp{
		db: cfg.DB,
	}
}

func (r *postsRepositoryImp) CreatePost(newPost *entity.Post) error {
	if err := r.db.Create(newPost).Error; err != nil {
		// if errors.Is(err, gorm.ErrDuplicatedKey) {
		// 	return httperror.ErrEmailAlreadyRegistered
		// }

		return httperror.ErrCreatePost
	}

	return nil
}

func (r *postsRepositoryImp) GetPosts(params *constant.Params) ([]*entity.Post, error) {
	fmt.Println(params.Category)
	posts := []*entity.Post{}

	query := r.db.Joins("JOIN categories ON categories.id = posts.category_id").Where("title ILIKE ?", "%"+params.Title+"%")

	if params.Category != "" {
		query = query.Where("categories.name = ? ", params.Category)
	}

	if params.NewsType != "" {
		switch params.NewsType {
		case "free":
			query = query.Where("type_id = ?", 1)
		case "paid":
			query = query.Where("type_id IN (?)", []int{2, 3})
		}
	}

	if params.Date != "" {
		query = query.Order("created_at " + params.Date)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, httperror.ErrFindNews
	}

	return posts, nil
}

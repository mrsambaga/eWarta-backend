package repository

import (
	"stage01-project-backend/entity"
	"stage01-project-backend/httperror"

	"gorm.io/gorm"
)

type PostsRepository interface {
	CreatePost(newPost *entity.Post) error
	GetPosts() ([]*entity.Post, error)
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

func (r *postsRepositoryImp) GetPosts() ([]*entity.Post, error) {
	users := []*entity.Post{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, httperror.ErrFindNews
	}

	return users, nil
}

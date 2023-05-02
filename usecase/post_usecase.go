package usecase

import (
	"stage01-project-backend/entity"
	"stage01-project-backend/repository"
)

type PostsUsecase interface {
	FindAllNews() ([]*entity.Post, error)
}

type postsUsecaseImp struct {
	postsRepository repository.PostsRepository
}

type PostsUsecaseConfig struct {
	PostsRepository repository.PostsRepository
}

func NewPostsUsecase(cfg *PostsUsecaseConfig) PostsUsecase {
	return &postsUsecaseImp{
		postsRepository: cfg.PostsRepository,
	}
}

func (u *postsUsecaseImp) FindAllNews() ([]*entity.Post, error) {
	posts, err := u.postsRepository.GetPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

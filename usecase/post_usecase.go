package usecase

import (
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/repository"
)

type PostsUsecase interface {
	FindAllNews() ([]*entity.Post, error)
	FindAllNewsHighlight() ([]*dto.PostHighlight, error)
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

func (u *postsUsecaseImp) FindAllNewsHighlight() ([]*dto.PostHighlight, error) {
	posts, err := u.postsRepository.GetPosts()
	if err != nil {
		return nil, err
	}

	highlights := make([]*dto.PostHighlight, 0, len(posts))
	for _, post := range posts {
		highlight := &dto.PostHighlight{
			Title:       post.Title,
			SummaryDesc: post.SummaryDesc,
			ImgUrl:      post.ImgUrl,
			Author:      post.AuthorName,
		}
		highlights = append(highlights, highlight)
	}

	return highlights, nil
}

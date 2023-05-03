package usecase

import (
	"stage01-project-backend/constant"
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/repository"
)

type PostsUsecase interface {
	FindAllNews(params *constant.Params) ([]*entity.Post, error)
	FindAllNewsHighlight(*constant.Params) ([]*dto.PostHighlight, error)
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

func (u *postsUsecaseImp) FindAllNews(params *constant.Params) ([]*entity.Post, error) {
	posts, err := u.postsRepository.GetPosts(params)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (u *postsUsecaseImp) FindAllNewsHighlight(params *constant.Params) ([]*dto.PostHighlight, error) {
	posts, err := u.postsRepository.GetPosts(params)
	if err != nil {
		return nil, err
	}

	highlights := make([]*dto.PostHighlight, 0, len(posts))
	for _, post := range posts {
		highlight := &dto.PostHighlight{
			PostId:      post.PostId,
			Title:       post.Title,
			SummaryDesc: post.SummaryDesc,
			ImgUrl:      post.ImgUrl,
			Author:      post.AuthorName,
		}
		highlights = append(highlights, highlight)
	}

	return highlights, nil
}

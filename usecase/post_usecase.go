package usecase

import (
	"stage01-project-backend/constant"
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/repository"
	"stage01-project-backend/util"
)

type PostsUsecase interface {
	FindAllNews(params *constant.Params) ([]*dto.PostDTO, error)
	// FindAllNewsHighlight(*constant.Params) ([]*dto.PostHighlight, error)
	FindNewsDetail(id uint64) (*dto.PostDetail, error)
	SoftDeleteNews(deletedPost *dto.DeletePostDTO) error
	CreateNewPost(newPostDTO *dto.NewPostRequestDTO) error
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

func (u *postsUsecaseImp) FindAllNews(params *constant.Params) ([]*dto.PostDTO, error) {
	posts, err := u.postsRepository.GetPosts(params)
	if err != nil {
		return nil, err
	}

	postDTO := make([]*dto.PostDTO, 0, len(posts))
	for _, post := range posts {
		post := &dto.PostDTO{
			PostId:      post.PostId,
			Title:       post.Title,
			SummaryDesc: post.SummaryDesc,
			ImgUrl:      post.ImgUrl,
			Author:      post.AuthorName,
			Slug:        post.AuthorName,
			TypeId:      post.TypeId,
			CategoryId:  post.CategoryId,
			Content:     post.Content,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			DeletedAt:   post.DeletedAt,
		}
		postDTO = append(postDTO, post)
	}

	return postDTO, nil
}

func (u *postsUsecaseImp) FindNewsDetail(id uint64) (*dto.PostDetail, error) {
	post, err := u.postsRepository.GetPostById(id)
	if err != nil {
		return nil, err
	}

	postDetail := &dto.PostDetail{
		Title:       post.Title,
		SummaryDesc: post.SummaryDesc,
		ImgUrl:      post.ImgUrl,
		Content:     post.Content,
		Author:      post.AuthorName,
	}

	return postDetail, nil
}

func (u *postsUsecaseImp) SoftDeleteNews(deletedPost *dto.DeletePostDTO) error {
	id := deletedPost.PostId

	err := u.postsRepository.SoftDeletePost(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *postsUsecaseImp) CreateNewPost(newPostDTO *dto.NewPostRequestDTO) error {
	cld, err := util.InitiateCloudinary()
	if err != nil {
		return err
	}

	imageURL, err := util.UploadImage(cld, &newPostDTO.Image)
	if err != nil {
		return err
	}

	newPost := &entity.Post{
		Title:       newPostDTO.Title,
		AuthorName:  newPostDTO.Author,
		SummaryDesc: newPostDTO.SummaryDesc,
		Slug:        newPostDTO.Slug,
		ImgUrl:      imageURL,
		TypeId:      newPostDTO.TypeId,
		CategoryId:  newPostDTO.CategoryId,
		Content:     newPostDTO.Content,
	}

	err = u.postsRepository.CreatePost(newPost)
	if err != nil {
		return err
	}

	return nil
}

// func (u *postsUsecaseImp) FindAllNewsHighlight(params *constant.Params) ([]*dto.PostHighlight, error) {
// 	posts, err := u.postsRepository.GetPosts(params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	highlights := make([]*dto.PostHighlight, 0, len(posts))
// 	for _, post := range posts {
// 		highlight := &dto.PostHighlight{
// 			PostId:      post.PostId,
// 			Title:       post.Title,
// 			SummaryDesc: post.SummaryDesc,
// 			ImgUrl:      post.ImgUrl,
// 			Author:      post.AuthorName,
// 		}
// 		highlights = append(highlights, highlight)
// 	}

// 	return highlights, nil
// }



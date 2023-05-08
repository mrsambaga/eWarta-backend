package usecase

import (
	"errors"
	"fmt"
	"stage01-project-backend/constant"
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/repository"
	"stage01-project-backend/util"
)

type PostsUsecase interface {
	FindAllNews(params *constant.Params) ([]*dto.PostDTO, error)
	FindNewsDetail(id uint64) (*dto.PostDetail, error)
	SoftDeleteNews(deletedPost *dto.DeletePostDTO) error
	CreateNews(newPostDTO *dto.NewPostRequestDTO) error
	EditNews(editedPostDTO *dto.EditPostRequestDTO, postId uint64) error
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
			Type:        post.Type.Type,
			Category:    post.Category.Name,
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

	imageUrl, err := u.postsRepository.SoftDeletePost(id)
	if err != nil {
		return err
	}

	publicId := util.GetPublicIdFromUrl(imageUrl)

	cld, err := util.InitiateCloudinary()
	if err != nil {
		return err
	}

	if err := util.DeleteImage(cld, publicId); err != nil {
		return err
	}

	return nil
}

func (u *postsUsecaseImp) CreateNews(newPostDTO *dto.NewPostRequestDTO) error {
	categoryId, err := util.ConvertCategoryToCategoryId(newPostDTO.Category)
	if err != nil {
		return err
	}

	typeId, err := util.ConvertTypeToTypeId(newPostDTO.Type)
	if err != nil {
		return err
	}

	cld, err := util.InitiateCloudinary()
	if err != nil {
		return err
	}

	imageURL := ""
	if newPostDTO.Image.Filename != "" || newPostDTO.Image.Size != 0 || newPostDTO.Image.Header.Get("Content-Type") != "" {
		imageURL, err = util.UploadImage(cld, &newPostDTO.Image)
		if err != nil {
			return err
		}
	}

	newPost := &entity.Post{}

	newPost = &entity.Post{
		Title:       newPostDTO.Title,
		AuthorName:  newPostDTO.Author,
		SummaryDesc: newPostDTO.SummaryDesc,
		Slug:        newPostDTO.Slug,
		TypeId:      typeId,
		CategoryId:  categoryId,
		Content:     newPostDTO.Content,
		ImgUrl:      imageURL,
	}

	err = u.postsRepository.CreatePost(newPost)
	if err != nil {
		return err
	}

	return nil
}

func (u *postsUsecaseImp) EditNews(editedPostDTO *dto.EditPostRequestDTO, postId uint64) error {
	categoryId, err := util.ConvertCategoryToCategoryId(editedPostDTO.Category)
	if err != nil {
		return err
	}

	typeId, err := util.ConvertTypeToTypeId(editedPostDTO.Type)
	if err != nil {
		return err
	}

	cld, err := util.InitiateCloudinary()
	if err != nil {
		return err
	}

	newImageURL := ""
	if editedPostDTO.Image.Filename != "" || editedPostDTO.Image.Size != 0 || editedPostDTO.Image.Header.Get("Content-Type") != "" {

		if editedPostDTO.Image.Header.Get("Content-Type") != "image/jpeg" && editedPostDTO.Image.Header.Get("Content-Type") != "image/png" {
			return errors.New("upload file is not a jpg or png")
		}

		newImageURL, err = util.UploadImage(cld, &editedPostDTO.Image)
		if err != nil {
			return err
		}
	}

	existingPost, err := u.postsRepository.GetPostById(postId)
	if err != nil {
		return err
	}
	editedPost := &entity.Post{}

	if newImageURL != "" {
		editedPost = &entity.Post{
			PostId:      postId,
			Title:       editedPostDTO.Title,
			AuthorName:  editedPostDTO.Author,
			SummaryDesc: editedPostDTO.SummaryDesc,
			Slug:        editedPostDTO.Slug,
			TypeId:      typeId,
			CategoryId:  categoryId,
			Content:     editedPostDTO.Content,
			ImgUrl:      newImageURL,
		}
	} else if newImageURL == "" {
		editedPost = &entity.Post{
			PostId:      postId,
			Title:       editedPostDTO.Title,
			AuthorName:  editedPostDTO.Author,
			SummaryDesc: editedPostDTO.SummaryDesc,
			Slug:        editedPostDTO.Slug,
			TypeId:      typeId,
			CategoryId:  categoryId,
			Content:     editedPostDTO.Content,
			ImgUrl:      existingPost.ImgUrl,
		}
	}

	fmt.Println("EDITED POST : ", editedPost)

	err = u.postsRepository.UpdatePost(editedPost)
	if err != nil {
		return err
	}

	return nil
}

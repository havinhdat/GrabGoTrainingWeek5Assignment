package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/repository"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
)

// GetPostsV1 usecase
type getPostsV1 struct {
	Repo repository.PostRepository
}

// NewGetPostsUsecase func
func NewGetPostsUsecase(repo repository.PostRepository) GetPostsUsecase {
	return &getPostsV1 {
		Repo: repo,
	}
}

// GetPosts GetPostsUsecase
func (usecase *getPostsV1) GetPosts() ([]entity.Post, error) {
	return usecase.Repo.GetPosts()
}


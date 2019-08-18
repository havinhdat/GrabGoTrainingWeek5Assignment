package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/repository"
)

// GetPostsUsecase usecase
type GetPostsUsecase struct {
	Repo repository.PostRepository
}

// NewGetPostsUsecase func
func NewGetPostsUsecase(repo repository.PostRepository) *GetPostsUsecase {
	return &GetPostsUsecase {
		Repo: repo,
	}
}

// Execute GetPostsUsecase
func (usecase *GetPostsUsecase) Execute(params... interface{}) (interface{}, error) {
	return usecase.Repo.GetPosts()
}


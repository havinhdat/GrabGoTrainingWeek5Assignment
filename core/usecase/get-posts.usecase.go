package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/repository"
)

// GetPostsUsecase usecase
type GetPostsUsecase struct {
	Repo repository.PostRepository
}

// Execute GetPostsUsecase
func (usecase *GetPostsUsecase) Execute(params... interface{}) (interface{}, error) {
	return usecase.Repo.GetPosts()
}


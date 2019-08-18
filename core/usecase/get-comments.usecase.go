package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/repository"
)

// GetCommentsUsecase usecase
type GetCommentsUsecase struct {
	Repo repository.CommentRepository
}

// Execute GetCommentsUsecase
func (usecase *GetCommentsUsecase) Execute(params... interface{}) (interface{}, error) {
	return usecase.Repo.GetComments()
}


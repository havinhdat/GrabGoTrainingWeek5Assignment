package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/repository"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
)

// GetCommentsV1 usecase
type getCommentsV1 struct {
	Repo repository.CommentRepository
}

// NewGetCommentsUsecase func
func NewGetCommentsUsecase(repo repository.CommentRepository) GetCommentsUsecase {
	return &getCommentsV1 {
		Repo: repo,
	}
}

// Execute GetCommentsUsecase
func (usecase *getCommentsV1) GetComments() ([]entity.Comment, error) {
	return usecase.Repo.GetComments()
}

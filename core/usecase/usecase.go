package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
)

// GetCommentsUsecase interface for usecases
type GetCommentsUsecase interface {
	GetComments() ([]entity.Comment, error)
}

type GetPostsUsecase interface {
	GetPosts() ([]entity.Post, error)
}
type GetPostsWithCommentsUsecase interface {
	GetPostsWithComments() (*entity.PostsWithComments, error)
}
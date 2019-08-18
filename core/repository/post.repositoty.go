package repository

import(
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
)

type PostRepository interface {
	GetPosts() ([]entity.Post, error)
}
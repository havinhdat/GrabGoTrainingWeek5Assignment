package repository

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
)

type CommentRepository interface {
	GetComments() ([]entity.Comment, error)
}
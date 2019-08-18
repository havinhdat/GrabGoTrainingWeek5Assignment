package postwithcomment

import "grab/week5/GrabGoTrainingWeek5Assignment/comment"

type PostWithComment struct {
	ID       int64             `json:"id"`
	Title    string            `json:"string"`
	Comments []comment.Comment `json:"comments,omitempty"`
}

type PostWithCommentsInterface interface {
	GetPostWithComments() ([]PostWithComment, error)
}

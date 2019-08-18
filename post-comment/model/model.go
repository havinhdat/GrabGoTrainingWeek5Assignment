package model

import comment_model "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/comment/model"

type PostWithComments struct {
	ID       int64                   `json:"id"`
	Title    string                  `json:"string"`
	Comments []comment_model.Comment `json:"comments,omitempty"`
}

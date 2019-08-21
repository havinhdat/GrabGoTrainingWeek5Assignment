package dataloader

import (
	"../defination"
	"../interfaces"
)

type CombineSeviceLoader struct {
	data interfaces.CombineData
}

func (service *CombineSeviceLoader) CombinePostWithComments(posts []defination.Post, comments []defination.Comment) []defination.PostWithComments {
	commentsByPostID := map[int64][]defination.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]defination.PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, defination.PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}
	return result
}

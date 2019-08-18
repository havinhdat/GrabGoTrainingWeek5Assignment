package postcomment

import (
	comment_getter "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/comment"
	post_getter "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/post"
)

type PostWithComments struct {
	ID       int64                    `json:"id"`
	Title    string                   `json:"string"`
	Comments []comment_getter.Comment `json:"comments,omitempty"`
}

type PostWithCommentsCombiner interface {
	CombinePostWithComments(posts []post_getter.Post, comments []comment_getter.Comment) []PostWithComments
}

type PostWithCommentsCombinerImpl struct{}

func (postCommentCombinerImpl *PostWithCommentsCombinerImpl) CombinePostWithComments(posts []post_getter.Post, comments []comment_getter.Comment) []PostWithComments {
	commentsByPostID := map[int64][]comment_getter.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result
}

func NewPostWithCommentsCombinerImpl() PostWithCommentsCombiner {
	return &PostWithCommentsCombinerImpl{}
}

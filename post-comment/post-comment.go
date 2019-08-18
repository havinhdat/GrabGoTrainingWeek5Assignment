package postcomment

import (
	comment_model "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/comment/model"
	post_comments_model "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/post-comment/model"
	post_model "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/post/model"
)

type PostWithCommentsCombiner interface {
	CombinePostWithComments(posts []post_model.Post, comments []comment_model.Comment) []post_comments_model.PostWithComments
}

type PostWithCommentsCombinerImpl struct{}

func (postCommentCombinerImpl *PostWithCommentsCombinerImpl) CombinePostWithComments(posts []post_model.Post,
	comments []comment_model.Comment) []post_comments_model.PostWithComments {
	commentsByPostID := map[int64][]comment_model.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]post_comments_model.PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, post_comments_model.PostWithComments{
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

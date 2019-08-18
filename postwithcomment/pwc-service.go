package postwithcomment

import (
	"grab/week5/GrabGoTrainingWeek5Assignment/comment"
	"grab/week5/GrabGoTrainingWeek5Assignment/post"
)

type PostWithCommentService struct {
	postService    post.PostInterface
	commentService comment.CommentInterface
}

func NewPostWithCommentService(postService post.PostInterface, commentService comment.CommentInterface) *PostWithCommentService {
	service := &PostWithCommentService{postService, commentService}
	return service
}

func (service *PostWithCommentService) GetPostWithComments() ([]PostWithComment, error) {
	// Get posts from api
	posts, err := service.postService.GetPosts()
	if err != nil {
		return nil, err
	}

	// Get comments from api
	comments, err := service.commentService.GetComments()
	if err != nil {
		return nil, err
	}

	// Combine and return response
	postWithComments := combinePostWithComments(posts, comments)
	return postWithComments, nil
}

func combinePostWithComments(posts []post.Post, comments []comment.Comment) []PostWithComment {
	commentsByPostID := map[int64][]comment.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]PostWithComment, 0, len(posts))
	for _, post := range posts {
		result = append(result, PostWithComment{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result
}

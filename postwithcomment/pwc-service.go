package postwithcomment

import (
	"grab/week5/GrabGoTrainingWeek5Assignment/comment"
	"grab/week5/GrabGoTrainingWeek5Assignment/post"
)

type PostInterface interface {
	GetPosts() ([]post.Post, error)
}

type CommentInterface interface {
	GetComments() ([]comment.Comment, error)
}

type PostWithCommentService struct {
	postService    PostInterface
	commentService CommentInterface
}

func NewPostWithCommentService(postService PostInterface, commentService CommentInterface) *PostWithCommentService {
	service := &PostWithCommentService{postService, commentService}
	return service
}

func (service *PostWithCommentService) GetPostWithComments() ([]PostWithComments, error) {
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

func combinePostWithComments(posts []post.Post, comments []comment.Comment) []PostWithComments {
	commentsByPostID := map[int64][]comment.Comment{}
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

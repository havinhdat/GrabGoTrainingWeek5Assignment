package postwithcomment

import (
	"grab/week5/GrabGoTrainingWeek5Assignment/comment"
	"grab/week5/GrabGoTrainingWeek5Assignment/post"
)

type PostWithCommentsResponse struct {
	Posts []PostWithComments `json:"posts"`
}

type PostWithComments struct {
	ID       int64             `json:"id"`
	Title    string            `json:"string"`
	Comments []comment.Comment `json:"comments,omitempty"`
}

type PostWithCommentsInterface interface {
	GetPostWithComments() ([]PostWithComments, error)
}

type PostWithCommentService struct {
	postService    post.PostInterface
	commentService comment.CommentInterface
}

func NewPostWithCommentService(postService post.PostInterface, commentService comment.CommentInterface) *PostWithCommentService {
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

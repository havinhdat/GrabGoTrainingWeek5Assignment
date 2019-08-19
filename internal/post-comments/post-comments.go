package postcomments

import (
	"grab/internal/comment"
	"grab/internal/post"
)

type PostWithCommentsResponse struct {
	Posts []PostWithComments `json:"posts"`
}

type PostWithComments struct {
	ID       int64     `json:"id"`
	Title    string    `json:"string"`
	Comments []comment.Comment `json:"comments,omitempty"`
}

type PostCommentsWorker interface {
	GetPostWithComments() (PostWithCommentsResponse, error)
	combinePostWithComments([]post.Post, []comment.Comment) []PostWithComments
}

type PostCommentsWorkerImp struct {}

func CreatePostCommentsWorker() *PostCommentsWorkerImp{
	return &PostCommentsWorkerImp{}
}
func (pwc *PostCommentsWorkerImp) GetPostWithComments() (PostWithCommentsResponse, error) {

	postwker := post.CreatePostWorker()
	posts, error := postwker.GetPosts()
	if error != nil {
		return PostWithCommentsResponse{}, error
	}
	
	cmtwker := comment.CreateCommentsWorker()
	comments, error := cmtwker.GetComments()
	if error != nil {
		return PostWithCommentsResponse{}, error
	}

	postWithComments := pwc.combinePostWithComments(posts, comments)

	resp := PostWithCommentsResponse{Posts: postWithComments}
	
	return resp, nil

}

func (pwc *PostCommentsWorkerImp) combinePostWithComments(posts []post.Post, comments []comment.Comment) []PostWithComments {
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


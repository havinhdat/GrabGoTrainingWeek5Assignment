package comment

import (
	"encoding/json"
	"grab/internal/httpworker"
	"grab/config/endpoint"
)

// Export Comment struct
type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}

type CommentsWorker interface {
	GetComments() ([]Comment, error)
}

type CommentsWorkerImp struct {}

func CreateCommentsWorker() *CommentsWorkerImp{
	return &CommentsWorkerImp{}
}

// Get all comment from endpoint.
func (cmtwker *CommentsWorkerImp) GetComments() ([]Comment, error) {

	httpWorker := httpworker.CreateHTTPWorker()
	body, error := httpWorker.Get(endpoint.GetCommentsEndpoint)
	if error != nil {
		return nil, error
	}
	var comments []Comment
	if error = json.Unmarshal(body, &comments); error != nil {
		return nil, error
	}

	return comments, nil
}
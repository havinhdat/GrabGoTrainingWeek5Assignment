package comment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}

type CommentInterface interface {
	GetComments() ([]Comment, error)
}

type CommentService struct{}

func (commentService *CommentService) GetComments() ([]Comment, error) {
	resp, err := http.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

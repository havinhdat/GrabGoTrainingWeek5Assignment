package comment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CommentService struct {
	endpoint string
}

func NewCommentService(endpoint string) *CommentService {
	service := &CommentService{endpoint}
	return service
}

func (commentService *CommentService) GetComments() ([]Comment, error) {
	resp, err := http.Get(commentService.endpoint)
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

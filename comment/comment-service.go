package comment

import (
	"encoding/json"
	"grab/week5/GrabGoTrainingWeek5Assignment/httpclient"
	"io/ioutil"
)

const (
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type CommentService struct {
	httpClient httpclient.HTTPClient
}

func NewCommentService(httpClient httpclient.HTTPClient) *CommentService {
	service := &CommentService{httpClient}
	return service
}

func (commentService *CommentService) GetComments() ([]Comment, error) {
	resp, err := commentService.httpClient.Get(getCommentsEndpoint)
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

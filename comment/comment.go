package comment

import (
	"encoding/json"
	"io/ioutil"

	comment_model "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/comment/model"
	httpclient "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/http-client"
)

const (
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type CommentGetter interface {
	GetComments() ([]comment_model.Comment, error)
}

type CommentGetterImpl struct {
	httpClient httpclient.HTTPClient
}

func (commentGetter *CommentGetterImpl) GetComments() ([]comment_model.Comment, error) {
	resp, err := commentGetter.httpClient.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []comment_model.Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func New(httpClient httpclient.HTTPClient) (CommentGetter, error) {
	return &CommentGetterImpl{
		httpClient: httpClient,
	}, nil
}

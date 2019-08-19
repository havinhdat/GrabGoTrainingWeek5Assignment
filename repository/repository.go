package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"bathien/assignment/entity"
	httpclient "bathien/assignment/http-client"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

//go:generate mockery -name=DataGetter
type DataGetter interface {
	GetPosts() ([]entity.Post, error)
	GetComments() ([]entity.Comment, error)
}

type dataGetterImpl struct {
	httpClient httpclient.HTTPClient
}

func (dataGetter *dataGetterImpl) GetComments() ([]entity.Comment, error) {
	resp, err := dataGetter.httpClient.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []entity.Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func (dataGetter *dataGetterImpl) GetPosts() ([]entity.Post, error) {
	resp, err := dataGetter.httpClient.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []entity.Post
	if err = json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

type Option func(dataGetter *dataGetterImpl)

func New(options ...Option) (DataGetter, error) {
	dataGetter := &dataGetterImpl{}
	for _, o := range options {
		o(dataGetter)
	}
	if dataGetter.httpClient == nil {
		return nil, errors.New("missing http client")
	}
	return dataGetter, nil
}

func WithHTTPClient(httpClient httpclient.HTTPClient) Option {
	return func(dataGetter *dataGetterImpl) {
		dataGetter.httpClient = httpClient
	}
}

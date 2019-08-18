package post

import (
	"encoding/json"
	"io/ioutil"

	httpclient "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/http-client"
)

const (
	getPostsEndpoint = "https://my-json-server.typicode.com/typicode/demo/posts"
)

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type PostGetter interface {
	GetPosts() ([]Post, error)
}

type PostGetterImpl struct {
	httpClient httpclient.HTTPClient
}

func (postGetter *PostGetterImpl) GetPosts() ([]Post, error) {
	resp, err := postGetter.httpClient.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []Post
	if err = json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func New(httpClient httpclient.HTTPClient) (PostGetter, error) {
	return &PostGetterImpl{
		httpClient: httpClient,
	}, nil
}

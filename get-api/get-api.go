package getapi

import (
	"encoding/json"
	"io/ioutil"

	httpclient "../http-client"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}

type GetApi interface {
	GetPosts() ([]Post, error)
	GetComments() ([]Comment, error)
}

type getApiImpl struct {
	httpClient httpclient.HTTPClient
}

func (api *getApiImpl) GetPosts() ([]Post, error) {
	resp, err := api.httpClient.Get(getPostsEndpoint)
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

func (api *getApiImpl) GetComments() ([]Comment, error) {
	resp, err := api.httpClient.Get(getCommentsEndpoint)
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

func New(httpClient httpclient.HTTPClient) (GetApi, error) {
	return &getApiImpl{httpClient: httpClient}, nil
}

package post

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	getPostsEndpoint = "https://my-json-server.typicode.com/typicode/demo/posts"
)

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type PostInterface interface {
	GetPosts() ([]Post, error)
}

type PostService struct{}

func (postService *PostService) GetPosts() ([]Post, error) {
	resp, err := http.Get(getPostsEndpoint)
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

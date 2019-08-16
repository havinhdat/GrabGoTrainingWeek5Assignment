package post

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PostService struct {
	endpoints string
}

func NewPostService(endpoint string) *PostService {
	service := &PostService{endpoint}
	return service
}

func (postService *PostService) GetPosts() ([]Post, error) {
	resp, err := http.Get(postService.endpoints)
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

package post

import (
	"encoding/json"
	"grab/week5/GrabGoTrainingWeek5Assignment/httpclient"
	"io/ioutil"
)

const (
	getPostsEndpoint = "https://my-json-server.typicode.com/typicode/demo/posts"
)

type PostService struct {
	httpClient httpclient.HTTPClient
}

func NewPostService(httpClient httpclient.HTTPClient) *PostService {
	service := &PostService{httpClient}
	return service
}

func (postService *PostService) GetPosts() ([]Post, error) {
	resp, err := postService.httpClient.Get(getPostsEndpoint)
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

package post


import (
	"encoding/json"
	"grab/service/httpworker"
	"grab/config/endpoint"
)

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type PostWorker interface {
	GetPosts() ([]Post, error)
}

type PostWorkerImp struct {}

// PostWorker Contructor
func CreatePostWorker() *PostWorkerImp{
	return &PostWorkerImp{}
}
// Get all post from post endpoint
func (post *PostWorkerImp) GetPosts() ([]Post, error) {

	httpWorker := httpworker.CreateHTTPWorker()
	body, error := httpWorker.Get(endpoint.GetPostsEndpoint)
	if error != nil {
		return nil, error
	}
	var posts []Post
	if error = json.Unmarshal(body, &posts); error != nil {
		return nil, error
	}

	return posts, nil
}
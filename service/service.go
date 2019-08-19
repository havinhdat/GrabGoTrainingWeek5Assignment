package service

import (
	"encoding/json"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/model"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/service/HttpClient"
	"io/ioutil"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

///go:generate mockery -name=Service -inpkg
type Service interface {
	GetPostWithComments() ([]model.PostWithComments, error)
}

type serviceImpl struct {
	HttpClient HttpClient.HttpClient
}

func NewService(HttpClient HttpClient.HttpClient) Service {
	return &serviceImpl{HttpClient: HttpClient}
}

func (s *serviceImpl) GetPostWithComments() (postWithComments []model.PostWithComments, err error) {
	commentsByPostID := map[int64][]model.Comment{}
	comments, err := s.getComments()
	if err != nil {
		return
	}
	posts, err := s.getPosts()
	if err != nil {
		return
	}
	if len(comments) > 0 {
		for _, comment := range comments {
			commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
		}
	}

	if len(posts) > 0 {
		postWithComments = combinePostWithComments(posts, comments)
	}
	return
}

func (s *serviceImpl) getPosts() (posts []model.Post, err error) {
	resp, err := s.HttpClient.Get(getPostsEndpoint)
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(body, &posts); err != nil {
		return
	}
	return posts, nil
}

func (s *serviceImpl) getComments() (comments []model.Comment, err error) {
	resp, err := s.HttpClient.Get(getCommentsEndpoint)
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	defer resp.Body.Close()
	if err = json.Unmarshal(body, &comments); err != nil {
		return
	}
	return comments, nil
}

func combinePostWithComments(posts []model.Post, comments []model.Comment) []model.PostWithComments {
	commentsByPostID := map[int64][]model.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]model.PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, model.PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result
}

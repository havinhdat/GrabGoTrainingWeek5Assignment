package service

import (
	"encoding/json"
	"grab/internal/bootcamp/model"
	"io/ioutil"
	"net/http"
)

type BlogService interface {
	GetPostWithComments() (*[]model.PostWithComments, error)
}

type BlogServiceImpl struct {
	postsEndpoint    string
	commentsEndpoint string
}

func NewBlogServiceImpl(postsEndpoint string, commentsEndpoint string) *BlogServiceImpl {
	return &BlogServiceImpl{postsEndpoint: postsEndpoint, commentsEndpoint: commentsEndpoint}
}

func (p *BlogServiceImpl) GetPostWithComments() (*[]model.PostWithComments, error) {
	commentsByPostID := map[int64][]model.Comment{}

	comments, err := p.getComments()
	if err != nil {
		return nil, err
	}
	posts, err := p.getPosts()
	if err != nil {
		return nil, err
	}

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

	return &result, nil
}

func (p *BlogServiceImpl) getPosts() ([]model.Post, error) {
	resp, err := http.Get(p.postsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []model.Post
	if err = json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *BlogServiceImpl) getComments() ([]model.Comment, error) {
	resp, err := http.Get(p.commentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []model.Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

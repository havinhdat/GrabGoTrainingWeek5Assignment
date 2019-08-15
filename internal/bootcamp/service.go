package bootcamp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Blog interface {
	GetPostWithComments() (*[]PostWithComments, error)
}

type BlogImpl struct {
	postsEndpoint    string
	commentsEndpoint string
}

func NewBlogImpl(postsEndpoint string, commentsEndpoint string) *BlogImpl {
	return &BlogImpl{postsEndpoint: postsEndpoint, commentsEndpoint: commentsEndpoint}
}

func (p *BlogImpl) GetPostWithComments() (*[]PostWithComments, error) {
	commentsByPostID := map[int64][]Comment{}

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

	result := make([]PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return &result, nil
}

func (p *BlogImpl) getPosts() ([]Post, error) {
	resp, err := http.Get(p.postsEndpoint)
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

func (p *BlogImpl) getComments() ([]Comment, error) {
	resp, err := http.Get(p.commentsEndpoint)
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

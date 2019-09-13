package business

import (
	apiservice "GrabBootCamp2019/GrabGoTrainingWeek5Assignment/api"
	"GrabBootCamp2019/GrabGoTrainingWeek5Assignment/model"
	"errors"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

//ApiService provides posts and comments
type ApiService interface {
	GetPosts(url string) ([]model.Post, error)
	GetComments(url string) ([]model.Comment, error)
}

//BusinessServiceImpl does logic works
type BusinessServiceImpl struct {
	api ApiService
}

//CombineData gets posts and comments from server then adds corresponding comments into posts
func (bs *BusinessServiceImpl) CombineData() (model.PostWithCommentsResponse, error) {
	bs.api = new(apiservice.PostCommentApi)

	posts, err := bs.api.GetPosts(getPostsEndpoint)
	if err != nil {
		return model.PostWithCommentsResponse{}, errors.New("Can't get posts")
	}
	comments, err := bs.api.GetComments(getCommentsEndpoint)
	if err != nil {
		return model.PostWithCommentsResponse{}, errors.New("Can't get comments")
	}

	m := make(map[int64][]model.Comment)

	for _, c := range comments {
		m[c.PostID] = append(m[c.PostID], c)
	}

	var pwcArr model.PostWithCommentsResponse
	for _, post := range posts {
		pwc := model.PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: m[post.ID],
		}
		pwcArr.Posts = append(pwcArr.Posts, pwc)
	}

	return pwcArr, nil
}

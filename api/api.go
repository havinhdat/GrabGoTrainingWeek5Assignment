package apiservice

import (
	"GrabBootCamp2019/GrabGoTrainingWeek5Assignment/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//PostCommentApi retrieves posts and comments from server
type PostCommentApi struct {
}

func readBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Unable to connect to " + url)
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	return ioutil.ReadAll(resp.Body)
}

//GetPosts retrieve all posts from server
func (a PostCommentApi) GetPosts(url string) ([]model.Post, error) {
	body, err := readBody(url)

	var posts []model.Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, errors.New("Incorrect json format, unable to unmarshal")
	}
	return posts, nil
}

//GetComments retrieve all comments from server
func (a PostCommentApi) GetComments(url string) ([]model.Comment, error) {
	body, err := readBody(url)

	var comments []model.Comment
	err = json.Unmarshal(body, &comments)
	if err != nil {
		return nil, errors.New("Incorrect json format, unable to unmarshal")
	}
	return comments, nil
}

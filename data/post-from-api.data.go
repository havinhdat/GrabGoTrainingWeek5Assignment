package data

import(
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/mapper"
	"net/http"
	"io/ioutil"
)

// PostData data
type PostData struct {
	Mapper mapper.Mapper
}

// GetPosts get all posts
func (data *PostData) GetPosts() ([]entity.Post, error) {
	resp, err := http.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []entity.Post
	if err = data.Mapper.Decode(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
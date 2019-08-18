package data

import(
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/data/dataformatter"
	"net/http"
	"io/ioutil"
)

// PostData data
type PostData struct {
	DataFormatter dataformatter.DataFormatter
}

// NewPostData func
func NewPostData(dataformatter dataformatter.DataFormatter) *PostData {
	return &PostData{
		DataFormatter: dataformatter,
	}
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
	if err = data.DataFormatter.Decode(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
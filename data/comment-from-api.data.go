package data

import(
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/data/dataformatter"
	"net/http"
	"io/ioutil"
)

// CommentData data
type CommentData struct {
	DataFormatter dataformatter.DataFormatter
}

// NewCommentData func
func NewCommentData(dataformatter dataformatter.DataFormatter) *CommentData {
	return &CommentData{
		DataFormatter: dataformatter,
	}
}

// GetComments get all comments
func (data *CommentData) GetComments() ([]entity.Comment, error) {
	resp, err := http.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []entity.Comment
	if err = data.DataFormatter.Decode(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}
package comment_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"grab/internal/comment"
	"grab/internal/comment/mocks"
	"encoding/json"
	"fmt"
)

func TestGetComments(t *testing.T) {
	mockCommentRep := &mocks.CommentsWorker{}
	var getComments []comment.Comment
	// var getComments comment.Comment
	// assert := assert.New(t)

	mockCommentRep.On("GetComments").Return(getComments, nil)
	// service := ToDo{ToDoRepo: mockToDoRep}
	fmt.Println(getComments)
	// res, err := service.GetTodo(nil, req)

	trueData := []byte(`[ { "id": 1, "body": "some comment", "postId": 1 }, { "id": 2, "body": "some comment", "postId": 1 } ]`)

	var expectedRes []comment.Comment
	error := json.Unmarshal(trueData, &expectedRes)

	fmt.Println(expectedRes)
	assert.Nil(t, error)
	assert.Equal(t, expectedRes, expectedRes)
	mockCommentRep.AssertExpectations(t)
}
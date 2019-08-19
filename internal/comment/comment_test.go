package comment_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"grab/internal/comment"
	"grab/internal/comment/mocks"
	"encoding/json"
)

func TestGetComments(t *testing.T) {
	mockCommentRep := &mocks.CommentsWorker{}
	var getComments []comment.Comment
	var expectedRes []comment.Comment

	mockCommentRep.On("GetComments").Return(getComments, nil)
	trueData := []byte(`[ { "id": 1, "body": "some comment", "postId": 1 }, { "id": 2, "body": "some comment", "postId": 1 } ]`)
	error := json.Unmarshal(trueData, &expectedRes)

	assert.Nil(t, error)
	assert.Equal(t, expectedRes, expectedRes)
	mockCommentRep.AssertExpectations(t)
}
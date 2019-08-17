package endpoint

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	_ "github.com/stretchr/testify/mock"
	tmocks "grab/internal/bootcamp/encode/mocks"
	"grab/internal/bootcamp/model"
	"grab/internal/bootcamp/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestBlogEndpoint_GetPostsWithComments(t *testing.T) {
	blogService := &mocks.BlogService{}
	encodeResponse := &tmocks.EncodeResponse{}
	endpoint := NewBlogEndpointImpl(blogService, encodeResponse)

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("get", "localhost:8080", nil)

	var res *[]model.PostWithComments

	blogService.On("GetPostWithComments").Return(res, nil)
	encodeResponse.On("Encode", mock.Anything, mock.Anything).Return(nil)

	endpoint.GetPostsWithComments(writer, request)

	blogService.AssertExpectations(t)

	assert.Equal(t, 200, writer.Code)
}

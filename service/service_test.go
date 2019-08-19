package service

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/model"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/service/HttpClient"
)

func TestNewService(t *testing.T) {

}

func Test_serviceImpl_GetPostWithComments(t *testing.T) {
	mockPost := []model.Post{{ID: 1, Title: "title test"}}
	mockComment := []model.Comment{{ID: 1, Body: "body comment test", PostID: mockPost[0].ID}}
	expect := []model.PostWithComments{{ID: mockPost[0].ID, Title: mockPost[0].Title, Comments: []model.Comment{{ID: mockComment[0].ID, Body: mockComment[0].Body, PostID: mockPost[0].ID}}}}
	tests := []struct {
		name                 string
		wantPostWithComments []model.PostWithComments
		expectErr            error
		doMock               func(httpClientMocks *HttpClient.MockHttpClient)
	}{
		{"test ok ", expect, nil, func(httpClientMocks *HttpClient.MockHttpClient) {
			bytes, _ := json.Marshal(mockPost)
			respMocks := &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
			}
			httpClientMocks.On("Get", getPostsEndpoint).Return(respMocks, nil)
			bytes, _ = json.Marshal(mockComment)
			respMocks = &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
			}
			httpClientMocks.On("Get", getCommentsEndpoint).Return(respMocks, nil)
		}},
		//{"test fail when get post ", nil, errors.New("anyError"), func(httpClientMocks *HttpClient.MockHttpClient) {
		//	bytes, _ := json.Marshal(mockComment)
		//	respMocks := &http.Response{
		//		Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
		//	}
		//	httpClientMocks.On("Get", getCommentsEndpoint).Return(respMocks, nil)
		//	httpClientMocks.On("Get", getPostsEndpoint).Return(nil, errors.New("anyError"))
		//}},
		{"test fail when get comment", nil, errors.New("anyError"), func(httpClientMocks *HttpClient.MockHttpClient) {
			httpClientMocks.On("Get", getCommentsEndpoint).Return(nil, errors.New("anyError"))
		}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			httpClientMocks := &HttpClient.MockHttpClient{}
			service := NewService(httpClientMocks)

			if tc.doMock != nil {
				tc.doMock(httpClientMocks)
			}

			postWithComments, err := service.GetPostWithComments()

			assert.Equal(t, tc.wantPostWithComments, postWithComments)
			assert.Equal(t, tc.expectErr, err)

		})
	}
}

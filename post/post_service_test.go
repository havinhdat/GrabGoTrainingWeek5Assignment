package post

import (
	"encoding/json"
	"errors"
	"grab/week5/GrabGoTrainingWeek5Assignment/httpclient"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPost(t *testing.T) {
	testcases := []getPostScenario{
		getCommentHappyPath(),
		getCommentFailed(),
	}
	for _, tc := range testcases {
		testCase := tc
		t.Run(testCase.Desc, func(t *testing.T) {
			httpMock := &httpclient.MockHTTPClient{}
			postService := NewPostService(httpMock)

			if testCase.DoMock != nil {
				testCase.DoMock(httpMock)
			}

			resp, err := postService.GetPosts()

			assert.Equal(t, testCase.ExpectedResult, resp)
			assert.Equal(t, testCase.ExpectedError, err)
		})
	}
}

type getPostScenario struct {
	Desc           string
	DoMock         func(httpClientMock *httpclient.MockHTTPClient)
	ExpectedResult []Post
	ExpectedError  error
}

func getCommentHappyPath() getPostScenario {
	expectedResult := []Post{
		Post{
			ID:    1,
			Title: "Hello",
		},
		Post{
			ID:    2,
			Title: "Bello",
		},
	}
	return getPostScenario{
		Desc: "Get posts happy path",
		DoMock: func(mock *httpclient.MockHTTPClient) {
			listMocks := expectedResult
			bytes, _ := json.Marshal(listMocks)
			respMocks := &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
			}
			mock.On("Get", getPostsEndpoint).Return(respMocks, nil)
		},
		ExpectedResult: expectedResult,
		ExpectedError:  nil,
	}
}

func getCommentFailed() getPostScenario {
	expectedError := errors.New("Failed to get posts")
	return getPostScenario{
		Desc: "Get post failed",
		DoMock: func(mock *httpclient.MockHTTPClient) {
			mock.On("Get", getPostsEndpoint).Return(nil, expectedError)
		},
		ExpectedResult: nil,
		ExpectedError:  expectedError,
	}
}

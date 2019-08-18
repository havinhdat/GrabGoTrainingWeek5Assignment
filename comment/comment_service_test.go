package comment

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

func TestGetComment(t *testing.T) {
	testcases := []getCommentScenario{
		getCommentHappyPath(),
		getCommentFailed(),
	}
	for _, tc := range testcases {
		testCase := tc
		t.Run(testCase.Desc, func(t *testing.T) {
			httpMock := &httpclient.MockHTTPClient{}
			commentService := NewCommentService(httpMock)

			if testCase.DoMock != nil {
				testCase.DoMock(httpMock)
			}

			resp, err := commentService.GetComments()

			assert.Equal(t, testCase.ExpectedResult, resp)
			assert.Equal(t, testCase.ExpectedError, err)
		})
	}
}

type getCommentScenario struct {
	Desc           string
	DoMock         func(httpClientMock *httpclient.MockHTTPClient)
	ExpectedResult []Comment
	ExpectedError  error
}

func getCommentHappyPath() getCommentScenario {
	expectedResult := []Comment{
		Comment{
			ID:     1,
			Body:   "ABC",
			PostID: 1,
		},
		Comment{
			ID:     2,
			Body:   "XYZ",
			PostID: 2,
		},
	}
	return getCommentScenario{
		Desc: "Get comment happy path",
		DoMock: func(mock *httpclient.MockHTTPClient) {
			listMocks := expectedResult
			bytes, _ := json.Marshal(listMocks)
			respMocks := &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
			}
			mock.On("Get", getCommentsEndpoint).Return(respMocks, nil)
		},
		ExpectedResult: expectedResult,
		ExpectedError:  nil,
	}
}

func getCommentFailed() getCommentScenario {
	expectedError := errors.New("Failed to get comments")
	return getCommentScenario{
		Desc: "Get comment failed",
		DoMock: func(mock *httpclient.MockHTTPClient) {
			mock.On("Get", getCommentsEndpoint).Return(nil, expectedError)
		},
		ExpectedResult: nil,
		ExpectedError:  expectedError,
	}
}

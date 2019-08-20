package getapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	httpclient "thien.com/http-client"
)

type testCaseGetPost struct {
	desc          string
	doMock        func(httpClientMocks *httpclient.MockHTTPClient)
	expectedPosts []Post
	expectedError error
}

type testCaseGetComment struct {
	desc             string
	doMock           func(httpClientMocks *httpclient.MockHTTPClient)
	expectedComments []Comment
	expectedError    error
}

func getScenarioGetPost() []testCaseGetPost {
	var expected = []Post{
		Post{ID: 1, Title: "Post 1"},
	}
	return []testCaseGetPost{
		testCaseGetPost{
			desc: "success",
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				bytes, _ := json.Marshal(expected)
				respMocks := &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
				}
				httpClientMocks.On("Get", getPostsEndpoint).Return(respMocks, nil)

			},
			expectedPosts: expected,
			expectedError: nil,
		},
		{
			desc:          "fail",
			expectedPosts: nil,
			expectedError: errors.New("failed http"),
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				httpClientMocks.On("Get", getPostsEndpoint).Return(nil, errors.New("failed http"))
			},
		},
	}
}

func getScenarioGetComment() []testCaseGetComment {
	var expected = []Comment{
		{ID: 1, Body: "Comment 1", PostID: 1},
	}
	return []testCaseGetComment{
		{
			desc: "success",
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				bytes, _ := json.Marshal(expected)
				respMocks := &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
				}
				httpClientMocks.On("Get", getCommentsEndpoint).Return(respMocks, nil)

			},
			expectedComments: expected,
			expectedError:    nil,
		},
		{
			desc:             "fail",
			expectedComments: nil,
			expectedError:    errors.New("failed http"),
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				httpClientMocks.On("Get", getCommentsEndpoint).Return(nil, errors.New("failed http"))
			},
		},
	}
}

func TestGetPost(t *testing.T) {
	testCases := getScenarioGetPost()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			httpClientMocks := &httpclient.MockHTTPClient{}
			if tc.doMock != nil {
				tc.doMock(httpClientMocks)
			}

			api, _ := New(httpClientMocks)
			res, err := api.GetPosts()
			assert.Equal(t, res, tc.expectedPosts)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestGetComment(t *testing.T) {
	testCases := getScenarioGetComment()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			httpClientMocks := &httpclient.MockHTTPClient{}
			if tc.doMock != nil {
				tc.doMock(httpClientMocks)
			}

			api, _ := New(httpClientMocks)
			res, err := api.GetComments()
			assert.Equal(t, res, tc.expectedComments)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

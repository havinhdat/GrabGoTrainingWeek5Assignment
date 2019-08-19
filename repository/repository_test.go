package repository

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"bathien/assignment/entity"
	httpclient "bathien/assignment/http-client"
)

func TestGetPosts(t *testing.T) {
	var (
		expectedPosts = []entity.Post{
			{
				ID:    1,
				Title: "POST 1",
			},
			{
				ID:    2,
				Title: "POST 2",
			},
		}
		errUnmarshal = json.Unmarshal([]byte("<"), &entity.Post{})
	)
	testCases := []struct {
		desc          string
		doMock        func(httpClientMocks *httpclient.MockHTTPClient)
		expectedPosts []entity.Post
		expectedError error
	}{
		{
			desc:          "success",
			expectedPosts: expectedPosts,
			expectedError: nil,
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				listMocks := expectedPosts
				bytes, _ := json.Marshal(listMocks)
				respMocks := &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
				}
				httpClientMocks.On("Get", getPostsEndpoint).Return(respMocks, nil)
			},
		},
		{
			desc:          "failed",
			expectedPosts: nil,
			expectedError: errors.New("anything"),
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				httpClientMocks.On("Get", getPostsEndpoint).Return(nil, errors.New("anything"))
			},
		},
		{
			desc:          "failed unmarshall",
			expectedPosts: nil,
			expectedError: errUnmarshal,
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				listMocks := expectedPosts
				bytes, _ := xml.Marshal(listMocks)
				respMocks := &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
				}
				httpClientMocks.On("Get", getPostsEndpoint).Return(respMocks, nil)
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			httpClientMocks := &httpclient.MockHTTPClient{}
			conf, _ := New(WithHTTPClient(httpClientMocks))

			if tc.doMock != nil {
				tc.doMock(httpClientMocks)
			}

			posts, err := conf.GetPosts()

			assert.Equal(t, tc.expectedPosts, posts)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetComments(t *testing.T) {
	var (
		expectedComments = []entity.Comment{
			{
				ID:     1,
				Body:   "Comment 1",
				PostID: 1,
			},
			{
				ID:     2,
				Body:   "Comment 2",
				PostID: 1,
			},
			{
				ID:     3,
				Body:   "Comment a",
				PostID: 2,
			},
		}
		errUnmarshal = json.Unmarshal([]byte("<"), &entity.Comment{})
	)
	testCases := []struct {
		desc             string
		doMock           func(httpClientMocks *httpclient.MockHTTPClient)
		expectedComments []entity.Comment
		expectedError    error
	}{
		{
			desc:             "success",
			expectedComments: expectedComments,
			expectedError:    nil,
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				listMocks := expectedComments
				bytes, _ := json.Marshal(listMocks)
				respMocks := &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
				}
				httpClientMocks.On("Get", getCommentsEndpoint).Return(respMocks, nil)
			},
		},
		{
			desc:             "failed unmarshall",
			expectedComments: nil,
			expectedError:    errUnmarshal,
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				listMocks := expectedComments
				bytes, _ := xml.Marshal(listMocks)
				respMocks := &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(string(bytes))),
				}
				httpClientMocks.On("Get", getCommentsEndpoint).Return(respMocks, nil)
			},
		},
		{
			desc:             "failed",
			expectedComments: nil,
			expectedError:    errors.New("anything"),
			doMock: func(httpClientMocks *httpclient.MockHTTPClient) {
				httpClientMocks.On("Get", getCommentsEndpoint).Return(nil, errors.New("anything"))
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			httpClientMocks := &httpclient.MockHTTPClient{}
			conf, _ := New(WithHTTPClient(httpClientMocks))

			if tc.doMock != nil {
				tc.doMock(httpClientMocks)
			}

			comments, err := conf.GetComments()

			assert.Equal(t, tc.expectedComments, comments)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestMissingHttpClient(t *testing.T) {
	conf, err := New(WithHTTPClient(nil))

	assert.Equal(t, nil, conf)
	assert.Equal(t, "missing http client", err.Error())
}

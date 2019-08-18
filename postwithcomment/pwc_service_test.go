package postwithcomment

import (
	"errors"
	"grab/week5/GrabGoTrainingWeek5Assignment/comment"
	"grab/week5/GrabGoTrainingWeek5Assignment/post"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPostWithComments(t *testing.T) {
	testcases := []getPostWithCommentsScenario{
		getPostWithCommentsHappyPath(),
		getPostWithCommentsFailedGetPosts(),
		getPostWithCommentsFailedGetComments(),
	}
	for _, testCase := range testcases {
		tc := testCase
		t.Run(tc.Desc, func(t *testing.T) {
			commentMock := &comment.MockCommentInterface{}
			postMock := &post.MockPostInterface{}
			pwcService := NewPostWithCommentService(postMock, commentMock)

			if tc.DoMock != nil {
				tc.DoMock(commentMock, postMock)
			}

			resp, err := pwcService.GetPostWithComments()

			assert.Equal(t, tc.ExpectedResult, resp)
			assert.Equal(t, tc.ExpectedError, err)
		})
	}
}

type getPostWithCommentsScenario struct {
	Desc           string
	DoMock         func(commentService *comment.MockCommentInterface, postService *post.MockPostInterface)
	ExpectedResult []PostWithComment
	ExpectedError  error
}

func getPostWithCommentsHappyPath() getPostWithCommentsScenario {
	comments := []comment.Comment{
		comment.Comment{
			ID:     1,
			Body:   "Comment 1",
			PostID: 1,
		},
		comment.Comment{
			ID:     2,
			Body:   "Comment 2",
			PostID: 1,
		},
	}
	expectedResut := []PostWithComment{
		PostWithComment{
			Comments: comments,
			ID:       1,
			Title:    "Post 1",
		},
		PostWithComment{
			Comments: nil,
			ID:       2,
			Title:    "Post 2",
		},
	}
	return getPostWithCommentsScenario{
		Desc: "Get post with comments happy path",
		DoMock: func(commentService *comment.MockCommentInterface, postService *post.MockPostInterface) {
			commentService.On("GetComments").Return(comments, nil)
			posts := []post.Post{
				post.Post{
					ID:    1,
					Title: "Post 1",
				},
				post.Post{
					ID:    2,
					Title: "Post 2",
				},
			}
			postService.On("GetPosts").Return(posts, nil)
		},
		ExpectedResult: expectedResut,
		ExpectedError:  nil,
	}
}

func getPostWithCommentsFailedGetPosts() getPostWithCommentsScenario {
	comments := []comment.Comment{
		comment.Comment{
			ID:     1,
			Body:   "Comment 1",
			PostID: 1,
		},
		comment.Comment{
			ID:     2,
			Body:   "Comment 2",
			PostID: 1,
		},
	}
	expectedError := errors.New("Failed to get posts")
	return getPostWithCommentsScenario{
		Desc: "Get post with comments failed get posts",
		DoMock: func(commentService *comment.MockCommentInterface, postService *post.MockPostInterface) {
			commentService.On("GetComments").Return(comments, nil)
			postService.On("GetPosts").Return(nil, expectedError)
		},
		ExpectedResult: nil,
		ExpectedError:  expectedError,
	}
}

func getPostWithCommentsFailedGetComments() getPostWithCommentsScenario {
	expectedError := errors.New("Failed to get comments")
	return getPostWithCommentsScenario{
		Desc: "Get post with comments failed get posts",
		DoMock: func(commentService *comment.MockCommentInterface, postService *post.MockPostInterface) {
			commentService.On("GetComments").Return(nil, expectedError)
			posts := []post.Post{
				post.Post{
					ID:    1,
					Title: "Post 1",
				},
				post.Post{
					ID:    2,
					Title: "Post 2",
				},
			}
			postService.On("GetPosts").Return(posts, nil)
		},
		ExpectedResult: nil,
		ExpectedError:  expectedError,
	}
}

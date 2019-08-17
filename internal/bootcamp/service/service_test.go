package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

func TestBlogServiceImpl_GetPostWithComments(t *testing.T) {
	bs := NewBlogServiceImpl(getPostsEndpoint, getCommentsEndpoint)

	comments, err := bs.GetPostWithComments()

	assert.NoError(t, err)
	assert.NotEmpty(t, comments)

}

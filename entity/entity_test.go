package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinePostWithComments(t *testing.T) {
	posts := []Post{
		{
			ID:    1,
			Title: "POST 1",
		},
		{
			ID:    2,
			Title: "POST 2",
		},
	}
	comments := []Comment{
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
	}
	expectedResut := []PostWithComments{
		{
			Comments: comments,
			ID:       1,
			Title:    "POST 1",
		},
		{
			Comments: nil,
			ID:       2,
			Title:    "POST 2",
		},
	}

	postWithComments := CombinePostWithComments(posts, comments)
	assert.Equal(t, expectedResut, postWithComments)

}

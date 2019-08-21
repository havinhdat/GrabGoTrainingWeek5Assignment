package interfaces

import (
	"net/http"

	"../defination"
)

type LoadData interface {
	GetPosts() ([]defination.Post, error)
	GetComments() ([]defination.Comment, error)
}

type CombineData interface {
	CombinePostWithComments(posts []defination.Post, comments []defination.Comment) []defination.PostWithComments
}

type RenderData interface {
	Render(writer http.ResponseWriter, postWithComment []defination.PostWithComments, typeOutput string) ([]byte, error)
}

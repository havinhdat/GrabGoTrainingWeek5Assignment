package bootcamp

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type BlogEndpoint interface {
	GetPostsWithComments(writer http.ResponseWriter, request *http.Request)
}

type BlogEndpointImpl struct {
	service BlogService
}

func NewBlogEndpointImpl() *BlogEndpointImpl {
	blog := NewBlogImpl(getPostsEndpoint, getCommentsEndpoint)
	return &BlogEndpointImpl{blog}
}

func (b *BlogEndpointImpl) GetPostsWithComments(writer http.ResponseWriter, request *http.Request) {
	posts, err := b.service.GetPostWithComments()
	if err == nil {
		buf, err := json.Marshal(posts)
		if err == nil {
			writer.Header().Set("Content-Type", "application/json")
			_, err = writer.Write(buf)

		} else {
			log.Println("unable to parse response: ", err)
			writer.WriteHeader(500)
		}

	} else {
		log.Println("get comments or posts failed with error: ", err)
		writer.WriteHeader(500)
	}
}

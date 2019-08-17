package bootcamp

import (
	"log"
	"net/http"
)

type BlogEndpoint interface {
	GetPostsWithComments(writer http.ResponseWriter, request *http.Request)
}

type blogEndpoint struct {
	service BlogService
	encode  EncodeResponseFunc
}

func newBlogEndpoint(service BlogService, encode EncodeResponseFunc) *blogEndpoint {
	return &blogEndpoint{service: service, encode: encode}
}

func (b *blogEndpoint) GetPostsWithComments(writer http.ResponseWriter, request *http.Request) {
	posts, err := b.service.GetPostWithComments()
	if err == nil {
		err := b.encode(writer, posts)
		if err != nil {
			log.Println("unable to parse response: ", err)
			writer.WriteHeader(500)
		}
	} else {
		log.Println("get comments or posts failed with error: ", err)
		writer.WriteHeader(500)
	}
}

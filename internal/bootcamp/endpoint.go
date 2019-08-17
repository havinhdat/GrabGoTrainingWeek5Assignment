package bootcamp

import (
	"encoding/json"
	"log"
	"net/http"
)

type BlogEndpoint interface {
	GetPostsWithComments(writer http.ResponseWriter, request *http.Request)
}

type blogEndpoint struct {
	service BlogService
}

func newBlogEndpoint(service BlogService) *blogEndpoint {
	return &blogEndpoint{service: service}
}

func (b *blogEndpoint) GetPostsWithComments(writer http.ResponseWriter, request *http.Request) {
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

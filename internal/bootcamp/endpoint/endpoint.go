package endpoint

import (
	"grab/internal/bootcamp/service"
	"grab/internal/bootcamp/transport"
	"log"
	"net/http"
)

type BlogEndpoint interface {
	GetPostsWithComments(writer http.ResponseWriter, request *http.Request)
}

type BlogEndpointImpl struct {
	service service.BlogService
	encoder transport.EncodeResponse
}

func NewBlogEndpointImpl(service service.BlogService, encoder transport.EncodeResponse) *BlogEndpointImpl {
	return &BlogEndpointImpl{service: service, encoder: encoder}
}

func (b *BlogEndpointImpl) GetPostsWithComments(writer http.ResponseWriter, request *http.Request) {
	posts, err := b.service.GetPostWithComments()
	if err == nil {
		err := b.encoder.Encode(writer, posts)
		if err != nil {
			log.Println("unable to parse response: ", err)
			writer.WriteHeader(500)
		}
	} else {
		log.Println("get comments or posts failed with error: ", err)
		writer.WriteHeader(500)
	}
}

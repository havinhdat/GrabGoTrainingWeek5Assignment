package bootcamp

import (
	"grab/internal/bootcamp/endpoint"
	"grab/internal/bootcamp/service"
	"grab/internal/bootcamp/transport"
	"log"
	"net/http"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

func StartServer() {
	blogService := service.NewBlogServiceImpl(getPostsEndpoint, getCommentsEndpoint)
	jsonEndpoint := endpoint.NewBlogEndpointImpl(blogService, transport.NewJsonResponseEncoder())

	http.HandleFunc("/postWithComments", jsonEndpoint.GetPostsWithComments)

	xmlEndpoint := endpoint.NewBlogEndpointImpl(blogService, transport.NewXmlResponseEncoder())
	http.HandleFunc("/postWithCommentsInXml", xmlEndpoint.GetPostsWithComments)

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

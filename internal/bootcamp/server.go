package bootcamp

import (
	"grab/internal/bootcamp/encode"
	"grab/internal/bootcamp/endpoint"
	"grab/internal/bootcamp/service"
	"log"
	"net/http"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

func StartServer() {
	blogService := service.NewBlogServiceImpl(getPostsEndpoint, getCommentsEndpoint)
	jsonEndpoint := endpoint.NewBlogEndpointImpl(blogService, encode.NewJsonResponseEncoder())

	http.HandleFunc("/postWithComments", jsonEndpoint.GetPostsWithComments)

	xmlEndpoint := endpoint.NewBlogEndpointImpl(blogService, encode.NewXmlResponseEncoder())
	http.HandleFunc("/postWithCommentsInXml", xmlEndpoint.GetPostsWithComments)

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package bootcamp

import (
	"log"
	"net/http"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

func StartServer() {
	service := newBlogService(getPostsEndpoint, getCommentsEndpoint)
	endpoint := newBlogEndpoint(service, jsonResponseEncoder)

	http.HandleFunc("/postWithComments", endpoint.GetPostsWithComments)

	xmlEndpoint := newBlogEndpoint(service, xmlResponseEncoder)
	http.HandleFunc("/postWithCommentsInXml", xmlEndpoint.GetPostsWithComments)

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

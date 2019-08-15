package main

import (
	"grab/internal/bootcamp"
	"log"
	"net/http"
)

func main() {
	endpoint := bootcamp.NewBlogEndpointImpl()
	http.HandleFunc("/postWithComments", endpoint.GetPostsWithComments)

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

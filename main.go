package main

import (
	"log"
	"net/http"

	"bathien/assignment/entity"
	"bathien/assignment/renderer"
	"bathien/assignment/repository"
)

func main() {
	httpClient := http.DefaultClient
	xmlRenderer := &renderer.JSONRenderer{}
	httpRepository, httpRepositoryErr := repository.New(repository.WithHTTPClient(httpClient))

	if httpRepositoryErr != nil {
		log.Println("ERROR: failed to init api service")
		return
	}

	http.HandleFunc("/postWithComments", func(writer http.ResponseWriter, request *http.Request) {
		// Get posts from config
		posts, err := httpRepository.GetPosts()
		if err != nil {
			log.Println("get posts failed with error: ", err)
			writer.WriteHeader(500)
			return
		}

		// Get comments from config
		comments, err := httpRepository.GetComments()
		if err != nil {
			log.Println("get comments failed with error: ", err)
			writer.WriteHeader(500)
			return
		}

		// Combine and return response
		postWithComments := entity.CombinePostWithComments(posts, comments)
		resp := entity.PostWithCommentsResponse{Posts: postWithComments}
		xmlRenderer.Render(writer, resp)
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"

	handler "GrabBootCamp2019/GrabGoTrainingWeek5Assignment/handler"
)

//TODO: how to separate API logic, business logic and response format logic
func main() {
	http.Handle("/postWithComments", handler.NewPostCommentHandler())

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

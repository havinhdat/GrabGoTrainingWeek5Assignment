package main

import (
	"log"
	"net/http"

	"dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/api"
	comment_getter "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/comment"
	encoder "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/encode"
	post_getter "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/post"
	post_comment_combiner "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/post-comment"
)

//TODO: how to separate API logic, business logic and response format logic
func main() {
	httpClient := http.DefaultClient
	postGetter, err := post_getter.New(httpClient)
	if err != nil {
		log.Println("ERROR: failed to init getter")
		return
	}
	commentGetter, err := comment_getter.New(httpClient)
	if err != nil {
		log.Println("ERROR: failed to init getter")
		return
	}

	jsonEncoder := encoder.NewJsonEncoderResponse()
	postCommentCombiner := post_comment_combiner.NewPostWithCommentsCombinerImpl()

	//xmlEncoder := encoder.NewXmlEncoderResponse()

	apiIml, _ := api.New(postGetter, commentGetter, postCommentCombiner, jsonEncoder)

	http.HandleFunc("/postWithComments", apiIml.GetPostsWithComments)

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

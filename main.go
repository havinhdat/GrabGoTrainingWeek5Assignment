package main

import (
	"grab/week5/GrabGoTrainingWeek5Assignment/comment"
	"grab/week5/GrabGoTrainingWeek5Assignment/httpservice"
	"grab/week5/GrabGoTrainingWeek5Assignment/post"
	pwc "grab/week5/GrabGoTrainingWeek5Assignment/postwithcomment"
	"grab/week5/GrabGoTrainingWeek5Assignment/renderer"
	"net/http"
)

func main() {
	httpClient := http.DefaultClient
	postSerice := post.NewPostService(httpClient)
	commentService := comment.NewCommentService(httpClient)
	pwcService := pwc.NewPostWithCommentService(postSerice, commentService)
	renderService := &renderer.XmlRender{}
	httpService := httpservice.NewPostWithCommentHttpService(pwcService, renderService)
	httpService.StartServer()
}

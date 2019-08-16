package main

import (
	"grab/week5/GrabGoTrainingWeek5Assignment/comment"
	"grab/week5/GrabGoTrainingWeek5Assignment/httpservice"
	"grab/week5/GrabGoTrainingWeek5Assignment/post"
	pwc "grab/week5/GrabGoTrainingWeek5Assignment/postwithcomment"
	"grab/week5/GrabGoTrainingWeek5Assignment/renderer"
)

func main() {
	postSerice := &post.PostService{}
	commentService := &comment.CommentService{}
	pwcService := pwc.NewPostWithCommentService(postSerice, commentService)
	renderService := &renderer.JsonRender{}
	httpService := httpservice.NewPostWithCommentHttpService(pwcService, renderService)
	httpService.StartServer()
}

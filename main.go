package main

import (
	"grab/week5/GrabGoTrainingWeek5Assignment/comment"
	"grab/week5/GrabGoTrainingWeek5Assignment/httpservice"
	"grab/week5/GrabGoTrainingWeek5Assignment/post"
	pwc "grab/week5/GrabGoTrainingWeek5Assignment/postwithcomment"
	"grab/week5/GrabGoTrainingWeek5Assignment/renderer"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

func main() {
	postSerice := post.NewPostService(getPostsEndpoint)
	commentService := comment.NewCommentService(getCommentsEndpoint)
	pwcService := pwc.NewPostWithCommentService(postSerice, commentService)
	renderService := &renderer.JsonRender{}
	httpService := httpservice.NewPostWithCommentHttpService(pwcService, renderService)
	httpService.StartServer()
}

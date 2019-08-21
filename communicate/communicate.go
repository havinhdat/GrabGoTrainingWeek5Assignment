package communicate

import (
	"net/http"

	"../dataloader"
	"../defination"
	"../handleerror"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

func HandleFunc(writer http.ResponseWriter, request *http.Request, isJSON bool) {
	process := new(dataloader.SeviceLoader)
	combineData := new(dataloader.CombineSeviceLoader)
	render := new(dataloader.RenderServiceLoader)
	// Get posts from api
	posts, err := process.GetPosts()
	handleerror.HandleError(writer, err, "get posts failed with error")

	// Get comments from api
	comments, err := process.GetComments()
	handleerror.HandleError(writer, err, "get comments failed with error")

	// Combine and return response
	postWithComments := combineData.CombinePostWithComments(posts, comments)
	handleerror.HandleError(writer, err, "unable to parse response")

	resp := defination.PostWithCommentsResponse{Posts: postWithComments}
	render.Render(writer, resp, isJSON)
}

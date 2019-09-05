package main

import (
	"log"
	"net/http"

	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/usecase"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/data"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/data/dataformatter"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/formatter"
)

func newGetPostsWithCommentsUsecase() usecase.GetPostsWithCommentsUsecase {
	dataFormatter := &dataformatter.JSONDataFormatter{}
	postData := data.NewPostData(dataFormatter)
	getPostsUsecase := usecase.NewGetPostsUsecase(postData)
	commentData := data.NewCommentData(dataFormatter)
	getCommentsUsecase := usecase.NewGetCommentsUsecase(commentData)

	return usecase.NewGetPostsWithCommentsUsecase(getPostsUsecase, getCommentsUsecase)
}

func getPostsWithCommentsHandler(writer http.ResponseWriter, request *http.Request) {
	// define my mapper
	myFormatter := &formatter.JSONFormatter{
		ContentType: "application/json",
	}
	// define usecase
	getPostsWithCommentsUsecase := newGetPostsWithCommentsUsecase()

	// executing usecase
	ret, err := getPostsWithCommentsUsecase.GetPostsWithComments()
	if err != nil {
		log.Println("execute usecase failed with error: ", err)
		writer.WriteHeader(500)
		return
	}

	// encoding data
	buf, err := myFormatter.Encode(ret)
	if err != nil {
		log.Println("unable to parse response: ", err)
		writer.WriteHeader(500)
	}

	// response to client
	writer.Header().Set("Content-Type", myFormatter.GetContentType())
	_, err = writer.Write(buf)
}

func main() {
	// define handler
	http.HandleFunc("/posts-with-comments", getPostsWithCommentsHandler)

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

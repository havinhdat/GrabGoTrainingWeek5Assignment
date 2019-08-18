package main

import (
	"log"
	"net/http"

	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/usecase"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/data"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/data/dataformatter"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/formatter"
)

type requestHandler struct {
	endPoint  string
	useCase   usecase.Usecase
	formatter formatter.Formatter
}

func handleRequest(rq *requestHandler) {
	http.HandleFunc(rq.endPoint, func(writer http.ResponseWriter, request *http.Request) {
		// executing usecase
		ret, err := rq.useCase.Execute()
		if err != nil {
			log.Println("execute usecase failed with error: ", err)
			writer.WriteHeader(500)
			return
		}

		// encoding data
		buf, err := rq.formatter.Encode(ret)
		if err != nil {
			log.Println("unable to parse response: ", err)
			writer.WriteHeader(500)
		}

		// response to client
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(buf)
	})
}

func newGetPostsWithCommentsUsecase() *usecase.GetPostsWithCommentsUsecase {
	dataFormatter := &dataformatter.JSONDataFormatter{}
	postData := data.NewPostData(dataFormatter)
	getPostsUsecase := usecase.NewGetPostsUsecase(postData)
	commentData := data.NewCommentData(dataFormatter)
	getCommentsUsecase := usecase.NewGetCommentsUsecase(commentData)

	return usecase.NewGetPostsWithCommentsUsecase(getPostsUsecase, getCommentsUsecase)
}

func main() {
	// define my mapper
	myFormatter := &formatter.JSONFormatter{}
	// define usecase
	getPostsWithCommentsUsecase := newGetPostsWithCommentsUsecase()
	// define handler
	handleRequest(&requestHandler{
		endPoint:  "/posts-with-comments",
		useCase:   getPostsWithCommentsUsecase,
		formatter: myFormatter,
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

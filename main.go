package main

import (
	"log"
	"net/http"

	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/usecase"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/data"
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/mapper"
)

type requestHandler struct {
	endPoint string
	useCase  usecase.Usecase
	mapper   mapper.Mapper
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
		buf, err := rq.mapper.Encode(ret)
		if err != nil {
			log.Println("unable to parse response: ", err)
			writer.WriteHeader(500)
		}

		// response to client
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(buf)
	})
}

func main() {
	// define mapper
	// notice: when use XMLMapper, GetPosts and GetComments have to return xml format
	// currently return json
	mapper := &mapper.JSONMapper{}
	// define usecase
	getPostsWithCommentsUsecase := &usecase.GetPostsWithCommentsUsecase{
		GetPosts: usecase.GetPostsUsecase{
			Repo: &data.PostData{
				Mapper: mapper,
			},
		},
		GetComments: usecase.GetCommentsUsecase{
			Repo: &data.CommentData{
				Mapper: mapper,
			},
		},
	}
	// define handler
	handleRequest(&requestHandler{
		endPoint: "/posts-with-comments",
		useCase:  getPostsWithCommentsUsecase,
		mapper:   mapper,
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package errorhandler

import (
	"log"
	"net/http"
)

type ErrorHandler interface {
	Handle(err error, writer http.ResponseWriter)
}

type CommentErrorHandler struct{}

type PostErrorHandler struct{}

type PostWithCommentsErrorHandler struct{}

func (e *PostErrorHandler) Handle(err error, writer http.ResponseWriter) {
	if err != nil {
		log.Println("get posts failed with error: ", err)
		writer.WriteHeader(500)
		return
	}
}

func (e *CommentErrorHandler) Handle(err error, writer http.ResponseWriter) {
	if err != nil {
		log.Println("get comments failed with error: ", err)
		writer.WriteHeader(500)
		return
	}
}

func (e *PostWithCommentsErrorHandler) Handle(err error, writer http.ResponseWriter) {
	if err != nil {
		log.Println("unable to parse response: ", err)
		writer.WriteHeader(500)
	}
}

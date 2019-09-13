package handler

import (
	"GrabBootCamp2019/GrabGoTrainingWeek5Assignment/business"
	"GrabBootCamp2019/GrabGoTrainingWeek5Assignment/model"
	"GrabBootCamp2019/GrabGoTrainingWeek5Assignment/render"
	_ "encoding/json"
	"fmt"
	_ "log"
	"net/http"
)

type BusinessService interface {
	CombineData() (model.PostWithCommentsResponse, error)
}

type Renderer interface {
	Serialize(postsWithComments model.PostWithCommentsResponse) ([]byte, string, error)
}

//PostCommentHandler handles requests from /postsWithComments
type PostCommentHandler struct {
	Business BusinessService
	Renderer Renderer
}

func NewPostCommentHandler() *PostCommentHandler {
	return &PostCommentHandler{
		Business: new(business.BusinessServiceImpl),
		Renderer: new(render.XmlRenderer),
	}
}

func (h *PostCommentHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {

	case http.MethodGet:
		h.Get(writer, request)
	}
}

//Get handles requests with method GET
func (h *PostCommentHandler) Get(writer http.ResponseWriter, request *http.Request) {
	postsWithComments, err := h.Business.CombineData()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	body, contentType, err := h.Renderer.Serialize(postsWithComments)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	fmt.Println(contentType)
	writer.Header().Set("Content-Type", contentType)
	writer.Write(body)
	writer.WriteHeader(http.StatusOK)
	return
}

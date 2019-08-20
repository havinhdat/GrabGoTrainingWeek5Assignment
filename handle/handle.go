package handle

import (
	"errors"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/mimeGenerate"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/service"
	"net/http"
)

//go:generate mockery -name=Handle -inpkg
type Handler interface {
	Get(writer http.ResponseWriter, request *http.Request)
}

type hanlderImpl struct {
	mimeGenerate mimeGenerate.MimeGenerator
	service      service.Service
}

func (h *hanlderImpl) Get(writer http.ResponseWriter, request *http.Request) {
	postWithComments, err := h.service.GetPostWithComments()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := h.mimeGenerate.Generate(postWithComments)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(body.Body)
	writer.WriteHeader(http.StatusOK)

}

func NewHandler(service service.Service, mimeGenerate mimeGenerate.MimeGenerator) (handle Handler, err error) {
	if service == nil || mimeGenerate == nil {
		return nil, errors.New("missing argument to get new handle")
	}
	return &hanlderImpl{service: service, mimeGenerate: mimeGenerate}, nil
}
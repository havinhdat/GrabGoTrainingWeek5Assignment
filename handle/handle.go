package handle

import (
	"errors"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/mimeGenerate"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/service"
	"net/http"
)

//go:generate mockery -name=Handle -inpkg
type Handle interface {
	Get(writer http.ResponseWriter, request *http.Request)
}

type hanldeImpl struct {
	mimeGenerate mimeGenerate.MimeGenerate
	service      service.Service
}

func (h hanldeImpl) Get(writer http.ResponseWriter, request *http.Request) {
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

func NewHandle(service service.Service, mimeGenerate mimeGenerate.MimeGenerate) (handle Handle, err error) {
	if service == nil || mimeGenerate == nil {
		return nil, errors.New("missing argument to get new hanle")
	}
	return &hanldeImpl{service: service, mimeGenerate: mimeGenerate}, nil
}
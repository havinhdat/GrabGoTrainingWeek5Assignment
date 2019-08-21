package dataloader

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"

	"../defination"
	"../interfaces"
)

type RenderServiceLoader struct {
	data interfaces.RenderData
}

func (service *RenderServiceLoader) Render(writer http.ResponseWriter, postWithComment defination.PostWithCommentsResponse, isJSON bool) {

	var buf []byte
	var err error
	var contentType string

	if isJSON == true {
		buf, err = json.Marshal(postWithComment)
		contentType = "application/json"
	} else {
		buf, err = xml.Marshal(postWithComment)
		contentType = "application/xml"
	}

	if err != nil {
		log.Println("unable to parse response: ", err)
		writer.WriteHeader(500)
	}

	writer.Header().Set("Content-Type", contentType)
	_, err = writer.Write(buf)
}

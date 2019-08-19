package main

import (
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/mimeGenerate"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/handle"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/service"
	"log"
	"net/http"
)

func main() {
	httpClient := http.DefaultClient
	service := service.NewService(httpClient)

	generateXml := mimeGenerate.NewXmlMimeGenerate()
	generateJson := mimeGenerate.NewJsonMimeGenerate()

	handleXml, err := handle.NewHandle(service, generateXml)
	handleJson, err := handle.NewHandle(service, generateJson)
	if err != nil {
		log.Fatal("could not init components to start program")
	}
	http.HandleFunc("/posts", func(writer http.ResponseWriter, request *http.Request) {
		contentType := request.Header.Get("Content-type")
		if contentType == generateXml.GetGenerateType() {
			handleXml.Get(writer, request)
			return
		}
		handleJson.Get(writer, request)
		return
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

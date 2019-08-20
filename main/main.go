package main

import (
	"fmt"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/handle"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/mimeGenerate"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/service"
	"log"
	"net/http"
)

func main() {
	httpClient := http.DefaultClient
	service := service.NewService(httpClient)

	xmlGenerator := mimeGenerate.NewXmlMimeGenerator()
	jsonGenerator := mimeGenerate.NewJsonMimeGenerator()

	handleXml, err := handle.NewHandler(service, xmlGenerator)
	handleJson, err := handle.NewHandler(service, jsonGenerator)
	if err != nil {
		log.Fatal("could not init components to start program")
	}
	generators := map[string]handle.Handler{xmlGenerator.GetGenerateType(): handleXml, jsonGenerator.GetGenerateType(): handleJson}
	http.HandleFunc("/posts", func(writer http.ResponseWriter, request *http.Request) {
		contentType := request.Header.Get("Content-type")
		if _, ok := generators[contentType]; ok {
			generators[contentType].Get(writer, request)
			return
		}
		writer.Write([]byte(fmt.Sprintf("Content type %s not support",contentType)))
		return
	})

	log.Println("httpServer starts ListenAndServe at 8182")
	log.Fatal(http.ListenAndServe(":8182", nil))
}

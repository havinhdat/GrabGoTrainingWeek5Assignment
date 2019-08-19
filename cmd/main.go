package main

import (
	// "encoding/json"
	// "io/ioutil"
	"log"
	"net/http"
	"grab/internal/postcomments"
	"grab/service/jsonservice"
	"grab/service/xmlservice"
)

func main() {

	http.HandleFunc("/postWithComments", func(writer http.ResponseWriter, request *http.Request) {
		postsWithCommentsWorker := postcomments.CreatePostCommentsWorker()
		buf, error := postsWithCommentsWorker.GetPostWithComments()
		if error != nil {
			log.Println(error)
			return
		}
		accept := request.Header.Get("Accept")

		switch accept {
		case "application/xml" :
			writer.Header().Set("Content-Type", "application/xml")
			xmlWorker := xmlservice.CreateXMLWorker()
			newBuf, error := xmlWorker.Marshal(buf)
			if error != nil {
				log.Println(error)
				return
			}
			_, _ = writer.Write(newBuf)
		default:
			writer.Header().Set("Content-Type", "application/json")
			jsonWorker := jsonservice.CreateJSONWorker()
			newBuf, error := jsonWorker.Marshal(buf)
			if error != nil {
				log.Println(error)
				return
			}
			_, _ = writer.Write(newBuf)
		}

		
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

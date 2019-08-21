package handleerror

import (
	"log"
	"net/http"
)

func HandleError(writer http.ResponseWriter, err error, errorString string) {
	if err != nil {
		log.Printf("%v : %v", errorString, err)
		writer.WriteHeader(500)
		return
	}
}

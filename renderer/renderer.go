package renderer

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

type Renderer interface {
	Render(w http.ResponseWriter, src interface{})
}

type JSONRenderer struct{}
type XMLRenderer struct{}

func (de *JSONRenderer) Render(w http.ResponseWriter, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		log.Println("unable to parse response: ", err)
		http.Error(w, "Internal server error", 500)
		return err
	}
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
	return nil
}

func (de *XMLRenderer) Render(w http.ResponseWriter, src interface{}) error {
	b, err := xml.MarshalIndent(src, "", "  ")
	if err != nil {
		log.Println("unable to parse response: ", err)
		http.Error(w, err.Error(), 500)
		return err
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(b)
	return nil
}

package bootcamp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type EncodeResponseFunc func(http.ResponseWriter, interface{}) error

func jsonResponseEncoder(writer http.ResponseWriter, data interface{}) error {
	buf, err := json.Marshal(data)
	if err == nil {
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(buf)
		return nil
	} else {
		return err
	}
}

func xmlResponseEncoder(writer http.ResponseWriter, data interface{}) error {
	buf, err := xml.Marshal(data)
	if err == nil {
		writer.Header().Set("Content-Type", "application/xml")
		_, err = writer.Write(buf)
		return nil
	} else {
		return err
	}
}

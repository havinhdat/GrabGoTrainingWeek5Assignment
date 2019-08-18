package encode

import (
	"encoding/json"
	"net/http"
)

type EncodeResponse interface {
	Encode(http.ResponseWriter, interface{}) error
}

type JsonEncoderResponse struct{}

type XmlEncoderResponse struct{}

func (j *JsonEncoderResponse) Encode(writer http.ResponseWriter, resp interface{}) error {
	buf, err := json.Marshal(resp)
	if err == nil {
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(buf)
		return nil
	} else {
		return err
	}
}

func (j *XmlEncoderResponse) Encode(writer http.ResponseWriter, resp interface{}) error {
	buf, err := json.Marshal(resp)
	if err == nil {
		writer.Header().Set("Content-Type", "application/xml")
		_, err = writer.Write(buf)
		return nil
	} else {
		return err
	}
}

func NewJsonEncoderResponse() EncodeResponse {
	return &JsonEncoderResponse{}
}

func NewXmlEncoderResponse() EncodeResponse {
	return &XmlEncoderResponse{}
}

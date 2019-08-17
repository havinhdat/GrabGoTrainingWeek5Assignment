package encode

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type EncodeResponse interface {
	Encode(http.ResponseWriter, interface{}) error
}

type JsonResponseEncoder struct {
}

func NewJsonResponseEncoder() *JsonResponseEncoder {
	return &JsonResponseEncoder{}
}

func (j *JsonResponseEncoder) Encode(writer http.ResponseWriter, data interface{}) error {
	buf, err := json.Marshal(data)
	if err == nil {
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(buf)
		return nil
	} else {
		return err
	}
}

type XmlResponseEncoder struct {
}

func NewXmlResponseEncoder() *XmlResponseEncoder {
	return &XmlResponseEncoder{}
}

func (x *XmlResponseEncoder) Encode(writer http.ResponseWriter, data interface{}) error {
	buf, err := xml.Marshal(data)
	if err == nil {
		writer.Header().Set("Content-Type", "application/xml")
		_, err = writer.Write(buf)
		return nil
	} else {
		return err
	}
}

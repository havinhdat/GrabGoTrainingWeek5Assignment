package xmlservice

import (
	"encoding/xml"
)

type XMLWorker interface {
	Marshal(v interface{}) ([]byte, error)
}

type XMLWorkerImp struct {}

func CreateXMLWorker() *XMLWorkerImp{
	return &XMLWorkerImp{}
}

func (xmlwker *XMLWorkerImp) Marshal(v interface{}) ([]byte, error) {

	buf, error := xml.Marshal(v)
	if error != nil {
		return nil,error
	}

	return buf, nil
}

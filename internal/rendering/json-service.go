package jsonservice

import (
	"encoding/json"
)

type JSONWorker interface {
	Marshal(v interface{}) ([]byte, error)
}

type JSONWorkerImp struct {}

func CreateJSONWorker() *JSONWorkerImp{
	return &JSONWorkerImp{}
}

func (jsonwker *JSONWorkerImp) Marshal(v interface{}) ([]byte, error) {

	buf, error := json.Marshal(v)
	if error != nil {
		return nil,error
	}

	return buf, nil
}

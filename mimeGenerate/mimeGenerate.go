package mimeGenerate

import (
	"encoding/json"
	"encoding/xml"
)

const (
	Application_json ContentType = "application/json"
	Application_xml  ContentType = "application/xml"
)

func (e ContentType) GetAsString() string {
	switch e {
	case Application_xml:
		return "application/xml"
	default:
		return "application/json"
	}
}

//go:generate mockery -name=MimeGenerate -inpkg
type MimeGenerate interface {
	Generate(orgObj interface{}) (result MimeResult, err error)
	GetGenerateType() string
}

type MimeResult struct {
	Body        []byte
	ContentType ContentType
}

type ContentType string

type JsonMimeGenerate struct {
}

func NewJsonMimeGenerate() MimeGenerate {
	return &JsonMimeGenerate{}
}

func (j JsonMimeGenerate) Generate(orgObj interface{}) (result MimeResult, err error) {
	var buf []byte
	if buf, err = json.Marshal(orgObj); err != nil {
		return
	}
	result = MimeResult{Body: buf, ContentType: Application_json}
	return
}
func (j JsonMimeGenerate) GetGenerateType() string {
	return Application_json.GetAsString()
}

type XmlMimeGenerate struct {
}

func NewXmlMimeGenerate() MimeGenerate {
	return &XmlMimeGenerate{}
}

func (x XmlMimeGenerate) Generate(orgObj interface{}) (result MimeResult, err error) {
	var buf []byte
	if buf, err = xml.Marshal(orgObj); err != nil {
		return
	}
	result = MimeResult{Body: buf, ContentType: Application_xml}
	return
}
func (j XmlMimeGenerate) GetGenerateType() string {
	return Application_xml.GetAsString()
}

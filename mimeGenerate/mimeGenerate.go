package mimeGenerate

import (
	"encoding/json"
	"encoding/xml"
)

const (
	ApplicationJson ContentType = "application/json"
	ApplicationXml  ContentType = "application/xml"
)

func (e ContentType) GetAsString() string {
	switch e {
	case ApplicationXml:
		return "application/xml"
	default:
		return "application/json"
	}
}

//go:generate mockery -name=MimeGenerator -inpkg
type MimeGenerator interface {
	Generate(orgObj interface{}) (result MimeResult, err error)
	GetGenerateType() string
}

type MimeResult struct {
	Body        []byte
	ContentType ContentType
}

type ContentType string

type jsonMimeGenerator struct {
}

func NewJsonMimeGenerator() MimeGenerator {
	return &jsonMimeGenerator{}
}

func (j *jsonMimeGenerator) Generate(orgObj interface{}) (result MimeResult, err error) {
	var buf []byte
	if buf, err = json.Marshal(orgObj); err != nil {
		return
	}
	result = MimeResult{Body: buf, ContentType: ApplicationJson}
	return
}
func (j *jsonMimeGenerator) GetGenerateType() string {
	return ApplicationJson.GetAsString()
}

type xmlMimeGenerator struct {
}

func NewXmlMimeGenerator() MimeGenerator {
	return &xmlMimeGenerator{}
}

func (x *xmlMimeGenerator) Generate(orgObj interface{}) (result MimeResult, err error) {
	var buf []byte
	if buf, err = xml.Marshal(orgObj); err != nil {
		return
	}
	result = MimeResult{Body: buf, ContentType: ApplicationXml}
	return
}
func (j *xmlMimeGenerator) GetGenerateType() string {
	return ApplicationXml.GetAsString()
}

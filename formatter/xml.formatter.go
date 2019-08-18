package formatter

import (
	"encoding/xml"
)

// XMLFormatter data
type XMLFormatter struct {
	ContentType string
}

// Encode data
func (formatter *XMLFormatter) Encode(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

// Decode data
func (formatter *XMLFormatter) Decode(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// GetContentType data
func (formatter *XMLFormatter) GetContentType() string {
	return formatter.ContentType
}

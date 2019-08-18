package formatter

import (
	"encoding/xml"
)

// XMLFormatter data
type XMLFormatter struct {
}

// Encode data
func (formatter *XMLFormatter) Encode(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

// Decode data
func (formatter *XMLFormatter) Decode(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

package dataformatter

import (
	"encoding/xml"
)

// XMLDataFormatter data
type XMLDataFormatter struct {
}

// Encode data
func (formatter *XMLDataFormatter) Encode(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

// Decode data
func (formatter *XMLDataFormatter) Decode(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

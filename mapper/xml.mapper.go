package mapper

import (
	"encoding/xml"
)

// XMLMapper mapper
type XMLMapper struct {
}

// Encode data
func (mapper *XMLMapper) Encode(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

// Decode data
func (mapper *XMLMapper) Decode(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

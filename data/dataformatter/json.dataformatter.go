package dataformatter

import (
	"encoding/json"
)

// JSONDataFormatter data
type JSONDataFormatter struct {
}

// Encode data
func (formatter *JSONDataFormatter) Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// Decode data
func (formatter *JSONDataFormatter) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

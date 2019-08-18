package formatter

import (
	"encoding/json"
)

// JSONFormatter data
type JSONFormatter struct {
	ContentType string
}

// Encode data
func (formatter *JSONFormatter) Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// Decode data
func (formatter *JSONFormatter) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// GetContentType data
func (formatter *JSONFormatter) GetContentType() string {
	return formatter.ContentType
}

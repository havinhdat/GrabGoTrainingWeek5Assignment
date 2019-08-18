package mapper

import (
	"encoding/json"
)

// JSONMapper mapper
type JSONMapper struct {
}

// Encode data
func (mapper *JSONMapper) Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// Decode data
func (mapper *JSONMapper) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

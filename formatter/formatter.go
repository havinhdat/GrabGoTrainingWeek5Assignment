package formatter

// Formatter data
type Formatter interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}
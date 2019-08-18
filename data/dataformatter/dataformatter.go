package dataformatter

// DataFormatter data
type DataFormatter interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}
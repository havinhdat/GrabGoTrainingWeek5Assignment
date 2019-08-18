package mapper

type Mapper interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}
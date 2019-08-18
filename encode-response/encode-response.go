package encoderesponse

type EncodeResponse interface {
	Encode(v interface{}) error
}

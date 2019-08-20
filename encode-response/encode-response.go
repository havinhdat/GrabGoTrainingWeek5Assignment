package encoderesponse

//go:generate mockery -name=EncodeResponse -inpkg
type EncodeResponse interface {
	Encode(v interface{}) error
}

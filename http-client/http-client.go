package httpclient

import (
	"net/http"
)

//go:generate mockery -name=HTTPClient -inpkg
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

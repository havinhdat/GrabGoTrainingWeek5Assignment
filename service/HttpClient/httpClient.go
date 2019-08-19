package HttpClient

import "net/http"


//go:generate mockery -name=HttpClient -inpkg
type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

package httpworker

import (
	"net/http"
	"io/ioutil"
)
// HTTP Worker Interface provide Get method from input URL.
type HttpWorker interface {
	Get(url string) ([]byte, error)
}

// Imple ment HTTP worker
type HTTPWorkerImpgo struct {}

func CreateHTTPWorker() *HTTPWorkerImpgo{
	return &HTTPWorkerImpgo{}
}
// Get data from input url using GET Method.
func (httpworkerimp *HTTPWorkerImpgo) Get(url string) ([]byte, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return nil, err
	}

	return body, nil
}
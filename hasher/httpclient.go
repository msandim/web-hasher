package hasher

import "net/http"

// HTTPClient is an interface that represents an HTTP client.
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// DefaultHTTPClient represents an implementation of an HTTPClient that uses the "net/http" package.
type DefaultHTTPClient struct{}

// Get makes an HTTP GET request using the "net/http" package.
func (client DefaultHTTPClient) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

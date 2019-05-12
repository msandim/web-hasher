package hasher

import (
	"net/http"
	"net/http/httptest"
)

// Client is an interface that represents an HTTP client.
type Client interface {
	Get(url string) (resp *http.Response, err error)
}

// HTTPClient represents an implementation of an HTTPClient that uses the "net/http" package.
type HTTPClient struct{}

// Get makes an HTTP GET request using the "net/http" package.
func (client *HTTPClient) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

// mockClient represents an implementation of a Client that does not perform real HTTP requests.
type mockClient struct {
	fakeServer *httptest.Server
}

func newMockClient(handler func(url string) (body string)) *mockClient {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := handler(r.URL.String())
		w.Write([]byte(body))
	}))

	return &mockClient{fakeServer: fakeServer}
}

func (client *mockClient) Close() {
	client.fakeServer.Close()
}

func (client *mockClient) Get(url string) (resp *http.Response, err error) {
	return http.Get(client.fakeServer.URL)
}

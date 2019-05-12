package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	neturl "net/url"
)

// Hasher is an interface that represents the capability of hashing the contents present in a URL.
type Hasher interface {
	// Hash retrieves the md5 hash (as a string) of the contents of the url.
	Hash(url string) (string, error)
}

// HTTPHasher is an implementation of the interface Hasher, by doing an HTTP GET to the URL supplied
// and hashing the results with MD5.
type HTTPHasher struct {
	client Client
}

// NewHTTPHasher returns a new http-based hasher, using the client supplied.
func NewHTTPHasher(client Client) *HTTPHasher {
	return &HTTPHasher{client: client}
}

// Hash makes a http GET to url defined in the parameter and returns the MD5 hashed value of the response,
// or an error if that was not possible.
func (hasher *HTTPHasher) Hash(url string) (string, error) {

	urlParsed, err := neturl.Parse(url)

	if err != nil {
		return "", err
	}

	if urlParsed.Scheme == "" {
		urlParsed.Scheme = "http"
	}

	resp, err := hasher.getClient().Get(urlParsed.String())

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	hashArr := md5.Sum(body)
	hash := hex.EncodeToString(hashArr[:])
	return hash, nil
}

func (hasher *HTTPHasher) getClient() Client {
	return hasher.client
}

// MockHasher is an implementation of the interface Hasher, that has an internal map
// of urls to hashes.
type MockHasher struct {
	hashMap map[string]string
}

// NewMockHasher returns a MockHasher with a certain map of urls to hashes.
func NewMockHasher(hashMap map[string]string) *MockHasher {
	return &MockHasher{hashMap: hashMap}
}

// Hash returns the corresponding hash to the url parameter, or an url if that is not possible.
func (hasher *MockHasher) Hash(url string) (string, error) {

	hash, ok := hasher.hashMap[url]

	if !ok {
		return "", fmt.Errorf("URl is not present in the hashMap")
	}

	return hash, nil
}

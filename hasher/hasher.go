package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
)

// Hasher is an interface that represents the capability of hashing the contents present in a URL.
type Hasher interface {
	Hash(url string) (string, error)
}

// HTTPHasher is an implementation of the interface Hasher, by doing an HTTP GET to the URL supplied
// and hashing the results with MD5.
type HTTPHasher struct{}

// Hash makes a http GET to url defined in the parameter and returns the MD5
// hashed value of the response, or an error if that was not possible.
func (hasher HTTPHasher) Hash(url string) (string, error) {

	urlParsed, err := neturl.Parse(url)

	if err != nil {
		fmt.Println(err)
	}

	if urlParsed.Scheme == "" {
		urlParsed.Scheme = "http"
	}

	resp, err := http.Get(urlParsed.String())

	if err != nil {
		fmt.Println(err)
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

package hasher

import (
	"crypto/md5"
	"io/ioutil"
	"net/http"
)

// Hash makes a http GET to url defined in the parameter and returns the MD5
// hashed value of the response, or an error if that was not possible.
func Hash(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	hashArr := md5.Sum(body)
	hash := string(hashArr[:])
	return hash, nil
}

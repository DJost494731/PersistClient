package httpclient

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
}

//Get makes a get request at the given route and returns the response body or error
func (client HttpClient) Get(url string) (string, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}

	sb := string(body)
	return sb, resp.StatusCode, nil
}
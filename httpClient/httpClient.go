package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
}

//Get makes a get request at the given route and returns the response body or error
func (client HttpClient) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("expected 200 level error code. Got '%v'", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	sb := string(body)
	return sb, nil
}
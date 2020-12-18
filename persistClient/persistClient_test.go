package persistclient

import (
	"testing"

	"github.com/DJost494731/PersistClient/httpclient"
)


func Test_PersistClientReturnsHttpResponseBody(t *testing.T) {
	want := "foobar"
	httpClient := &httpclient.MockHttpClient{MockResponseBody: want}
	persistClient := PersistClient{HttpClient: httpClient}
	
	result := persistClient.GetData("stub folder", "stub file")
	

	if result != want { 
		t.Errorf("Expected '%v' got '%v'", want, result)
	}
}

func Test_PersistClientCallsHttpClientWithFormattedUrl(t *testing.T) {
	expectedUrl := "http://www.foo.com/stub folder/stub file"
	persistUrlPrefix := "http://www.foo.com"
	mockHttpClient := &httpclient.MockHttpClient{}
	persistClient := PersistClient{HttpClient: mockHttpClient, PersistUrlPrefix: persistUrlPrefix}
	
	persistClient.GetData("stub folder", "stub file")
	
	if !contains(mockHttpClient.GetUrlInvocations, expectedUrl) { 
		t.Errorf("Expected '%v' to contain '%v'", mockHttpClient.GetUrlInvocations, expectedUrl)
	}
}


//contains returns true when the array 'a' contains the element 'e'
func contains(a []string, e string) bool {
	for _, a := range a {
		if a == e {
				return true
		}
	}
	return false
}

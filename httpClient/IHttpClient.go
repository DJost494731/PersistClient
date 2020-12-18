package httpClient

type IHttpClient interface {
	Get(url string) string
}

type MockHttpClient struct {
	//The string that will be returned from a request.
	MockResponseBody string
	//'GET' request url invocations
	GetUrlInvocations []string
}

//Returns MockHttpClient's response body
func (client *MockHttpClient) Get(url string) string {
	client.initializeNilArrays()

	client.GetUrlInvocations = append(client.GetUrlInvocations, url)
	return client.MockResponseBody
}

func (client *MockHttpClient) initializeNilArrays() {
	if client.GetUrlInvocations == nil {
		client.GetUrlInvocations = []string{}
	}
}
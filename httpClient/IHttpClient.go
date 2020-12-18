package httpclient

type IHttpClient interface {
	Get(url string) (string, error)
}

type MockHttpClient struct {
	ErrortoReturn error
	//The string that will be returned from a request.
	MockResponseBody string
	//'GET' request url invocations
	GetUrlInvocations []string
}

//Returns MockHttpClient's response body
func (client *MockHttpClient) Get(url string) (string, error) {
	client.initializeNilArrays()

	client.GetUrlInvocations = append(client.GetUrlInvocations, url)
	return client.MockResponseBody, client.ErrortoReturn
}

func (client *MockHttpClient) initializeNilArrays() {
	if client.GetUrlInvocations == nil {
		client.GetUrlInvocations = []string{}
	}
}
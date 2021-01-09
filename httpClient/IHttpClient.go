package httpclient

type IHttpClient interface {
	Get(url string) (string, int, error)
}

type MockHttpClient struct {
	ErrortoReturn error
	//The string that will be returned from a request.
	MockResponseBody string
	//'GET' request url invocations
	GetUrlInvocations []string
	//Http Status code to return
	StatusCode int
}

//Returns MockHttpClient's response body
func (client *MockHttpClient) Get(url string) (string, int, error) {
	client.initializeNilArrays()

	client.GetUrlInvocations = append(client.GetUrlInvocations, url)
	return client.MockResponseBody, client.StatusCode, client.ErrortoReturn
}

func (client *MockHttpClient) initializeNilArrays() {
	if client.GetUrlInvocations == nil {
		client.GetUrlInvocations = []string{}
	}
}
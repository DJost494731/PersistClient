package persistClient

import "github.com/DJost494731/PersistClient/httpClient"


type PersistClient struct {
	HttpClient httpClient.IHttpClient
	//The url your prefix server is hosted at, without the folder or file path. Ex: "http://www.foo.com"
	PersistUrlPrefix string
}

func (p PersistClient) GetData(folder string, file string) string {
	return p.HttpClient.Get(p.PersistUrlPrefix + "/" + folder + "/" + file)
}
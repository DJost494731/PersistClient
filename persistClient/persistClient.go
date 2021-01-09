package persistclient

import "github.com/DJost494731/PersistClient/httpclient"


type PersistClient struct {
	HttpClient httpclient.IHttpClient
	//The url your prefix server is hosted at, without the folder or file path. Ex: "http://www.foo.com"
	PersistUrlPrefix string
}

//GetData gets data at the path from persist. Returns RecordNotFoundError if no data is found.
func (p PersistClient) GetData(folder string, file string) (string, error) {
	response, statusCode, err  :=  p.HttpClient.Get(p.PersistUrlPrefix + "/" + folder + "/" + file)
	
	if statusCode == 404 {
		return "", &RecordNotFoundError{Path: folder + "/" + file} 
	}

	return response, err
}
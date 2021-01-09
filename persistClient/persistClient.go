package persistclient

import "github.com/DJost494731/PersistClient/httpclient"

//PersistClient interacts with a Persist database over http, without forcing the user to do much setup.
type PersistClient struct {
	//HttpClient that the PersistClient will use.
	HttpClient httpclient.IHttpClient
	//PersistUrl is the url your prefix server is hosted at without the folder or file path. Ex: "http://www.foo.com"
	PersistUrl string
}

//GetData gets data at the path from persist. Returns RecordNotFoundError if no data is found.
func (p PersistClient) GetData(folder string, file string) (string, error) {
	response, statusCode, err  :=  p.HttpClient.Get(p.PersistUrl + "/" + folder + "/" + file)

	if statusCode == 404 {
		return "", &RecordNotFoundError{Path: folder + "/" + file} 
	}

	return response, err
}
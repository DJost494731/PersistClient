### Persist Client
The persist client aims to further simplify interacting with the persist database. 

### Usage
First initialize the client:

```golang
persistClient := PersistClient{
  HttpClient: httpclient.HttpClient{} //Default IHttpClient provided with persist
  PersistUrl: "http://www.example.com" //Url your persist server is hosted at. 
}
```
Then make requests:
```golang
dataAsString, err := persistClient.GetData("folder","file")
```

### Testing
To test applications that use the persist client, use IDatabaseClient and the included MockDatabaseClient.

For example:
```golang
want := "result" 

var mockDbClient IDatabaseClient

mockDbClient = persistclient.MockDatabaseClient{
  DataToReturn: want
}

got := mockDatabaseClient.GetData("any-folder", "any-file")

fmt.Println(got) //prints "result"
``` 
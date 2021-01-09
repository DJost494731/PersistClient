package persistclient

import "fmt"

//IDatabaseClient is an interface for getting/pushing data to and from a database. PersistClient implements IDatabaseClient.
type IDatabaseClient interface {
	GetData(folder string, file string) (string, error)
}

//MockDatabaseClient includes functionality for mocking out an IDatabaseClient
type MockDatabaseClient struct {
	DataToReturn      string
	ErrorToReturn     error
	FolderInvocations []string
	FileInvocations   []string
}

//RecordNotFoundError is an error that should be returned when an IDatabaseClient cannot find the requested record.
type RecordNotFoundError struct {
	Path string
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("Path " + e.Path + " was not found!")
}

//GetData returns MockDatabaseClient.DataToReturn and MockDatabaseClient.ErrorToReturn. Also records folder/file invocations.
func (mockDb *MockDatabaseClient) GetData(folder string, file string) (string, error) {
	mockDb.ensureArraysAreInitialized()

	mockDb.FolderInvocations = append(mockDb.FolderInvocations, folder)
	mockDb.FileInvocations = append(mockDb.FileInvocations, file)

	return mockDb.DataToReturn, mockDb.ErrorToReturn
}

func (mockDb *MockDatabaseClient) ensureArraysAreInitialized() {
	if mockDb.FolderInvocations == nil {
		mockDb.FolderInvocations = []string{}
	}
	if mockDb.FileInvocations == nil {
		mockDb.FileInvocations = []string{}
	}
}
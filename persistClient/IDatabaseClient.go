package persistclient

type IDatabaseClient interface {
	GetData(folder string, file string) (string, error)
}

type MockDatabaseClient struct {
	DataToReturn      string
	errorToReturn     error
	FolderInvocations []string
	FileInvocations   []string
}

func (mockDb *MockDatabaseClient) GetData(folder string, file string) (string, error) {
	mockDb.ensureArraysAreInitialized()

	mockDb.FolderInvocations = append(mockDb.FolderInvocations, folder)
	mockDb.FileInvocations = append(mockDb.FileInvocations, file)

	return mockDb.DataToReturn, mockDb.errorToReturn
}

func (mockDb *MockDatabaseClient) ensureArraysAreInitialized() {
	if mockDb.FolderInvocations == nil {
		mockDb.FolderInvocations = []string{}
	}
	if mockDb.FileInvocations == nil {
		mockDb.FileInvocations = []string{}
	}
}
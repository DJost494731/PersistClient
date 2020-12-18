package persistclient

type IDatabaseClient interface {
	GetData(folder string, file string) string
}

type MockDatabaseClient struct {
	DataToReturn      string
	FolderInvocations []string
	FileInvocations   []string
}

func (mockDb *MockDatabaseClient) GetData(folder string, file string) string {
	mockDb.ensureArraysAreInitialized()

	mockDb.FolderInvocations = append(mockDb.FolderInvocations, folder)
	mockDb.FileInvocations = append(mockDb.FileInvocations, file)

	return mockDb.DataToReturn
}

func (mockDb *MockDatabaseClient) ensureArraysAreInitialized() {
	if mockDb.FolderInvocations == nil {
		mockDb.FolderInvocations = []string{}
	}
	if mockDb.FileInvocations == nil {
		mockDb.FileInvocations = []string{}
	}
}
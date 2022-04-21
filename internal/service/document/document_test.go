package document_test

import (
	"context"
	"log"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/service/document"
	"notary-public-online/mocks"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockDB *mocks.MockDB

var mockStorage *mocks.MockStorage

func setupSuite(t testing.TB) func(t testing.TB) {
	log.Println("setup suite")

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	// Mock Database
	mockDB = mocks.NewMockDB(mockctl)

	// Mock Storage
	mockStorage = mocks.NewMockStorage(mockctl)

	// Return a function to teardown the test
	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

func TestStoreDocument(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	doc := document.New(mockDB, mockStorage)

	// file and file address
	file, _ := os.Open("../../../mocks/fakeFile.txt")

	mockDB.EXPECT().CreateDocument(gomock.Any(), "doc name", "doc description", "fakeFile.txt", gomock.Any(), 1, false).Return(nil).Times(1)
	mockStorage.EXPECT().StoreFile(file).Return("fakeFile.txt", nil).Times(1)

	err := doc.StoreDocument(context.TODO(), file, "doc name", "doc description", 1)

	assert.Nil(t, err)
}

func TestDocumentDetails(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	doc := document.New(mockDB, mockStorage)

	mockDB.EXPECT().GetDocument(gomock.Any(), 1).Return(&model.Document{}, nil).Times(1)

	_, err := doc.DocumentDetails(context.TODO(), 1)

	assert.Nil(t, err)
}

func TestReadDocument(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	doc := document.New(mockDB, mockStorage)

	// store new Document
	file, _ := os.Open("../../../mocks/fakeFile.txt")
	mockDB.EXPECT().CreateDocument(gomock.Any(), "doc name", "doc description", "fakeFile.txt", gomock.Any(), 1, false).Return(nil).Times(1)
	mockStorage.EXPECT().StoreFile(file).Return("fakeFile.txt", nil).Times(1)
	doc.StoreDocument(context.TODO(), file, "doc name", "doc description", 1)

	// Read Document
	mockDB.EXPECT().GetDocumentAddress(gomock.Any(), 1).Return("fakeFile.txt", nil).Times(1)
	mockStorage.EXPECT().ReadFile(gomock.Any()).Return(file, nil).Times(1)

	_, err := doc.ReadDocument(context.TODO(), 1)
	assert.Nil(t, err)
}

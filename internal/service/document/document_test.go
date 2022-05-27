package document_test

import (
	"context"
	"log"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/pkg/storage/disk"
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

func TestStoreDocumentWithDuplicatedIdempotenct(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	doc := document.New(mockDB, mockStorage)

	// file and file address
	file, _ := os.Open("../../../mocks/fakeFile.txt")

	user := model.User{
		Id:    1,
		Email: "test@gmail.com",
	}

	idempotentKey := "1qaz"

	mockDB.EXPECT().CheckDocumentIdempotency(gomock.Any(), idempotentKey).Return(true).Times(1)
	mockDB.EXPECT().GetUserWithEmail(gomock.Any(), user.Email).Return(user, nil).Times(1)

	err := doc.StoreDocument(context.TODO(), idempotentKey, file, "doc name", "doc description", user.Email)

	assert.Nil(t, err)
}

func TestStoreDocument(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	doc := document.New(mockDB, mockStorage)

	// file and file address
	file, _ := os.Open("../../../mocks/fakeFile.txt")

	user := model.User{
		Id:    1,
		Email: "test@gmail.com",
	}

	document := model.Document{
		Name:        "doc name",
		Description: "doc description",
		FileAddress: "fakeFile.txt",
		UserId:      user.Id,
	}

	idempotentKey := "1qaz"

	mockDB.EXPECT().CheckDocumentIdempotency(gomock.Any(), idempotentKey).Return(false).Times(1)
	mockDB.EXPECT().CreateDocument(gomock.Any(), idempotentKey, document.Name, document.Description, document.FileAddress, gomock.Any(), document.UserId, false).Return(nil).Times(1)
	mockDB.EXPECT().GetUserWithEmail(gomock.Any(), user.Email).Return(user, nil).Times(1)
	mockStorage.EXPECT().StoreFile(file).Return("fakeFile.txt", nil).Times(1)

	err := doc.StoreDocument(context.TODO(), idempotentKey, file, "doc name", "doc description", user.Email)

	assert.Nil(t, err)
}

func TestStoreDocumentIntegration(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	storage := disk.New()

	doc := document.New(mockDB, storage)

	// file and file address
	file, _ := os.Open("../../../mocks/fakeFile.txt")

	user := model.User{
		Id:    1,
		Email: "test@gmail.com",
	}

	document := model.Document{
		Name:        "doc name",
		Description: "doc description",
		FileAddress: "fakeFile.txt",
		UserId:      user.Id,
	}

	idempotentKey := "1qaz"

	mockDB.EXPECT().CheckDocumentIdempotency(gomock.Any(), idempotentKey).Return(false).Times(1)
	mockDB.EXPECT().CreateDocument(gomock.Any(), idempotentKey, document.Name, document.Description, gomock.Any(), gomock.Any(), 1, false).Return(nil).Times(1)
	mockDB.EXPECT().GetUserWithEmail(gomock.Any(), user.Email).Return(user, nil).Times(1)

	err := doc.StoreDocument(context.TODO(), idempotentKey, file, "doc name", "doc description", user.Email)

	assert.Nil(t, err)
}

func TestDocumentDetails(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	doc := document.New(mockDB, mockStorage)

	mockDB.EXPECT().GetDocument(gomock.Any(), 1).Return(model.Document{}, nil).Times(1)

	_, err := doc.DocumentDetails(context.TODO(), 1)

	assert.Nil(t, err)
}

func TestReadDocument(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	doc := document.New(mockDB, mockStorage)

	// store new Document
	file, _ := os.Open("../../../mocks/fakeFile.txt")

	user := model.User{
		Id:    1,
		Email: "test@gmail.com",
	}

	document := model.Document{
		Name:        "doc name",
		Description: "doc description",
		FileAddress: "fakeFile.txt",
		UserId:      user.Id,
	}

	idempotentKey := "1qaz"

	mockDB.EXPECT().CheckDocumentIdempotency(gomock.Any(), idempotentKey).Return(false).Times(1)
	mockDB.EXPECT().CreateDocument(gomock.Any(), idempotentKey, document.Name, document.Description, document.FileAddress, gomock.Any(), document.UserId, false).Return(nil).Times(1)
	mockDB.EXPECT().GetUserWithEmail(gomock.Any(), user.Email).Return(user, nil).Times(1)
	mockStorage.EXPECT().StoreFile(file).Return("fakeFile.txt", nil).Times(1)
	doc.StoreDocument(context.TODO(), idempotentKey, file, "doc name", "doc description", user.Email)

	// Read Document
	mockDB.EXPECT().GetDocumentAddress(gomock.Any(), 1).Return("fakeFile.txt", nil).Times(1)
	mockStorage.EXPECT().ReadFile(gomock.Any()).Return(file, nil).Times(1)

	_, err := doc.ReadDocument(context.TODO(), 1)
	assert.Nil(t, err)
}

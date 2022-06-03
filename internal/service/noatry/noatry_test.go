package noatry_test

import (
	"context"
	"log"
	"notary-public-online/internal/dto"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/pkg/hash/sha256"
	"notary-public-online/internal/pkg/pairKey/rsa"
	"notary-public-online/internal/service/noatry"
	"notary-public-online/mocks"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockDB *mocks.MockDB
var fileHash []byte

func setupSuite(t testing.TB) func(t testing.TB) {

	log.Println("setup suite")

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	// Mock Database
	mockDB = mocks.NewMockDB(mockctl)

	// make new document and its hash
	file, _ := os.Open("../../../mocks/fakeFile.txt")
	hasher := sha256.New()
	fileHash, _ = hasher.Hash(file)

	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

func TestCreateNoatry(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	notary := model.Notary{
		DocumentId:   1,
		UserId:       1,
		PartnerCount: 2,
		Completed:    false,
	}

	user := model.User{
		Id:    1,
		Email: "ahmad@gmail.com",
	}

	mockDB.EXPECT().CheckDocumentIdempotency(gomock.Any(), "1qaz").Return(1).Times(1)
	mockDB.EXPECT().GetUserWithEmail(gomock.Any(), "ahmad@gmail.com").Return(user, nil).Times(1)
	mockDB.EXPECT().CreateNoatry(gomock.Any(), notary.DocumentId, notary.UserId, notary.PartnerCount, notary.Completed).Return(nil).Times(1)

	noatry := noatry.New(mockDB)

	err := noatry.CreateNoatry(context.TODO(), &dto.StoreNoatryCredential{UserEmail: "ahmad@gmail.com", DocumentIdempotent: "1qaz", PartnerCount: 2})

	assert.Nil(t, err)
}

func TestGetNoatry(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	mockDB.EXPECT().GetNoatry(gomock.Any(), 1).Return(model.Notary{}, nil).Times(1)

	noatry := noatry.New(mockDB)

	_, err := noatry.GetNoatry(context.TODO(), 1)

	assert.Nil(t, err)
}

func TestSignNoatry(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	userModel := model.User{
		Id:    1,
		Email: "test@gmail.com",
	}

	noatryModel := model.Notary{
		Id:           1,
		DocumentId:   1,
		UserId:       1,
		PartnerCount: 1,
		Completed:    false,
	}

	signature := model.Signature{
		UserId:         1,
		NotaryId:       noatryModel.Id,
		SignedDocument: fileHash,
	}

	key := rsa.NewKeys()
	key.PairKeyGenerator(userModel.Email)

	mockDB.EXPECT().GetUserWithId(gomock.Any(), 1).Return(userModel, nil).Times(1)
	mockDB.EXPECT().GetDocumentHash(gomock.Any(), 1).Return(fileHash, nil).Times(1)
	mockDB.EXPECT().GetNoatry(gomock.Any(), 1).Return(noatryModel, nil).Times(1)
	mockDB.EXPECT().CreateSignature(gomock.Any(), 1, 1, gomock.Any()).Return(signature, nil).Times(1)

	noatry := noatry.New(mockDB)

	err := noatry.SignNoatry(context.TODO(), 1, 1)

	assert.Nil(t, err)
}

func TestVerifyNoatrySignature(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	userModel := model.User{
		Id:    1,
		Email: "test@gmail.com",
	}

	noatryModel := model.Notary{
		Id:           1,
		DocumentId:   1,
		UserId:       1,
		PartnerCount: 1,
		Completed:    false,
	}

	// generate pairKey for user
	key := rsa.NewKeys()
	pr, pu, _ := key.PairKeyGenerator("test@gmail.com")

	// sign document
	crypto := rsa.New(pr, pu)
	signedDoc, _ := crypto.Signature(fileHash)

	mockDB.EXPECT().GetUserWithId(gomock.Any(), 1).Return(userModel, nil).Times(1)
	mockDB.EXPECT().GetNoatry(gomock.Any(), 1).Return(noatryModel, nil).Times(1)
	mockDB.EXPECT().GetDocumentHash(gomock.Any(), 1).Return(fileHash, nil).Times(1)
	mockDB.EXPECT().GetSignatures(gomock.Any(), 1, 1).Return(&signedDoc, nil).Times(1)

	noatry := noatry.New(mockDB)
	verify, err := noatry.VerifyNoatrySignature(context.TODO(), 1, 1)

	assert.Nil(t, err)
	assert.True(t, verify)
}

package user_test

import (
	"context"
	"log"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/service/user"
	"notary-public-online/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockDB *mocks.MockDB
var mockPairKey *mocks.MockKeys
var mockPassHash *mocks.MockPasswordHash

func setupSuite(t testing.TB) func(t testing.TB) {
	log.Println("user service test suite setup")

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	// Mock Database
	mockDB = mocks.NewMockDB(mockctl)

	// Mock Pair Keys
	mockPairKey = mocks.NewMockKeys(mockctl)

	// Mock Password Hash
	mockPassHash = mocks.NewMockPasswordHash(mockctl)

	return func(t testing.TB) {
		log.Println("user service test suite teardown")
	}
}

func TestRegister(t *testing.T) {
	tearDown := setupSuite(t)
	defer tearDown(t)

	userInp := model.User{
		FirstName:   "ahmad",
		LastName:    "abd",
		Email:       "ahmad@gmail.com",
		Password:    "123456",
		Citizenship: "indonesia",
	}

	privateKey := "pr key"
	publicKey := "pu key"
	passHash := "123456"

	checkUserExistanceMK := mockDB.EXPECT().CheckUserExistanceWithEmail(gomock.Any(), userInp.Email).Return(false, nil).Times(1)
	passHashMK := mockPassHash.EXPECT().HashPassword(userInp.Password).Return(passHash, nil).After(checkUserExistanceMK).Times(1)
	pairKeyMK := mockPairKey.EXPECT().PairKeyGenerator().Return(privateKey, publicKey, nil).After(passHashMK).Times(1)
	mockDB.EXPECT().CreateUser(gomock.Any(), &userInp).Return(func() model.User {
		userInp.PrivateKey = privateKey
		userInp.PublicKey = publicKey
		userInp.Password = passHash
		return userInp
	}(), nil).After(pairKeyMK).Times(1)

	userSrv := user.New(mockDB, mockPairKey, mockPassHash)

	user, err := userSrv.Register(context.TODO(), userInp)

	assert.Equal(t, user, userInp)
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	tearDown := setupSuite(t)
	defer tearDown(t)

	userInp := model.User{
		FirstName:   "ahmad",
		LastName:    "abd",
		Email:       "ahmad@gmail.com",
		Password:    "654321",
		Citizenship: "indonesia",
	}

	pass := "123456"

	getUserMK := mockDB.EXPECT().GetUserWithEmail(gomock.Any(), userInp.Email).Return(userInp, nil).Times(1)
	mockPassHash.EXPECT().CheckPasswordHash(pass, userInp.Password).Return(true).After(getUserMK).Times(1)

	userSrv := user.New(mockDB, mockPairKey, mockPassHash)
	userSrv.Login(context.TODO(), userInp.Email, pass)
}

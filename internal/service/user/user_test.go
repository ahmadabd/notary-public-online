package user_test

import (
	"context"
	"errors"
	"log"
	"notary-public-online/internal/dto"
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

	inp := dto.RegisterCredential{
		FirstName:       "ahmad",
		LastName:        "abd",
		Email:           "ahmad@gmail.com",
		Password:        "123456",
		ConfirmPassword: "123456",
		Citizenship:     "indonesia",
	}

	privateKey := "pr key"
	publicKey := "pu key"
	passHash := "123456"

	userModel := model.User{
		FirstName:   inp.FirstName,
		LastName:    inp.LastName,
		Email:       inp.Email,
		Citizenship: inp.Citizenship,
	}

	checkUserExistanceMK := mockDB.EXPECT().CheckUserExistanceWithEmail(gomock.Any(), inp.Email).Return(false, nil).Times(1)
	passHashMK := mockPassHash.EXPECT().HashPassword(inp.Password).Return(passHash, nil).After(checkUserExistanceMK).Times(1)
	pairKeyMK := mockPairKey.EXPECT().PairKeyGenerator(inp.Email).Return([]byte(privateKey), []byte(publicKey), nil).After(passHashMK).Times(1)
	mockDB.EXPECT().CreateUser(gomock.Any(), &userModel).Return(func() model.User {
		userModel.Password = passHash
		return userModel
	}(), nil).After(pairKeyMK).Times(1)

	userSrv := user.New(mockDB, mockPairKey, mockPassHash)
	err := userSrv.Register(context.TODO(), inp)
	assert.Equal(t, nil, err)
}

func TestRegisterWhenEmailExists(t *testing.T) {
	tearDown := setupSuite(t)
	defer tearDown(t)

	inp := dto.RegisterCredential{
		FirstName:       "ahmad",
		LastName:        "abd",
		Email:           "ahmad@gmail.com",
		Password:        "123456",
		ConfirmPassword: "123456",
		Citizenship:     "indonesia",
	}

	mockDB.EXPECT().CheckUserExistanceWithEmail(gomock.Any(), inp.Email).Return(true, nil).Times(1)

	userSrv := user.New(mockDB, mockPairKey, mockPassHash)
	err := userSrv.Register(context.TODO(), inp)
	assert.Equal(t, errors.New("user with this email exist in system"), err)
}

func TestRegisterWithNotSamePassword(t *testing.T) {
	tearDown := setupSuite(t)
	defer tearDown(t)

	inp := dto.RegisterCredential{
		FirstName:       "ahmad",
		LastName:        "abd",
		Email:           "ahmad@gmail.com",
		Password:        "123456",
		ConfirmPassword: "2",
		Citizenship:     "indonesia",
	}

	mockDB.EXPECT().CheckUserExistanceWithEmail(gomock.Any(), inp.Email).Return(false, nil).Times(1)

	userSrv := user.New(mockDB, mockPairKey, mockPassHash)
	err := userSrv.Register(context.TODO(), inp)
	assert.Equal(t, errors.New("password and confirm password does not match"), err)
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
	userSrv.Login(context.TODO(), dto.LoginCredential{Email: userInp.Email, Password: pass})
}

package user

import (
	"context"
	"errors"
	"log"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/pkg/pairKey"
	"notary-public-online/internal/pkg/passwordHash"
	"notary-public-online/internal/repository"
)

type user struct {
	Db       repository.DB
	Key      pairKey.Keys
	PassHash passwordHash.PasswordHash
}

func New(db repository.DB, keys pairKey.Keys, passHash passwordHash.PasswordHash) UserImp {
	return &user{Db: db, Key: keys, PassHash: passHash}
}

func (u *user) Register(ctx context.Context, user model.User) (model.User, error) {

	// check user with this email dosent exist
	email := user.Email
	if exists, err := u.Db.CheckUserExistanceWithEmail(ctx, email); err != nil {
		log.Panicln("CheckUserExistanceWithEmail failed in registering user: ", err)
		return model.User{}, err
	} else if exists {
		return model.User{}, errors.New("user with this email exist in system")
	}

	// decrypet password
	password, _ := u.PassHash.HashPassword(user.Password)
	user.Password = password

	prKey, puKey, err := u.Key.PairKeyGenerator()

	if err != nil {
		log.Panicln("Generating pairKey failed")
		return model.User{}, err
	}

	user.PrivateKey = prKey
	user.PublicKey = puKey

	user, err = u.Db.CreateUser(ctx, &user)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *user) Login(ctx context.Context, email string, password string) (bool, error) {
	user, err := u.Db.GetUserWithEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if !u.PassHash.CheckPasswordHash(password, user.Password) {
		return false, errors.New("invalid password")
	}

	return true, nil
}

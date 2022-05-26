package user

import (
	"context"
	"errors"
	"log"
	"notary-public-online/internal/dto"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/pkg/pairKey"
	"notary-public-online/internal/pkg/passwordHash"
	"notary-public-online/internal/repository"
)

type userImp struct {
	Db       repository.DB
	Key      pairKey.Keys
	PassHash passwordHash.PasswordHash
}

func New(db repository.DB, keys pairKey.Keys, passHash passwordHash.PasswordHash) User {
	return &userImp{Db: db, Key: keys, PassHash: passHash}
}

func (u *userImp) Register(ctx context.Context, inp dto.RegisterCredential) error {

	// check user with this email dosent exist
	email := inp.Email
	if exists, err := u.Db.CheckUserExistanceWithEmail(ctx, email); err != nil {
		return err
	} else if exists {
		return errors.New("user with this email exist in system")
	}

	// check password and confirm password match
	if inp.Password != inp.ConfirmPassword {
		return errors.New("password and confirm password does not match")
	}

	user := model.User{
		Email:       inp.Email,
		FirstName:   inp.FirstName,
		LastName:    inp.LastName,
		Citizenship: inp.Citizenship,
	}

	user.Password = decrypetPassword(u, inp.Password)

	_, _, err := u.Key.PairKeyGenerator(email)

	if err != nil {
		log.Println("Generating pairKey failed")
		return err
	}

	_, err = u.Db.CreateUser(ctx, &user)

	if err != nil {
		return err
	}

	return nil
}

func decrypetPassword(u *userImp, inpPass string) string {
	password, _ := u.PassHash.HashPassword(inpPass)

	return password
}

func (u *userImp) Login(ctx context.Context, inp dto.LoginCredential) (bool, error) {
	user, err := u.Db.GetUserWithEmail(ctx, inp.Email)
	if err != nil {
		return false, err
	}

	if !u.PassHash.CheckPasswordHash(inp.Password, user.Password) {
		return false, errors.New("invalid password")
	}

	return true, nil
}

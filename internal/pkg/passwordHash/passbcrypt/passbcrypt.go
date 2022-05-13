package passbcrypt

import (
	passwordhash "notary-public-online/internal/pkg/passwordHash"

	"golang.org/x/crypto/bcrypt"
)

type passwordHashImp struct{}

func New() passwordhash.PasswordHash {
	return &passwordHashImp{}
}

func (p *passwordHashImp) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *passwordHashImp) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"notary-public-online/internal/pkg/pairKey"
)

type Crypto struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func New(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) pairKey.CryptoImp {
	return &Crypto{privateKey: privateKey, publicKey: publicKey}
}

func PairKeyGenerator() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	publicKey := privateKey.PublicKey

	return privateKey, &publicKey, err
}

func (c *Crypto) Signature(hashedInput *string) (string, error) {

	signature, err := rsa.SignPSS(rand.Reader, c.privateKey, crypto.SHA256, []byte(*hashedInput), nil)

	return string(signature), err
}

func (c *Crypto) VerifySignature(hashedInput *string, signature *string) bool {
	if err := rsa.VerifyPSS(c.publicKey, crypto.SHA256, []byte(*hashedInput), []byte(*signature), nil); err != nil {
		return true
	}

	return false
}

func (c *Crypto) Encryption(input *string) (string, error) {
	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, c.publicKey, []byte(*input), nil)

	return string(encrypted), err
}

func (c *Crypto) Decryption(input *string) (string, error) {
	decrypted, err := c.privateKey.Decrypt(nil, []byte(*input), &rsa.OAEPOptions{Hash: crypto.SHA256})

	return string(decrypted), err
}

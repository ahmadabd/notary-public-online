package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"notary-public-online/internal/pkg/pairKey"
)

type Crypto struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func New(privateKey string, publicKey string) pairKey.CryptoImp {

	pb, errPr := x509.ParsePKCS1PublicKey([]byte(publicKey))
	pr, errPu := x509.ParsePKCS1PrivateKey([]byte(privateKey))

	if errPr != nil || errPu != nil {
		panic("Error parsing key")
	}

	return &Crypto{privateKey: pr, publicKey: pb}
}

func PairKeyGenerator() (string, string, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey

	bytePr := x509.MarshalPKCS1PrivateKey(privateKey)
	bytePu := x509.MarshalPKCS1PublicKey(&publicKey)

	return string(bytePr), string(bytePu), err
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

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"notary-public-online/internal/pkg/pairKey"
)

type keyImp struct{}

func NewKeys() pairKey.Keys {
	return &keyImp{}
}

func (k *keyImp) PairKeyGenerator(email string) ([]byte, []byte, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey

	filename := "/tmp/" + email

	bytePr := x509.MarshalPKCS1PrivateKey(privateKey)
	bytePu := x509.MarshalPKCS1PublicKey(&publicKey)

	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: bytePr,
		},
	)
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: bytePu,
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile(filename+".rsa", keyPEM, 0700); err != nil {
		panic(err)
	}

	// Write public key to file.
	if err := ioutil.WriteFile(filename+".rsa.pub", pubPEM, 0755); err != nil {
		panic(err)
	}

	return bytePr, bytePu, err
}

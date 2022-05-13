package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"notary-public-online/internal/pkg/pairKey"
)

type keyImp struct{}

func NewKeys() pairKey.Keys {
	return &keyImp{}
}

func (k *keyImp) PairKeyGenerator() (string, string, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey

	bytePr := x509.MarshalPKCS1PrivateKey(privateKey)
	bytePu := x509.MarshalPKCS1PublicKey(&publicKey)

	return string(bytePr), string(bytePu), err
}

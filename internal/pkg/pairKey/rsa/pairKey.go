package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"notary-public-online/internal/pkg/pairKey"
)

var baseDir = "/tmp/"

type keyImp struct{}

func NewKeys() pairKey.Keys {
	return &keyImp{}
}

func (k *keyImp) PairKeyGenerator(email string) ([]byte, []byte, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		log.Println("Cannot generate RSA key")
	}

	publicKey := &privateKey.PublicKey

	filename := baseDir + email

	var bytePr []byte = x509.MarshalPKCS1PrivateKey(privateKey)
	var bytePu []byte = x509.MarshalPKCS1PublicKey(publicKey)

	privateKeyBlock := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: bytePr,
		},
	)
	publicKeyBlock := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: bytePu,
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile(filename+".rsa", privateKeyBlock, 0700); err != nil {
		panic(err)
	}

	// Write public key to file.
	if err := ioutil.WriteFile(filename+".rsa.pub", publicKeyBlock, 0755); err != nil {
		panic(err)
	}

	return bytePr, bytePu, err
}

func (k *keyImp) PairKeyReader(filename string) ([]byte, []byte, error) {

	// Read private key from file.
	privateKeyBlock, err := ioutil.ReadFile(baseDir + filename + ".rsa")

	if err != nil {
		log.Println("Cannot read RSA private key")
	}

	// Read public key from file.
	publicKeyBlock, err := ioutil.ReadFile(baseDir + filename + ".rsa.pub")

	if err != nil {
		log.Println("Cannot read RSA public key")
	}

	blockPr, _ := pem.Decode(privateKeyBlock)
	blockPu, _ := pem.Decode(publicKeyBlock)

	return blockPr.Bytes, blockPu.Bytes, err
}

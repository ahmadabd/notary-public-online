package rsa_test

import (
	"notary-public-online/internal/pkg/hash/sha256"
	"notary-public-online/internal/pkg/pairKey/rsa"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairKeyGenerate(t *testing.T) {
	key := rsa.NewKeys()
	pr, pu, err := key.PairKeyGenerator("test@gmail.com")

	assert.Nil(t, err)
	assert.NotNil(t, pr)
	assert.NotNil(t, pu)
}

func TestPairKeyReader(t *testing.T) {
	key := rsa.NewKeys()
	key.PairKeyGenerator("test@gmail.com")

	t.Run("Check_reader", func(t *testing.T) {
		prReader, puReader, err := key.PairKeyReader("test@gmail.com")
		assert.Nil(t, err)

		assert.NotNil(t, prReader)
		assert.NotNil(t, puReader)
	})
}

func createNewFile() *os.File {
	// create new file
	os.WriteFile("/tmp/test.txt", []byte("test"), 0644)

	// open file
	input, _ := os.Open("/tmp/test.txt")

	defer input.Close()

	return input
}

func TestSign(t *testing.T) {
	key := rsa.NewKeys()
	pr, pu, _ := key.PairKeyGenerator("test@gmail.com")
	crypto := rsa.New(pr, pu)

	input := createNewFile()

	hash := sha256.New()
	hashedInput, _ := hash.Hash(input)

	t.Run("Check_signuture", func(t *testing.T) {
		signature, err := crypto.Signature(&hashedInput)
		assert.Nil(t, err)
		assert.NotNil(t, signature)
	})
}

func TestSignVerification(t *testing.T) {
	key := rsa.NewKeys()
	pr, pu, _ := key.PairKeyGenerator("test@gmail.com")
	crypto := rsa.New(pr, pu)

	input := createNewFile()

	hash := sha256.New()
	hashedInput, _ := hash.Hash(input)

	t.Run("Check_verification", func(t *testing.T) {
		signature, err := crypto.Signature(&hashedInput)
		assert.Nil(t, err)
		assert.NotNil(t, signature)

		res := crypto.VerifySignature(&signature, &hashedInput)
		assert.True(t, res)
	})
}

func TestEncryption(t *testing.T) {
	key := rsa.NewKeys()
	pr, pu, _ := key.PairKeyGenerator("test@gmail.com")
	crypto := rsa.New(pr, pu)

	input := "Hello World"
	encrptedInp, err := crypto.Encryption(&input)

	assert.Nil(t, err)
	assert.NotNil(t, encrptedInp)
}

func TestDecryption(t *testing.T) {
	key := rsa.NewKeys()
	pr, pu, _ := key.PairKeyGenerator("test@gmail.com")
	crypto := rsa.New(pr, pu)

	input := "Hello World"
	encrptedInp, _ := crypto.Encryption(&input)

	decryptedInp, err := crypto.Decryption(&encrptedInp)

	assert.Nil(t, err)
	assert.NotNil(t, decryptedInp)

	assert.Equal(t, input, decryptedInp)
}

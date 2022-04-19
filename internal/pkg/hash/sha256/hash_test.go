package sha256_test

import (
	"notary-public-online/internal/pkg/hash/sha256"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createNewFile() *os.File {
	// create new file
	os.WriteFile("/tmp/test.txt", []byte("test"), 0644)

	// open file
	input, _ := os.Open("/tmp/test.txt")
	
	defer input.Close()

	return input
}

func TestHasher(t *testing.T) {

	hasher := sha256.New()

	input := createNewFile()

	res, err := hasher.Hash(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}
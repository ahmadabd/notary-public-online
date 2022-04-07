package sha256_test

import (
	"notary-public-online/internal/pkg/hash/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasher(t *testing.T) {

	hasher := sha256.New()

	input := "Hello Hash"
	res := hasher.Hash(input)

	assert.NotEmpty(t, res)
}

func TestCheckHasher(t *testing.T) {

	hasher := sha256.New()

	input := "Hello Hash"
	inputHash := hasher.Hash(input)

	res := hasher.HashChecker(input, inputHash)

	assert.Equal(t, res, true)
}

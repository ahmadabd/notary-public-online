package sha256

import (
	"crypto/sha256"
	"notary-public-online/internal/pkg/hash"
)

type Hasher struct{}

func New() hash.HasherImp {
	return &Hasher{}
}

func (h *Hasher) Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	return string(hasher.Sum(nil))
}

func (h *Hasher) HashChecker(data string, hash string) bool {
	return h.Hash(data) == hash
}

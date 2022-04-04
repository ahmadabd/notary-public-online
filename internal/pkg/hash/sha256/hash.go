package sha256

import (
	"crypto/sha256"
	"encoding/base64"
	"notary-public-online/internal/pkg/hash"
)

type Hasher struct {}

func NewHash() hash.HasherImp {
	return &Hasher{}
}

func (h *Hasher) Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (h *Hasher) HashChecker(data string, hash string) bool {
	return h.Hash(data) == hash
}

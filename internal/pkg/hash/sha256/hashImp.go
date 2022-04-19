package sha256

import (
	"crypto/sha256"
	"io"
	"log"
	"notary-public-online/internal/pkg/hash"
	"os"
)

type HasherImp struct{}

func New() hash.Hasher {
	return &HasherImp{}
}

func (h *HasherImp) Hash(data *os.File) ([]byte, error) {
	hasher := sha256.New()

	if _, err := io.Copy(hasher, data); err != nil {
		log.Println(err)
	}

	return hasher.Sum(nil), nil
}

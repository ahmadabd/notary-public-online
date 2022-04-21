package hash

import (
	"os"
)

type Hasher interface {
	Hash(*os.File) ([]byte, error)
}

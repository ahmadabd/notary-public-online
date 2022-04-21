package storage

import (
	"os"
)

type Storage interface {
	StoreFile(file *os.File) (string, error)

	ReadFile(fileName string) (*os.File, error)
}

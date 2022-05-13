package storage

import (
	"os"
)

//go:generate $HOME/go_projects/bin/mockgen -destination=../../../mocks/mock_storage.go -package=mocks notary-public-online/internal/pkg/storage Storage
type Storage interface {
	StoreFile(file *os.File) (string, error)

	ReadFile(fileName string) (*os.File, error)
}

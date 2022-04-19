package disk

import (
	"io/ioutil"
	"notary-public-online/internal/pkg/random"
	"notary-public-online/internal/pkg/storage"
	"os"
)

type storageImp struct{}

func New() storage.Storage {
	return &storageImp{}
}

func (s *storageImp) StoreFile(file *os.File) (string, error) {
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return "", err
	}

	fileName := random.RandStringRunes(64)

	err = ioutil.WriteFile(fileName, content, 0644)

	if err != nil {
		return "", err
	}

	return fileName, nil
}

func (s *storageImp) ReadFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

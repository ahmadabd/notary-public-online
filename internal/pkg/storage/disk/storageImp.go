package disk

import (
	"io/ioutil"
	"notary-public-online/internal/pkg/random"
	"notary-public-online/internal/pkg/storage"
	"os"
	"path/filepath"
)

type storageImp struct{}

func New() storage.Storage {
	return &storageImp{}
}

var baseDir = "/tmp/"

func (s *storageImp) StoreFile(file *os.File) (string, error) {
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return "", err
	}

	extention := filepath.Ext(file.Name())

	fileName := random.RandStringRunes(8) + extention

	err = ioutil.WriteFile(baseDir+fileName, content, 0644)

	if err != nil {
		return "", err
	}

	return fileName, nil
}

func (s *storageImp) ReadFile(fileName string) (*os.File, error) {
	return os.Open(baseDir + fileName)
}

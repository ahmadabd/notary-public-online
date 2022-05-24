package dto

import "os"

type StoreDocumentCredential struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Document    *os.File `json:"document"`
}

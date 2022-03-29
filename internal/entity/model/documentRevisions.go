package model

type DocumentRevisions struct {
	Id         int    `json:"id"`
	DocumentId int    `json:"documentId"`
	Slug       string `json:"slug"`
	Revision   int    `json:"revision"`
	Hash       string `json:"hash"`
	Active     bool   `json:"active"`			// If the document is active it cant be Delete or Update
}

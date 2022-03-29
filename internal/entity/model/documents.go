package model

type Document struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`			// creator id
	Hash        string `json:"hash"`
	Active      bool   `json:"active"`			// when a document is active that used in a contract, If the document is active it cant be Delete or Update
}

package model

type Document struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug		string `json:"slug"`
	UserId      int    `json:"userId"`
	Hash        string `json:"hash"`
	Active      bool   `json:"active"`			// If the document is active it cant be Delete or Update
}

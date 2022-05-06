package model

import "time"

type Document struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FileAddress string    `json:"fileAddress"`
	User        User      `json:"user"`
	UserId      int       `json:"-"` // creator id
	Hash        []byte    `json:"hash"`
	Active      bool      `json:"active"` // when a document is active that used in a contract, If the document is active it cant be Delete or Update
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

package model

import "time"

type Signature struct {
	Id             int       `json:"id"`
	User           User      `json:"user"`
	UserId         int       `json:"-"`
	SignedDocument []byte    `json:"signedDocument"`
	Notary         Notary    `json:"notary"`
	NotaryId       int       `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

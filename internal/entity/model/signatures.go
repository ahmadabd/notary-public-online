package model

type Signature struct {
	Id             int    `json:"id"`
	UserId         int    `json:"userId"`
	SignedDocument string `json:"signedDocument"`
	NoatryId       int    `json:"noatrytId"`
}

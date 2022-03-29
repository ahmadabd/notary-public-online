package model

type UsersDocuments struct {
	Id                 int    `json:"id"`
	UserId             int    `json:"userId"`
	SignedDocument     string `json:"signedDocument"`		// signed document is hash of document that signed using private key of users
	DocumentRevisionId int    `json:"documentRevisionId"`
	Step			   int    `json:"step"`					// Number of members that should sign the document
	Completed          bool   `json:"completed"`			// If all the members have signed the document Completed = true
}

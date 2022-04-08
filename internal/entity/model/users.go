package model

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Citizenship string `json:"citizenship"`
	PrivateKey  string `json:"privateKey"`
	PublicKey   string `json:"publicKey"`
}

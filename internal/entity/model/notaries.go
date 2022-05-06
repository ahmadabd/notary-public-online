package model

import "time"

type Notary struct {
	Id           int       `json:"id"`
	User         User      `json:"user"`
	UserId       int       `json:"-"` // owner
	Document     Document  `json:"document"`
	DocumentId   int       `json:"-"`
	PartnerCount int       `json:"partnerCount"` // Number of members that should sign the document
	Completed    bool      `json:"completed"`    // If all the members have signed the document Completed = true
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

package model

type Notary struct {
	Id           int  `json:"id"`
	UserId       int  `json:"userId"` // owner
	DocumentId   int  `json:"documentId"`
	PartnerCount int  `json:"partnerCount"` // Number of members that should sign the document
	Completed    bool `json:"completed"`    // If all the members have signed the document Completed = true
}

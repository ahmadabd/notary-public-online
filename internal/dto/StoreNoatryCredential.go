package dto

type StoreNoatryCredential struct {
	UserEmail          interface{} `json:"userEmail"`
	DocumentIdempotent string      `json:"documentIdempotent"`
	PartnerCount       int         `json:"partnerCount"`
}

package gorm

import (
	"notary-public-online/internal/entity/model"
	"time"
)

type User struct {
	Id          int       `json:"id" gorm:"primary_key;auto_increment"`
	FirstName   string    `json:"firstName" binding:"required" gorm:"type:varchar(32)"`
	LastName    string    `json:"lastName" binding:"required" gorm:"type:varchar(32)"`
	Email       string    `json:"email" binding:"required,email" gorm:"type:varchar(32)"`
	Password    string    `json:"password" binding:"required" gorm:"type:varchar(32)"`
	Citizenship string    `json:"citizenship" binding:"required" gorm:"type:varchar(32)"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func mapFromUserEntity(user model.User) User {
	return User{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		Citizenship: user.Citizenship,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func MapToUserEntity(user User) model.User {
	return model.User{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		Citizenship: user.Citizenship,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

type Document struct {
	Id          int       `json:"id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name" binding:"required" gorm:"type:varchar(32)"`
	Description string    `json:"description" binding:"required" gorm:"type:varchar(255)"`
	FileAddress string    `json:"fileAddress" gorm:"type:varchar(255)"`
	User        User      `json:"user" gorm:"foreignkey:UserId"`
	UserId      int       `json:"-"`
	Hash        []byte    `json:"hash" gorm:"type:varchar(255);UNIQUE"`
	Active      bool      `json:"active" gorm:"type:bool;default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func mapFromDocumentEntity(document model.Document) Document {
	return Document{
		Id:          document.Id,
		Name:        document.Name,
		Description: document.Description,
		FileAddress: document.FileAddress,
		User:        mapFromUserEntity(document.User),
		UserId:      document.UserId,
		Hash:        document.Hash,
		Active:      document.Active,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
	}
}

func mapToDocumentEntity(document Document) model.Document {
	return model.Document{
		Id:          document.Id,
		Name:        document.Name,
		Description: document.Description,
		FileAddress: document.FileAddress,
		UserId:      document.UserId,
		Hash:        document.Hash,
		Active:      document.Active,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
	}
}

type Notary struct {
	Id           int       `json:"id" gorm:"primary_key;auto_increment"`
	User         User      `json:"user" gorm:"foreignKey:UserId"`
	UserId       int       `json:"-"`
	Document     Document  `json:"document" gorm:"foreignKey:DocumentId"`
	DocumentId   int       `json:"-"`
	PartnerCount int       `json:"partnerCount" binding:"required" gorm:"type:int"`
	Completed    bool      `json:"completed" gorm:"type:bool;default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func mapFromNotaryEntity(notary model.Notary) Notary {
	return Notary{
		Id:           notary.Id,
		User:         mapFromUserEntity(notary.User),
		UserId:       notary.UserId,
		Document:     mapFromDocumentEntity(notary.Document),
		DocumentId:   notary.DocumentId,
		PartnerCount: notary.PartnerCount,
		Completed:    notary.Completed,
		CreatedAt:    notary.CreatedAt,
		UpdatedAt:    notary.UpdatedAt,
	}
}

func mapToNotaryEntity(notary Notary) model.Notary {
	return model.Notary{
		Id:           notary.Id,
		UserId:       notary.UserId,
		DocumentId:   notary.DocumentId,
		PartnerCount: notary.PartnerCount,
		Completed:    notary.Completed,
		CreatedAt:    notary.CreatedAt,
		UpdatedAt:    notary.UpdatedAt,
	}
}

type Signature struct {
	Id             int       `json:"id" gorm:"primary_key;auto_increment"`
	User           User      `json:"user" gorm:"foreignKey:UserId"`
	UserId         int       `json:"-"`
	SignedDocument []byte    `json:"signedDocument" gorm:"type:varchar(255)"`
	Notary         Notary    `json:"notary" gorm:"foreignKey:NoatryId"`
	NotaryId       int       `json:"-"`
	CreatedAt      time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func mapFromSignatureEntity(signature model.Signature) Signature {
	return Signature{
		Id:             signature.Id,
		User:           mapFromUserEntity(signature.User),
		UserId:         signature.UserId,
		SignedDocument: signature.SignedDocument,
		Notary:         mapFromNotaryEntity(signature.Notary),
		NotaryId:       signature.NotaryId,
		CreatedAt:      signature.CreatedAt,
		UpdatedAt:      signature.UpdatedAt,
	}
}

func mapToSignatureEntity(signature Signature) model.Signature {
	return model.Signature{
		Id:             signature.Id,
		UserId:         signature.UserId,
		SignedDocument: signature.SignedDocument,
		NotaryId:       signature.NotaryId,
		CreatedAt:      signature.CreatedAt,
		UpdatedAt:      signature.UpdatedAt,
	}
}

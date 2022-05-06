package gorm

import (
	"context"
	"notary-public-online/internal/entity/model"
)

func (db *Gorm) CreateSignature(ctx context.Context, noatryId int, userId int, documentSignature *[]byte) error {

	signature := mapFromSignatureEntity(model.Signature{
		NotaryId:       noatryId,
		SignedDocument: *documentSignature,
		UserId:         userId,
	})

	if err := db.Db.WithContext(ctx).Create(&signature).Error; err != nil {
		return err
	}

	return nil
}

func (db *Gorm) GetSignatures(ctx context.Context, noatryId int, userId int) (*[]byte, error) {

	signature := &model.Signature{}

	if err := db.Db.WithContext(ctx).First(signature, "notary_id = ? AND user_id = ?", noatryId, userId).Error; err != nil {
		return nil, err
	}

	return &signature.SignedDocument, nil
}

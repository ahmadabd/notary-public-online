package gorm

import (
	"context"
	"notary-public-online/internal/entity/model"
)

func (db *Gorm) CreateNoatry(ctx context.Context, documentId int, userId int, partnerCount int, completed bool) error {

	noatry := mapFromNotaryEntity(model.Notary{
		UserId:       userId,
		DocumentId:   documentId,
		PartnerCount: partnerCount,
		Completed:    completed,
	})

	if err := db.Db.WithContext(ctx).Create(&noatry).Error; err != nil {
		return err
	}

	return nil
}

func (db *Gorm) GetNoatry(ctx context.Context, noatryId int) (model.Notary, error) {

	var notary Notary

	if err := db.Db.WithContext(ctx).Where("Id", noatryId).First(notary).Error; err != nil {
		return model.Notary{}, err
	}

	return mapToNotaryEntity(notary), nil

}

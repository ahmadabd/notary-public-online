package gorm

import (
	"context"
	"notary-public-online/internal/entity/model"
)

func (db *Gorm) CreateDocument(ctx context.Context, name string, description string, fileAddress string, documentsHash *[]byte, userId int, active bool) (model.Document, error) {
	document := mapFromDocumentEntity(model.Document{
		Name:        name,
		Description: description,
		FileAddress: fileAddress,
		Hash:        *documentsHash,
		UserId:      userId,
		Active:      active,
	})

	if err := db.Db.WithContext(ctx).Create(&document).Error; err != nil {
		return model.Document{}, err
	}

	return mapToDocumentEntity(document), nil
}

func (db *Gorm) GetDocument(ctx context.Context, documentId int) (model.Document, error) {
	var document Document

	if err := db.Db.WithContext(ctx).Where("id", documentId).First(&document).Error; err != nil {
		return model.Document{}, err
	}

	return mapToDocumentEntity(document), nil
}

func (db *Gorm) GetDocumentAddress(ctx context.Context, documentId int) (string, error) {
	var document Document

	if err := db.Db.WithContext(ctx).Select("FileAddress").Where("id", documentId).First(&document).Error; err != nil {
		return "", err
	}

	return document.FileAddress, nil
}

func (db *Gorm) GetDocumentHash(ctx context.Context, documentId int) (*[]byte, error) {
	var document Document

	if err := db.Db.WithContext(ctx).Select("Hash").Where("id", documentId).First(&document).Error; err != nil {
		return nil, err
	}

	return &document.Hash, nil
}

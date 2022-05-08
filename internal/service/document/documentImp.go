package document

import (
	"context"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/pkg/hash/sha256"
	"notary-public-online/internal/pkg/storage"
	"notary-public-online/internal/repository"
	"os"
)

type documentImp struct {
	Db      repository.DB
	Storage storage.Storage
}

func New(db repository.DB, storage storage.Storage) Document {
	return &documentImp{Db: db, Storage: storage}
}

func (d *documentImp) StoreDocument(ctx context.Context, document *os.File, name string, description string, userId int) (model.Document, error) {

	// get document hash
	hasher := sha256.New()
	documentsHash, err := hasher.Hash(document)
	if err != nil {
		return model.Document{}, err
	}

	// store document in storage
	fileAddress, err := d.Storage.StoreFile(document)
	if err != nil {
		return model.Document{}, err
	}

	return d.Db.CreateDocument(ctx, name, description, fileAddress, &documentsHash, userId, false)
}

func (d *documentImp) DocumentDetails(ctx context.Context, documentId int) (model.Document, error) {
	return d.Db.GetDocument(ctx, documentId)
}

func (d *documentImp) ReadDocument(ctx context.Context, documentId int) (*os.File, error) {

	document, err := d.Db.GetDocumentAddress(ctx, documentId)

	if err != nil {
		return nil, err
	}

	return d.Storage.ReadFile(document)
}

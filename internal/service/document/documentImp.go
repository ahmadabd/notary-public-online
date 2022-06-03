package document

import (
	"context"
	"log"
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

func (d *documentImp) StoreDocument(ctx context.Context, idempotentKey string, document *os.File, name string, description string, userEmail string) error {

	// get user id
	user, err := d.Db.GetUserWithEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	// check idempotency
	if d.Db.CheckDocumentIdempotency(ctx, idempotentKey) != 0 {
		log.Println("Item with this idempotent key already exists")
		return nil
	}

	// get document hash
	hasher := sha256.New()
	documentsHash, err := hasher.Hash(document)
	if err != nil {
		return err
	}

	// store document in storage
	// TODO: store file or encrypted file
	fileAddress, err := d.Storage.StoreFile(document)
	if err != nil {
		return err
	}

	return d.Db.CreateDocument(ctx, idempotentKey, name, description, fileAddress, &documentsHash, user.Id, false)
}

func (d *documentImp) DocumentDetails(ctx context.Context, documentId int) (model.Document, error) {
	return d.Db.GetDocument(ctx, documentId)
}

func (d *documentImp) ReadDocument(ctx context.Context, idempotent string) (*os.File, error) {

	documentAddr, err := d.Db.GetDocumentAddress(ctx, idempotent)

	if err != nil {
		return nil, err
	}

	return d.Storage.ReadFile(documentAddr)
}

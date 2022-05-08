package repository

import (
	"context"
	"notary-public-online/internal/entity/model"
)

type document interface {
	CreateDocument(ctx context.Context, name string, description string, fileAddress string, documentsHash *[]byte, userId int, active bool) (model.Document, error)

	GetDocument(ctx context.Context, documentId int) (model.Document, error)

	GetDocumentAddress(ctx context.Context, documentId int) (string, error)

	GetDocumentHash(ctx context.Context, documentId int) (*[]byte, error)
}

type notary interface {
	CreateNoatry(ctx context.Context, documentId int, userId int, partnerCount int, completed bool) (model.Notary, error)

	GetNoatry(ctx context.Context, noatryId int) (model.Notary, error)
}

type signature interface {
	CreateSignature(ctx context.Context, noatryId int, userId int, documentSignature *[]byte) (model.Signature, error)

	GetSignatures(ctx context.Context, noatryId int, userId int) (*[]byte, error)
}

type user interface {
	GetUserKeys(ctx context.Context, userId int) (string, string, error)
}

type DB interface {
	document
	notary
	user
	signature
}

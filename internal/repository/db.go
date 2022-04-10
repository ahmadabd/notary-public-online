package repository

import "context"

type document interface {
	CreateDocument(ctx context.Context, document string, documentsHash string, user int) error

	GetDocumentHash(ctx context.Context, documentId int) (string, error)
}

type user interface {
	GetUserKeys(ctx context.Context, userId int) (string, string, error)
}


type DB interface {
	document
	user
}
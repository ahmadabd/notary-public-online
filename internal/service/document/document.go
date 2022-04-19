package document

import (
	"context"
	"notary-public-online/internal/entity/model"
	"os"
)

type Document interface {
	StoreDocument(ctx context.Context, document *os.File, name string, description string, userId int) error

	DocumentDetails(ctx context.Context, documentId int) (*model.Document, error)

	ReadDocument(ctx context.Context, documentId int) (*os.File, error)
}
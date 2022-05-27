package document

import (
	"context"
	"notary-public-online/internal/entity/model"
	"os"
)

type Document interface {
	StoreDocument(ctx context.Context, idempotentKey string, document *os.File, name string, description string, userEmail string) error

	DocumentDetails(ctx context.Context, documentId int) (model.Document, error)

	ReadDocument(ctx context.Context, documentId int) (*os.File, error)
}

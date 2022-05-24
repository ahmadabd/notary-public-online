package document

import (
	"context"
	"notary-public-online/internal/entity/model"
	"os"
)

type Document interface {
	StoreDocument(ctx context.Context, document *os.File, name string, description string, userEmail string) (model.Document, error)

	DocumentDetails(ctx context.Context, documentId int) (model.Document, error)

	ReadDocument(ctx context.Context, documentId int) (*os.File, error)
}

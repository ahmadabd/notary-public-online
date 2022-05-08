package noatry

import (
	"context"
	"notary-public-online/internal/entity/model"
)

type Noatry interface {
	CreateNoatry(ctx context.Context, documentId int, userId int, partnerCount int, completed bool) (model.Notary, error)

	GetNoatry(ctx context.Context, noatryId int) (model.Notary, error)

	SignNoatry(ctx context.Context, noatryId int, userId int) error

	VerifyNoatrySignature(ctx context.Context, userId int, noatryId int) (bool, error)

	// DocumentHashEncrypt(ctx context.Context, signedDocument string) (string, error)

	// DocumentHashDecrypt(ctx context.Context, encryptedDocument string) (string, error)
}

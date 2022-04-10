package document

import "context"

type DocumentImp interface {
	DocumentHash(ctx context.Context, document string, userId int) error

	DocumentSignature(ctx context.Context, documentId int, userId int) (string, error)

	DocumentSignatureVerify(ctx context.Context, documentHash string, signature string) (bool, error)

	DocumentHashEncrypt(ctx context.Context, signedDocument string) (string, error)

	DocumentHashDecrypt(ctx context.Context, encryptedDocument string) (string, error)
}
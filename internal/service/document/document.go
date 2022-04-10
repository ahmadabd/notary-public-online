package document

import (
	"context"
	"errors"
	"notary-public-online/internal/pkg/hash/sha256"
	"notary-public-online/internal/pkg/pairKey/rsa"
	"notary-public-online/internal/repository"
)

type doc struct {
	Db repository.DB
}

func New(db repository.DB) DocumentImp {
	return &doc{Db: db}
}

func (d *doc) DocumentHash(ctx context.Context, document string, userId int) error {

	hasher := sha256.New()
	documentsHash := hasher.Hash(document)

	return d.Db.CreateDocument(ctx, document, documentsHash, userId)
}

func (d *doc) DocumentSignature(ctx context.Context, documentId int, userId int) (string, error) {

	documentHash, err := d.Db.GetDocumentHash(ctx, documentId)

	if err != nil {
		return "", errors.New("document not found")
	}

	privateKey, publicKey, err := d.Db.GetUserKeys(ctx, userId)

	if err != nil {
		return "", errors.New("user not found")
	}

	crypto := rsa.New(privateKey, publicKey)
	if signedDoc, err := crypto.Signature(&documentHash); err == nil {
		return signedDoc, nil
	}

	return "", errors.New("error signing document")
}

func (d *doc) DocumentSignatureVerify(ctx context.Context, documentHash string, signature string) (bool, error) {
	return true, nil
}

func (d *doc) DocumentHashEncrypt(ctx context.Context, signedDocument string) (string, error) {
	return "", nil
}

func (d *doc) DocumentHashDecrypt(ctx context.Context, encryptedDocument string) (string, error) {
	return "", nil
}

package noatry

import (
	"context"
	"errors"
	"notary-public-online/internal/entity/model"
	"notary-public-online/internal/pkg/pairKey"
	"notary-public-online/internal/pkg/pairKey/rsa"
	"notary-public-online/internal/repository"
)

type doc struct {
	Db repository.DB
}

func New(db repository.DB) Noatry {
	return &doc{Db: db}
}

func (d *doc) CreateNoatry(ctx context.Context, documentId int, userId int, partnerCount int, completed bool) (model.Notary, error) {
	return d.Db.CreateNoatry(ctx, documentId, userId, partnerCount, completed)
}

func (d *doc) GetNoatry(ctx context.Context, noatryId int) (model.Notary, error) {
	return d.Db.GetNoatry(ctx, noatryId)
}

func (d *doc) SignNoatry(ctx context.Context, noatryId int, userId int) error {

	noatry, err := d.Db.GetNoatry(ctx, noatryId)
	if err != nil {
		return errors.New("noatry not found")
	}

	documentHash, err := d.Db.GetDocumentHash(ctx, noatry.DocumentId)
	if err != nil {
		return errors.New("document not found")
	}

	crypto, err := getUserCrypto(ctx, d, userId)
	if err != nil {
		return errors.New("error in getting user crypto object")
	}

	if signedDoc, err := crypto.Signature(documentHash); err == nil {
		if _, err := d.Db.CreateSignature(ctx, noatryId, userId, &signedDoc); err == nil {
			return err
		}
	}

	return errors.New("error in signing document")
}

func (d *doc) VerifyNoatrySignature(ctx context.Context, userId int, noatryId int) (bool, error) {

	noatry, err := d.Db.GetNoatry(ctx, noatryId)
	if err != nil {
		return false, errors.New("noatry not found")
	}

	documentId := noatry.DocumentId

	crypto, err := getUserCrypto(ctx, d, userId)
	if err != nil {
		return false, errors.New("error in getting user crypto object")
	}

	documentHash, err := d.Db.GetDocumentHash(ctx, documentId)
	if err != nil {
		return false, errors.New("error in getting document hash")
	}

	documentSignature, err := d.Db.GetSignatures(ctx, noatryId, userId)

	if err != nil {
		return false, errors.New("error in getting document signature")
	}

	// verify := crypto.VerifySignature(documentHash, documentSignature)
	verify := crypto.VerifySignature(documentSignature, documentHash)

	return verify, nil
}

// func (d *doc) DocumentHashEncrypt(ctx context.Context, signedDocument string) (string, error) {
// 	return "", nil
// }

// func (d *doc) DocumentHashDecrypt(ctx context.Context, encryptedDocument string) (string, error) {
// 	return "", nil
// }

func getUserCrypto(ctx context.Context, d *doc, userId int) (pairKey.Crypto, error) {
	// privateKey, publicKey, err := d.Db.GetUserKeys(ctx, userId)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return rsa.New(privateKey, publicKey), nil
}

package document

type DocumentImp interface {
	DocumentHash(document string) (string, error)

	DocumentSignature(documentHash string) (string, error)

	DocumentSignatureVerify(documentHash string, signature string) (bool, error)

	DocumentHashEncrypt(signedDocument string) (string, error)

	DocumentHashDecrypt(encryptedDocument string) (string, error)
}
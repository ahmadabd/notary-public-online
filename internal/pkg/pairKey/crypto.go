package pairKey

type Crypto interface {
	Signature([]byte) ([]byte, error) // Sign a document using pr

	VerifySignature(*[]byte, []byte) bool // Verify a signed document using third party pu for check that the document was signed using the third party or not, then Usnign the document to make sure the documents hash has not changed

	Encryption(*string) (string, error) // Sign document using third party pu to send

	Decryption(*string) (string, error) // Unsign secured document using pr
}

//go:generate $HOME/go_projects/bin/mockgen -destination=../../../mocks/mock_keys.go -package=mocks notary-public-online/internal/pkg/pairKey Keys
type Keys interface {
	PairKeyGenerator(email string) ([]byte, []byte, error)

	PairKeyReader(filename string) ([]byte, []byte, error)
}

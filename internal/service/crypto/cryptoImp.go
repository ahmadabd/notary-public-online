package crypto

type CryptoImp interface {
	PairKeyGenerator() (string, string, error)			// Generate a pair of keys
	Sign(string, string) (string, error)				// Sign a document using pr
	UnSign(string, string) (string, error)				// UnSign a signed document using pu, its one of the steps before doing hash-check
	Verify(string, string, string) (bool, error)		// Verify a signed document using third party pu for check that the document was signed using the third party or not, then Usnign the document to make sure the documents hash has not changed
	MakeSecure(string) (string, error)					// Sign document using third party pu to send
	Decoder(string) (string, error)						// Unsign secured document using pr
}
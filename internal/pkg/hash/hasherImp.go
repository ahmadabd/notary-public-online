package hash

type HasherImp interface {
	Hash(string) string
	
	HashChecker(string, string) bool
}
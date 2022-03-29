package hash

type HasherImp interface {
	hasher(string) (string, error)
	hashChecker(string, string) (bool, error)
}
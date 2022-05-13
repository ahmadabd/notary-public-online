package passwordHash

//go:generate $HOME/go_projects/bin/mockgen -destination=../../../mocks/mock_passhash.go -package=mocks notary-public-online/internal/pkg/passwordHash PasswordHash
type PasswordHash interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}

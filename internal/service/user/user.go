package user

type UserImp interface {
	Register(personalInfo []string, username string, password string) (string, error)

	Login(username string, password string) (string, error)
}
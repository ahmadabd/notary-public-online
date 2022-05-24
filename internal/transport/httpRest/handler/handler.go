package handler

import (
	"notary-public-online/internal/pkg/jwtPkg"
	"notary-public-online/internal/service/document"
	"notary-public-online/internal/service/user"
)

type handler struct {
	userServ user.User
	docServ  document.Document
	jwtPkg   jwtPkg.Jwt
}

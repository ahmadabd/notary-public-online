package handler

import (
	"context"
	"notary-public-online/internal/dto"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterController(ctx *gin.Context) error {
	cctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user dto.RegisterCredential
	ctx.ShouldBindJSON(&user)

	return h.userServ.Register(cctx, user)
}

func (h *handler) LoginController(ctx *gin.Context) string {
	cctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user dto.LoginCredential
	ctx.ShouldBindJSON(&user)

	result, err := h.userServ.Login(cctx, user)
	if err != nil || !result {
		return ""
	}

	return h.jwtPkg.GenerateToken(user.Email, false)
}

package handler

import (
	"context"
	"notary-public-online/internal/dto"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) StoreNoatry(ctx *gin.Context) error {
	cctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userEmail, _ := ctx.Get("user")

	var storeNoatryCredential dto.StoreNoatryCredential

	ctx.ShouldBindJSON(&storeNoatryCredential)
	storeNoatryCredential.UserEmail = userEmail

	return h.noatryServ.CreateNoatry(cctx, &storeNoatryCredential)
}

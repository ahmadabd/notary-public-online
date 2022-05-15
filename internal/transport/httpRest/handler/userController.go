package handler

import (
	"context"
	"notary-public-online/internal/entity/model"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) Register(ctx *gin.Context) (model.User, error) {
	cctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.User
	ctx.BindJSON(&user)

	return h.userServ.Register(cctx, user)
}

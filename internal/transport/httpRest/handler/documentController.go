package handler

import (
	"context"
	"fmt"
	"notary-public-online/internal/dto"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) StoreDocument(ctx *gin.Context) error {

	cctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userEmail, exists := ctx.Get("user")

	if exists {

		var document dto.StoreDocumentCredential

		if err := ctx.ShouldBindJSON(&document); err != nil {
			return err
		}

		h.docServ.StoreDocument(cctx, document.Document, document.Name, document.Description, fmt.Sprintf("%v", userEmail))
	}
	
	return nil
}

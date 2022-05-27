package handler

import (
	"context"
	"fmt"
	"notary-public-online/internal/dto"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) StoreDocument(ctx *gin.Context) error {

	cctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check idempotency
	idempotentKey := ctx.GetHeader("idempotent")
	if idempotentKey == "" {
		return fmt.Errorf("missing idempotency key")
	}

	userEmail, exists := ctx.Get("user")

	if exists {

		var document dto.StoreDocumentCredential

		file, err := handleFileUpload(ctx)

		if err != nil {
			return err
		}

		document.Document = file
		document.Name = ctx.PostForm("name")
		document.Description = ctx.PostForm("description")

		err = h.docServ.StoreDocument(cctx, idempotentKey, document.Document, document.Name, document.Description, fmt.Sprintf("%v", userEmail))

		if err != nil {
			return err
		}
	}

	return nil
}

func handleFileUpload(r *gin.Context) (*os.File, error) {

	file, err := r.FormFile("document")
	if err != nil {
		return nil, err
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := r.SaveUploadedFile(file, "/tmp/"+newFileName); err != nil {
		return nil, err
	}

	// Open the file
	f, err := os.Open("/tmp/" + newFileName)
	if err != nil {
		return nil, err
	}

	return f, nil
}

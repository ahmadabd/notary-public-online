package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (r *rest) routes(baseRoute string) {
	r.gin.GET(fmt.Sprintf("/api/%s", baseRoute), func(ctx *gin.Context) {
		ctx.JSON(200, r.handler.GetDocument())
	})
}
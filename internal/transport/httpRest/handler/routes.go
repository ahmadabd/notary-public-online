package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) routes(baseRoute string) {
	user := r.gin.Group(fmt.Sprintf("/%s/user", baseRoute))
	{
		user.POST("/register", func(ctx *gin.Context) {
			user, err := r.handler.Register(ctx)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "success",
					"user":    user,
				})
			}
		})
	}
}

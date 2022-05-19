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
			err := r.handler.RegisterController(ctx)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "success",
				})
			}
		})

		user.POST("/login", func(ctx *gin.Context) {
			token := r.handler.LoginController(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "unauthorized",
				})
			}
		})
	}

	// api := r.gin.Group(fmt.Sprint("/%s/api", baseRoute), middlewares.AuthorizeJWT())
	// {

	// }
}

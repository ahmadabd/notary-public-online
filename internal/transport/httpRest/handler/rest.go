package handler

import (
	"fmt"
	"notary-public-online/internal/configs/yaml"
	"notary-public-online/internal/service/user"
	"notary-public-online/internal/transport/httpRest"

	"github.com/gin-gonic/gin"
)

type rest struct {
	gin     *gin.Engine
	handler *handler
}

func New(userServ user.User) httpRest.Rest {
	return &rest{
		gin: gin.Default(),
		handler: &handler{
			userServ: userServ,
		},
	}
}

func (r *rest) Start(cfg *yaml.Config) error {

	r.routes("api")

	return r.gin.Run(serverConfig(cfg))
}

// func (r *rest) Shutdown() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

// 	defer cancel()

// 	return r.gin.
// }

func serverConfig(cfg *yaml.Config) string {
	return fmt.Sprintf(":%s", cfg.Server.Port)
}

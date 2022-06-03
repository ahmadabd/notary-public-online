package cmd

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"notary-public-online/internal/configs/yaml"
	"notary-public-online/internal/pkg/jwtPkg"
	"notary-public-online/internal/pkg/pairKey/rsa"
	"notary-public-online/internal/pkg/passwordHash/passbcrypt"
	"notary-public-online/internal/pkg/storage/disk"
	"notary-public-online/internal/repository/gorm"
	"notary-public-online/internal/service/document"
	"notary-public-online/internal/service/noatry"
	"notary-public-online/internal/service/user"
	"notary-public-online/internal/transport/httpRest/handler"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var ServeCMD = &cli.Command{
	Name:    "notary",
	Aliases: []string{"n"},
	Usage:   "Start the notary server",
	Action:  serve,
}

func setupLoggingOutPut() {
	f, _ := os.Create(filepath.Join("logs", fmt.Sprintf("%s-%d.log", "notary", rand.Intn(10000))))

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func serve(c *cli.Context) error {

	cfg, err := yaml.GetConfig("config.yaml")

	if err != nil {
		log.Println(err)
	}

	setupLoggingOutPut()
	gin.SetMode(cfg.App.Mode)

	db, err := gorm.SetupDatabase(cfg)
	if err != nil {
		log.Println("error while connecting to database: ", err)
	}

	jwtPkg := jwtPkg.JWTAuthService()

	documentService := document.New(db, disk.New())
	noatryService := noatry.New(db)
	userService := user.New(db, rsa.NewKeys(), passbcrypt.New())

	go func() {
		if err := handler.New(userService, documentService, noatryService, jwtPkg).Start(cfg); err != nil {
			log.Println(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	log.Println("Recived SIGTERM, exiting gracefully...")

	return nil
}

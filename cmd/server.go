package cmd

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"notary-public-online/internal/configs/yaml"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var ServeCMD = &cli.Command{
	Name:    "notary",
	Aliases: []string{"n"},
	Usage:   "Start the notary server",
	Action:  serve,
}

func setupLogOutPut() {
	f, _ := os.Create(filepath.Join("logs", fmt.Sprintf("%s.log", "notary-"+string(rand.Intn(10000)))))

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func serve(c *cli.Context) error {
	setupLogOutPut()

	cfg, err := yaml.GetConfig("config.yml")

	if err != nil {
		log.Println(err)
	}

	
}

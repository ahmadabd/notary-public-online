package main

import (
	"log"
	"notary-public-online/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
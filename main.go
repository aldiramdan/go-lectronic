package main

import (
	"lectronic/src/configs"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := configs.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/bereket-7/gin-template/internal/app"
)

func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

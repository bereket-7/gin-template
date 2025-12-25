package main

// @title           Gin Template API
// @version         1.0
// @description     Production-ready Gin starter template
// @termsOfService  https://github.com/<your-github-username>/gin-template

// @contact.name   Your Name
// @contact.url    https://github.com/<your-github-username>
// @contact.email  your@email.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

import (
	"log"

	_ "github.com/bereket-7/gin-template/docs"
	"github.com/bereket-7/gin-template/internal/app"
)

func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

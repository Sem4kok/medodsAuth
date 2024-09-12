package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"medodsAuth/internal/controller/users"
)

const (
	localhost = ":8081"
)

type App struct {
	*gin.Engine
}

func New() *App {
	return &App{
		gin.Default(),
	}
}

func (app *App) MustRunApp() {
	app.handleUrls()
	log.Fatal(app.Run(localhost))
}

func (app *App) handleUrls() {
	app.POST("/api/register", users.Register)
	app.POST("/api/login", users.Login)
}

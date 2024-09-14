package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"medodsAuth/internal/controller/tokens"
	"medodsAuth/internal/controller/users"
	storage "medodsAuth/internal/storage/postgresql"
	"net/http"
	"os/signal"
	"syscall"
)

const (
	localhost   = ":8081"
	storagePath = ""
)

type App struct {
	*gin.Engine
}

func New() *App {
	return &App{
		Engine: gin.Default(),
	}
}

func (app *App) MustRunApp() {
	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGABRT, syscall.SIGKILL,
		syscall.SIGALRM)
	defer stop()

	server := http.Server{
		Handler: app.Engine,
	}

	storage.ConnectDB(storagePath)
	app.handleUrls()

	go func() {
		log.Fatal(app.Run(localhost))
	}()

	<-quit.Done()
	log.Println("receive interrupt signal")

	server.Close()
}

func (app *App) handleUrls() {
	app.POST("/api/register", users.Register)
	app.POST("/api/token/get", tokens.GetTokens)
	app.POST("/api/token/refresh", tokens.RefreshTokens)
}

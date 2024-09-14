package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"medodsAuth/internal/controller/tokens"
	"medodsAuth/internal/controller/users"
	storage "medodsAuth/internal/storage/postgresql"
	"net/http"
	"os/signal"
	"syscall"
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

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: app.Engine,
	}

	go func() {
		storage.ConnectDB(viper.GetString("storagePath"))
		app.handleUrls()
		log.Fatal(app.Run(viper.GetString("port")))
	}()

	<-quit.Done()
	log.Println("receive interrupt signal")

	server.Close()
}

func (app *App) handleUrls() {
	app.POST("/api/register", users.Register)
	app.GET("/api/token/get", tokens.GetTokens)
	app.POST("/api/token/refresh", tokens.RefreshTokens)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

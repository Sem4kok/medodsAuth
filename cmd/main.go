package main

import (
	"medodsAuth/internal/app"
	storage "medodsAuth/internal/storage/postgresql"
)

func main() {
	application := app.New()
	application.MustRunApp()
	defer storage.DB.Close()
}

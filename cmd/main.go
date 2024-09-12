package main

import "medodsAuth/internal/app"

func main() {
	application := app.New()
	application.MustRunApp()
}

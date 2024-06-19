package main

import (
	"log"

	"khand.dev/msngr/cmd/app"
)

func main() {
	app := app.RootCmd()
	if err := app.Execute(); err != nil {
		log.Fatalf("unhandled exception, %s", err.Error())
	}
}

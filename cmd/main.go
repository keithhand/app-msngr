package main

import (
	"log"

	"khand.dev/msngr/cmd/app"
)

func main() {
	if err := app.RootCmd(); err != nil {
		log.Fatalf("unhandled exception, %s", err.Error())
	}
}

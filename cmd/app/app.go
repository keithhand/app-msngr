package main

import (
	"log"

	"khand.dev/msngr/pkg/cmd/app"
)

func main() {
	if err := app.Execute(); err != nil {
		log.Fatalf("unhandled exception, %s", err.Error())
	}
}

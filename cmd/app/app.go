package main

import (
	"os"

	"khand.dev/msngr/pkg/cmd/app"
)

func main() {
	cli := app.AppCmd()
	if err := cli.Run(); err != nil {
		os.Exit(1)
	}
}

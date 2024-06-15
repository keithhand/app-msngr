package main

import (
	"fmt"
	"os"

	"khand.dev/msngr/cmd/app/cmd"
)

func main() {
	cli := cmd.MainCmd()
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
	}
}

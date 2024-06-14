package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:   "app",
		Short: "An easier way to talk with Kubernetes",
		Long: `An easier way to talk with Kubernetes. 

Declare an application spec, for access to common Kubernetes tasks
without needing to know deployment specifics.
	`,
	}

	create := &cobra.Command{
		Use:   "create",
		Short: "Install an application.",
	}
	deploy := &cobra.Command{
		Use:   "deploy",
		Short: "Update an application.",
	}
	logs := &cobra.Command{
		Use:   "logs",
		Short: "Retrieve logs for the defined container set.",
	}
	reinit := &cobra.Command{
		Use:   "re-init",
		Short: "Uninstall and reinstall an application.",
	}
	destroy := &cobra.Command{
		Use:   "destroy",
		Short: "Uninstall an application.",
	}

	app.AddCommand(create, deploy, logs, reinit, destroy)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
	}
}

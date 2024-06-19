package app

import "github.com/spf13/cobra"

type Command = cobra.Command

var (
	slogan = "Keep the talking with your applications short"
	appCmd = &Command{
		Use:   "app",
		Short: slogan,
		Long: slogan + ".\n" + `
 Reads an application spec for access to common Kubernetes
 tasks without needing to know deployment specifics.`,
	}
)

func RootCmd() *Command {
	return appCmd
}

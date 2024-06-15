package app

import "khand.dev/msngr/pkg/cmd"

var (
	slogan = "Keep the talking with your applications short"
	appCmd = cmd.New(cmd.Config{
		Verb:      "app",
		ShortDesc: slogan,
		LongDesc: slogan + ".\n" + `
 Reads an application spec for access to common Kubernetes
 tasks without needing to know deployment specifics.`,
	})
)

func AppCmd() cmd.Command {
	return appCmd
}

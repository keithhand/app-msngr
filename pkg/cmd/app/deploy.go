package app

import "khand.dev/msngr/pkg/cmd"

var deploy = cmd.NewCommand(cmd.Config{
	Verb:      "deploy",
	ShortDesc: "Update an application.",
})

func init() {
	appCmd.AttachCommands(deploy)
}

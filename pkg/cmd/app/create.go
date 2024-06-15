package app

import "khand.dev/msngr/pkg/cmd"

var create = cmd.NewCommand(cmd.Config{
	Verb:      "create",
	ShortDesc: "Install an application.",
})

func init() {
	appCmd.AttachCommands(create)
}

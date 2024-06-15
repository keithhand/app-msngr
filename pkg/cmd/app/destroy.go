package app

import "khand.dev/msngr/pkg/cmd"

var destroy = cmd.NewCommand(cmd.Config{
	Verb:      "destroy",
	ShortDesc: "Uninstall an application.",
})

func init() {
	appCmd.AttachCommands(destroy)
}

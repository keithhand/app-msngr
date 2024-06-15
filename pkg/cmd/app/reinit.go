package app

import "khand.dev/msngr/pkg/cmd"

var reinit = cmd.NewCommand(cmd.Config{
	Verb:      "re-init",
	ShortDesc: "Uninstall and reinstall an application.",
})

func init() {
	appCmd.AttachCommands(reinit)
}

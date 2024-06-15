package app

import "khand.dev/msngr/pkg/cmd"

var destroy = cmd.New(cmd.Config{
	Verb:      "destroy",
	ShortDesc: "Uninstall an application.",
})

func init() {
	appCmd.AddCommand(destroy)
}

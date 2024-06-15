package app

import "khand.dev/msngr/pkg/cmd"

var create = cmd.New(cmd.Config{
	Verb:      "create",
	ShortDesc: "Install an application.",
})

func init() {
	appCmd.AddCommand(create)
}

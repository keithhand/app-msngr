package app

import "khand.dev/msngr/pkg/cmd"

var deploy = cmd.New(cmd.Config{
	Verb:      "deploy",
	ShortDesc: "Update an application.",
})

func init() {
	appCmd.AddCommand(deploy)
}

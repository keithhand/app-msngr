package app

import "khand.dev/msngr/pkg/cmd"

var logs = cmd.NewCommand(cmd.Config{
	Verb:      "logs",
	ShortDesc: "Retrieve logs for the defined container set.",
})

func init() {
	appCmd.AttachCommands(logs)
}

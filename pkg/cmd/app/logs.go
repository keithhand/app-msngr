package app

import "khand.dev/msngr/pkg/cmd"

var logs = cmd.New(cmd.Config{
	Verb:      "logs",
	ShortDesc: "Retrieve logs for the defined container set.",
})

func init() {
	appCmd.AddCommand(logs)
}

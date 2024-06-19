package app

import "khand.dev/msngr/pkg/kubernetes"

var logs = &Command{
	Use:   "logs",
	Short: "Retrieve logs for the defined container set.",
	Run: func(cmd *Command, args []string) {
		kubernetes.Logs()
	},
}

func init() {
	appCmd.AddCommand(logs)
}

package app

import "log"

var logs = &Command{
	Use:   "logs",
	Short: "Retrieve logs for the defined container set.",
	Run:   run,
}

func run(cmd *Command, args []string) {
	log.Printf("ahh: ")
}

func init() {
	appCmd.AddCommand(logs)
}

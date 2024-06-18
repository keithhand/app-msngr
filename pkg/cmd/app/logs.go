package app

var logs = &Command{
	Use:   "logs",
	Short: "Retrieve logs for the defined container set.",
}

func init() {
	appCmd.AddCommand(logs)
}

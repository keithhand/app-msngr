package cmd

var logs = &command{
	Use:   "logs",
	Short: "Retrieve logs for the defined container set.",
}

func init() {
	rootCmd.AddCommand(logs)
}

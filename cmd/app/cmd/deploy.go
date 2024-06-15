package cmd

var deploy = &command{
	Use:   "deploy",
	Short: "Update an application.",
}

func init() {
	rootCmd.AddCommand(deploy)
}

package cmd

var create = &command{
	Use:   "create",
	Short: "Install an application.",
}

func init() {
	rootCmd.AddCommand(create)
}

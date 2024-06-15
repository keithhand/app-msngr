package cmd

var destroy = &command{
	Use:   "destroy",
	Short: "Uninstall an application.",
}

func init() {
	rootCmd.AddCommand(destroy)
}

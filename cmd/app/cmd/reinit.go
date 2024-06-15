package cmd

var reinit = &command{
	Use:   "re-init",
	Short: "Uninstall and reinstall an application.",
}

func init() {
	rootCmd.AddCommand(reinit)
}

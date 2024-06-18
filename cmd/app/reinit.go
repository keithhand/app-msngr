package app

var reinit = &Command{
	Use:   "re-init",
	Short: "Uninstall and reinstall an application.",
}

func init() {
	appCmd.AddCommand(reinit)
}

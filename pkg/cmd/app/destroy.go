package app

var destroy = &Command{
	Use:   "destroy",
	Short: "Uninstall an application.",
}

func init() {
	appCmd.AddCommand(destroy)
}

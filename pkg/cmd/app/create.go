package app

var create = &Command{
	Use:   "create",
	Short: "Install an application.",
}

func init() {
	appCmd.AddCommand(create)
}

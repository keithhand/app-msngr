package app

var deploy = &Command{
	Use:   "deploy",
	Short: "Update an application.",
}

func init() {
	appCmd.AddCommand(deploy)
}

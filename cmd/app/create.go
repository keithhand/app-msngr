package app

import "khand.dev/msngr/pkg/helm"

var create = &Command{
	Use:   "create",
	Short: "Install an application.",
	Run: func(cmd *Command, args []string) {
		helm.Install()
	},
}

func init() {
	appCmd.AddCommand(create)
}

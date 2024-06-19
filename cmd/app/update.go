package app

import "khand.dev/msngr/pkg/helm"

var deploy = &Command{
	Use:   "update",
	Short: "Update an application.",
	Run: func(cmd *Command, args []string) {
		helm.Update()
	},
}

func init() {
	appCmd.AddCommand(deploy)
}

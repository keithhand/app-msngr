package app

import "khand.dev/msngr/pkg/helm"

var destroy = &Command{
	Use:   "destroy",
	Short: "Uninstall an application.",
	Run: func(cmd *Command, args []string) {
		helm.Uninstall()
	},
}

func init() {
	appCmd.AddCommand(destroy)
}

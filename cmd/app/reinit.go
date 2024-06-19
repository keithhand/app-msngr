package app

import "khand.dev/msngr/pkg/helm"

var reinit = &Command{
	Use:   "re-init",
	Short: "Uninstall and reinstall an application.",
	Run: func(cmd *Command, args []string) {
		helm.Uninstall()
		helm.Install()
	},
}

func init() {
	appCmd.AddCommand(reinit)
}

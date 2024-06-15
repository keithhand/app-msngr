package cmd

import "khand.dev/msngr/pkg/cmd"

type CliSvc struct {
	cmdr cmd.Commander
}

type Cli interface {
	Run() error
}

func MainCmd() Cli {
	svc := &CliSvc{
		cmdr: cmd.CobraCommander(),
	}

	slogan := "Keep the talking with your applications short"
	cmds := svc.cmdr.Command(cmd.Config{
		Verb:      "app",
		ShortDesc: slogan,
		LongDesc: slogan + ".\n" + `
 Reads an application spec for access to common Kubernetes
 tasks without needing to know deployment specifics.`,
	})

	create := svc.cmdr.Command(cmd.Config{
		Verb:      "create",
		ShortDesc: "Install an application.",
	})
	deploy := svc.cmdr.Command(cmd.Config{
		Verb:      "deploy",
		ShortDesc: "Update an application.",
	})
	logs := svc.cmdr.Command(cmd.Config{
		Verb:      "logs",
		ShortDesc: "Retrieve logs for the defined container set.",
	})
	reinit := svc.cmdr.Command(cmd.Config{
		Verb:      "re-init",
		ShortDesc: "Uninstall and reinstall an application.",
	})
	destroy := svc.cmdr.Command(cmd.Config{
		Verb:      "destroy",
		ShortDesc: "Uninstall an application.",
	})

	cmds.AttachCommands(create, deploy, logs, reinit, destroy)

	return cmds
}

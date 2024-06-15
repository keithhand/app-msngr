package cmd

import "github.com/spf13/cobra"

type cobraCommand struct {
	*cobra.Command
}

func newCobraCommand(cfg Config) *cobraCommand {
	return &cobraCommand{
		&cobra.Command{
			Use:   cfg.Verb,
			Short: cfg.ShortDesc,
			Long:  cfg.LongDesc,
		},
	}
}

func (cmd *cobraCommand) Run() error {
	return cmd.Command.Execute()
}

func (cmd *cobraCommand) AttachCommands(cmds ...Commander) {
	for _, c := range cmds {
		cmd.Command.AddCommand(c.(*cobraCommand).Command)
	}
}

package cmd

import "github.com/spf13/cobra"

type cobraCommander struct {
	cmd *cobra.Command
}

func CobraCommander() *cobraCommander {
	return &cobraCommander{}
}

func (c *cobraCommander) AttachCommands(cmds ...Commander) {
	for _, cmd := range cmds {
		c.cmd.AddCommand(cmd.(*cobraCommander).cmd)
	}
}

func (c *cobraCommander) Run() error {
	return c.cmd.Execute()
}

func (c *cobraCommander) Command(cfg Config) Commander {
	return &cobraCommander{
		cmd: &cobra.Command{
			Use:   cfg.Verb,
			Short: cfg.ShortDesc,
			Long:  cfg.LongDesc,
		}}
}

package cmd

import "github.com/spf13/cobra"

type cobraCommand struct {
	*cobra.Command
}

func CobraCommand(cfg Config) *cobraCommand {
	return &cobraCommand{
		&cobra.Command{
			Use:   cfg.Verb,
			Short: cfg.ShortDesc,
			Long:  cfg.LongDesc,
		},
	}
}

func (c *cobraCommand) AddCommand(cc ...Command) {
	for _, cmd := range cc {
		c.Command.AddCommand(cmd.(*cobraCommand).Command)
	}
}

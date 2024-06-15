package cmd

type Config struct {
	Verb      string
	ShortDesc string
	LongDesc  string
}

type Command interface {
	AddCommand(...Command)
	Execute() error
}

func New(cfg Config) Command {
	return CobraCommand(cfg)
}

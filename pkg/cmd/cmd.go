package cmd

type Config struct {
	Verb      string
	ShortDesc string
	LongDesc  string
}

type Commander interface {
	AttachCommands(...Commander)
	Run() error
}

func NewCommand(cfg Config) Commander {
	return newCobraCommand(Config{
		Verb:      cfg.Verb,
		ShortDesc: cfg.ShortDesc,
		LongDesc:  cfg.LongDesc,
	})
}

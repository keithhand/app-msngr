package cmd

type Config struct {
	Verb      string
	ShortDesc string
	LongDesc  string
}

type Commander interface {
	Command(Config) Commander
	AttachCommands(...Commander)
	Run() error
}

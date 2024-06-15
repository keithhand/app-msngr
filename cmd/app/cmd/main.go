package cmd

var (
	slogan  = "Keep the talking with your applications short"
	rootCmd = &command{
		Use:   "app",
		Short: slogan,
		Long: slogan + ".\n" + `
 Reads an application spec for access to common Kubernetes
 tasks without needing to know deployment specifics.`,
	}
)

type Cli interface {
	Execute() error
}

func MainCmd() Cli {
	return rootCmd
}

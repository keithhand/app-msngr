package config

type Config struct {
	Kubernetes Kubernetes
	Helm       Helm
}

type Kubernetes struct {
	Context       string
	Namespace     string
	LogContainers struct {
		Labels []map[string]string
	}
}

type Helm struct {
	Name  string
	Chart string
	Repo  string
}

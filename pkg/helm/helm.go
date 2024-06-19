package helm

import (
	"fmt"

	"khand.dev/msngr/pkg/config"
)

var (
	cfg   = config.New().Helm
	name  = cfg.Name
	chart = cfg.Chart
	repo  = cfg.Repo
)

func Install() error {
	fmtStr := "helm repo add %s %s\n"
	fmt.Printf(fmtStr, "ahh", repo)
	fmt.Println("helm repo update")
	fmtStr = "helm install %s %s/%s\n"
	fmt.Printf(fmtStr, name, chart, repo)
	return nil
}

func Update() error {
	fmtStr := "helm upgrade %s %s/%s\n"
	fmt.Printf(fmtStr, name, chart, repo)
	return nil
}

func Uninstall() error {
	fmtStr := "helm uninstall %s\n"
	fmt.Printf(fmtStr, name)
	return nil
}

package kube

import (
	"fmt"

	"khand.dev/msngr/pkg/config"
)

var cfg = config.New().Kubernetes

func Logs() error {
	fmtStr := "kubectl logs --context %s --namespace %s -l %s: %s\n"
	ctx, ns := cfg.Context, cfg.Namespace
	for _, ll := range cfg.LogContainers.Labels {
		for k, v := range ll {
			fmt.Printf(fmtStr, ctx, ns, k, v)
		}
	}
	return nil
}

package main

import (
	v1 "github.com/chefmtt/helm-controller-riscv64/pkg/apis/helm.cattle.io/v1"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/chefmtt/helm-controller-riscv64/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"helm.cattle.io": {
				Types: []interface{}{
					v1.HelmChart{},
					v1.HelmChartConfig{},
				},
				GenerateTypes: true,
			},
		},
	})
}

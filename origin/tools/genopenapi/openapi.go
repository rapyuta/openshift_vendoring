package main

import (
	"github.com/openshift/k8s.io/gengo/args"
	"github.com/openshift/kubernetes/cmd/libs/go2idl/openapi-gen/generators"

	"github.com/golang/glog"
)

func main() {
	arguments := args.Default()

	// Override defaults.
	arguments.OutputFileBaseName = "zz_generated.openapi"
	arguments.GoHeaderFilePath = "hack/boilerplate.txt"
	arguments.GeneratedBuildTag = "ignore_autogenerated_openshift"

	// Run it.
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.V(2).Info("Completed successfully.")
}

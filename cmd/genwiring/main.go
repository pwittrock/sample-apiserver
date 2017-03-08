package main

import (
	"os"
	"runtime"

	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/gengo/args"
	"k8s.io/sample-apiserver/cmd/genwiring/generators"
	"github.com/golang/glog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	arguments := args.Default()

	// Override defaults.
	arguments.OutputFileBaseName = "zz_generated.api.register"

	// Custom args.
	customArgs := &generators.CustomArgs{}
	arguments.CustomArgs = customArgs

	g := generators.Gen{}
	if err := g.Execute(arguments); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.V(2).Info("Completed successfully.")
}


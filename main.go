package main

import (
	"fmt"
	"plugin"
	"runtime"

	// This line causes the issue.
	_ "google.golang.org/grpc"
)

func main() {
	plugin.Open(fmt.Sprintf("plugin-%s.so", runtime.Version()))
}

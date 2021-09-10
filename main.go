package main

import (
	"os"
	"plugin"

	// This line causes the issue.
	_ "google.golang.org/grpc"
)

func main() {
	plugin.Open(os.Args[1])
}

package main

import (
	"fmt"
	"log"
	"os"
	"plugin"

	_ "google.golang.org/grpc"
)

func main() {
	p, err := plugin.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	sym, err := p.Lookup("Conn")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sym)
}

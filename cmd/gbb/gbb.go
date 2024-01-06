package main

import (
	"fmt"
	"os"

	"github.com/mortedecai/go-burn-bits/gbb"
)

var (
	version string = "<unknown>"
)

func greetings() string {
	return "Hello, World!"
}

func Version() string {
	return version
}

func main() {
	fmt.Printf("%s [%s]\n", greetings(), Version())
	if err := gbb.New("localhost:9990", "").Run(os.Args[1:]); err != nil {
		fmt.Printf("Error during execution: %s\n", err.Error())
	}
}

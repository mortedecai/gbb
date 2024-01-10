package main

import (
	"fmt"
	"os"

	"github.com/mortedecai/gbb/gbb"
)

var (
	version string = "<unknown>"
)

func greetings() string {
	return "Go Burn Bits"
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

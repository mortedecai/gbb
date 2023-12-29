package main

import (
	"fmt"
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
}

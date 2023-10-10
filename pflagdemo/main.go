package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	name := "verbose"
	short := name[:1]

	pflag.BoolP(name, short, false, "verbose output")

	// len(short) must be == 1
	flag := pflag.ShorthandLookup(short)

	fmt.Println(flag.Name)
}

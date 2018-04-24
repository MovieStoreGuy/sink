package main

import (
	"flag"
	"os"

	"github.com/MovieStoreGuy/sink/sink"
	"github.com/fatih/color"
)

func main() {
	flag.Parse()
	if err := sink.EnsureResolveOf(flag.Args()...); err != nil {
		color.Red("We have an error:%v", err)
		os.Exit(-1)
	}
	color.Green("Everything resolved correctly")
}

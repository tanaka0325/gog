package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tanaka0325/gog"
)

var (
	cmd      string
	isDryRun bool
)

func init() {
	flag.StringVar(&cmd, "cmd", "", "command name that you want to execute")
	flag.BoolVar(&isDryRun, "n", false, "dry run")
}

func main() {
	flag.Parse()
	args := flag.Args()

	g := gog.New(cmd, isDryRun, args)
	if err := g.Gen(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

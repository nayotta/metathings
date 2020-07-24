package main

import (
	"github.com/nayotta/metathings/cmd/metathings/cmd"
)

var (
	Version string
)

func main() {
	cmd.SetVersion(Version)
	cmd.RootCmd.Execute()
}

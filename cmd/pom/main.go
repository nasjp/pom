package main

import (
	"fmt"
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/service/command"
	"github.com/YukihiroTaniguchi/pom/interface/environment"
	"github.com/YukihiroTaniguchi/pom/interface/file"
	"github.com/spf13/cobra"
)

const (
	// GOPATH ...
	GOPATH = "GOPATH"
	// APPDIR ...
	APPDIR = "/src/github.com/YukihiroTaniguchi/pom"
	// FILE ...
	FILE = "/config/pom.json"
)

func init() {
	cobra.OnInitialize()
	command.RootCmd.AddCommand(command.StartCmd)

	ev, err := environment.GetEnvVar(GOPATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	if err := file.CreateOrNot(ev.Dir + APPDIR + FILE); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}

func main() {
	if err := command.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}

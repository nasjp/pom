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
)

func init() {
	cobra.OnInitialize()
	command.RootCmd.AddCommand(command.StartCmd)
}

func main() {
	if err := command.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	ev, err := environment.GetEnvVar(GOPATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	if err := file.CreateOrNot(ev.Dir + APPDIR + "/config/.pomrc"); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}

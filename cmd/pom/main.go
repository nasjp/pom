package main

import (
	"fmt"
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
	"github.com/YukihiroTaniguchi/pom/domain/service/command"
	"github.com/YukihiroTaniguchi/pom/infrastructure/environment"
	"github.com/YukihiroTaniguchi/pom/infrastructure/file"
	"github.com/spf13/cobra"
)

const (
	// GOPATH ...
	GOPATH = "GOPATH"
	// APPDIR ...
	APPDIR = "/src/github.com/YukihiroTaniguchi/pom"
	// CONFIGFILE ...
	CONFIGFILE = "/config/pom.json"
)

var (
	defaultConfig = timeset.Setting{
		Work:       25,
		ShortBreak: 10,
		LongBreak:  20,
		Set:        10,
	}
)

func init() {
	cobra.OnInitialize()
	command.RootCmd.AddCommand(command.StartCmd)

	ev, err := environment.GetEnvVar(GOPATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	path := ev.Dir + APPDIR + CONFIGFILE
	if err := file.InitConfigFile(path, &defaultConfig); err != nil {
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

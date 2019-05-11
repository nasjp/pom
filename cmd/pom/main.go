package main

import (
	"fmt"
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/service/command"
	"github.com/YukihiroTaniguchi/pom/interface/file"
	"github.com/spf13/cobra"
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
	if err := file.Cd(ev.Dir + APPDIR); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
	if err := file.CreateOrNot("hoge/huga"); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}

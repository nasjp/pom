package main

import (
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/service/command/loop"
	"github.com/YukihiroTaniguchi/pom/domain/service/command/root"
	"github.com/YukihiroTaniguchi/pom/domain/service/command/set"
	"github.com/YukihiroTaniguchi/pom/domain/service/command/start"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	root.Cmd.AddCommand(start.Cmd)
	root.Cmd.AddCommand(set.Cmd)
	root.Cmd.AddCommand(loop.Cmd)
}

func main() {
	if err := root.Cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

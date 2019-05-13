package main

import (
	"fmt"
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/service/command/root"
	"github.com/YukihiroTaniguchi/pom/domain/service/command/set"
	"github.com/YukihiroTaniguchi/pom/domain/service/command/start"
	"github.com/YukihiroTaniguchi/pom/infrastructure/file"
	"github.com/spf13/cobra"
)

func init() {
	if err := file.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
	cobra.OnInitialize()
	root.Cmd.AddCommand(start.Cmd)
	root.Cmd.AddCommand(set.Cmd)
}

func main() {
	if err := root.Cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}

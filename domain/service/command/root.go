package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:   "pom",
	Short: "pomodoro timer on shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

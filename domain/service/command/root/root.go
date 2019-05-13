package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd is root command
var Cmd = &cobra.Command{
	Use:   "pom",
	Short: "pomodoro timer on shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

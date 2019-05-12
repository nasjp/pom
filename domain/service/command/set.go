package command

import (
	"github.com/spf13/cobra"
)

// SetCmd represents the start command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "set pomodoro timer",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

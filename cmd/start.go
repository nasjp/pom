package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
  Use:   "start",
  Short: "start pomodoro timer",
  Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Args: cobra.RangeArgs(0, 1),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("start called")
  },
}

func init() {
  RootCmd.AddCommand(startCmd)
}

package cmd

import (
  "fmt"
  "time"

  "github.com/spf13/cobra"
  pb "gopkg.in/cheggaaa/pb.v2"
)

// Options represents startCmd's options
type Options struct {
  mins int
}

var (
  o = &Options{}
)

// startCmd represents the start command
var startCmd = &cobra.Command{
  Use:   "start",
  Short: "start pomodoro timer",
  Args:  cobra.NoArgs,
  Run: func(cmd *cobra.Command, args []string) {

    fmt.Printf("start pomodoro in %d minutes!!\n", o.mins)

    // Start parallel processing
    bar := `🍅  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | red}} `
    outputBar(bar)
  },
}

func outputBar(tmpl string) {
  secs := 60 * o.mins
  bar := pb.ProgressBarTemplate(tmpl).Start(secs)
  bar.Set("minutes", o.mins)
  defer bar.Finish()
  for i := 0; i < secs; i++ {
    bar.Add(1)
    time.Sleep(time.Second)
  }
}

func init() {
  RootCmd.AddCommand(startCmd)
  startCmd.Flags().IntVarP(&o.mins, "set", "s", 25, "set the timer")
}

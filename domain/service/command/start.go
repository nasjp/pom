package command

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	pb "gopkg.in/cheggaaa/pb.v2"
)

// Options represents startCmd's options
type Options struct {
	Mins int
}

var (
	o = &Options{}
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start pomodoro timer",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("start pomodoro in %d minutes!!\n", o.Mins)

		// Start parallel processing
		bar := `🍅  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | red}} `
		outputBar(bar)
	},
}

func outputBar(tmpl string) {
	secs := 60 * o.Mins
	bar := pb.ProgressBarTemplate(tmpl).Start(secs)
	bar.Set("minutes", o.Mins)
	defer bar.Finish()
	for i := 0; i < secs; i++ {
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

func init() {
	StartCmd.Flags().IntVarP(&o.Mins, "set", "s", 25, "set the timer")
}

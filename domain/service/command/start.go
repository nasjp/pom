package command

import (
	"fmt"
	"time"

	"github.com/YukihiroTaniguchi/pom/domain/service/option"
	"github.com/spf13/cobra"
	pb "gopkg.in/cheggaaa/pb.v2"
)

var (
	sto = &option.Options{}
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start pomodoro timer",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("start pomodoro in %d minutes!!\n", sto.Mins)

		// Start parallel processing
		bar := `🍅  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | red}} `
		outputBar(bar)
	},
}

func outputBar(tmpl string) {
	secs := 60 * sto.Mins
	bar := pb.ProgressBarTemplate(tmpl).Start(secs)
	bar.Set("minutes", sto.Mins)
	defer bar.Finish()
	for i := 0; i < secs; i++ {
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

func init() {
	StartCmd.Flags().IntVarP(&sto.Mins, "set", "s", 25, "set the timer")
}

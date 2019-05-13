package start

import (
	"fmt"
	"os"
	"time"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
	"github.com/YukihiroTaniguchi/pom/domain/service/option"
	"github.com/YukihiroTaniguchi/pom/infrastructure/file"
	"github.com/spf13/cobra"
	pb "gopkg.in/cheggaaa/pb.v2"
)

var (
	o = &option.Start{}
	s = &timeset.Setting{}
)

// Cmd represents the start command
var Cmd = &cobra.Command{
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
	bar := pb.ProgressBarTemplate(tmpl).Start(int(secs))
	bar.Set("minutes", o.Mins)
	defer bar.Finish()
	for i := 0; i < int(secs); i++ {
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

func init() {
	var err error
	s, err = file.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
	Cmd.Flags().UintVarP(&o.Mins, "set", "s", s.Work, "set the timer")
}

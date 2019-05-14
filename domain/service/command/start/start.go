package start

import (
	"fmt"
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
	"github.com/YukihiroTaniguchi/pom/domain/service/option"
	"github.com/YukihiroTaniguchi/pom/domain/service/progressbar"
	"github.com/YukihiroTaniguchi/pom/infrastructure/file"
	"github.com/spf13/cobra"
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
		progressbar.Work(o.Mins)
	},
}

func init() {
	var err error
	if err = file.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
	s, err = file.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
	Cmd.Flags().UintVarP(&o.Mins, "set", "s", s.Work, "set the timer")
}

package loop

import (
	"fmt"
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
	"github.com/YukihiroTaniguchi/pom/domain/service/progressbar"
	"github.com/YukihiroTaniguchi/pom/infrastructure/file"
	"github.com/spf13/cobra"
)

var (
	s = &timeset.Setting{}
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "loop",
	Short: "loop pomodoro timer",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 1; i <= int(s.Times); i++ {
			fmt.Printf("Start %d / %d loops!!\n", i, s.Times)
			progressbar.Work(s.Work)
			progressbar.ShortBreak(s.ShortBreak)
			if i%4 == 0 {
				progressbar.LongBreak(s.LongBreak)
			}
		}
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
}

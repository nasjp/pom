package command

import (
	"fmt"
	"os"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
	"github.com/YukihiroTaniguchi/pom/infrastructure/file"
	"github.com/spf13/cobra"
)

var (
	d             bool
	s             = &timeset.Setting{}
	defaultConfig = &timeset.Setting{
		Work:       25,
		ShortBreak: 10,
		LongBreak:  20,
		Times:      10,
	}
)

// SetCmd represents the start command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "set pomodoro timer",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if d {
			file.Update(defaultConfig)
			fmt.Printf("set default config!!\n")
			return
		}
		file.Update(s)
		fmt.Printf("config is updated!!\n")
	},
}

func init() {
	var err error
	s, err = file.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	SetCmd.Flags().BoolVarP(&d, "default", "d", false, "set default config")
	SetCmd.Flags().UintVarP(&s.Work, "work", "w", s.Work, "set work minutes")
	SetCmd.Flags().UintVarP(&s.ShortBreak, "short", "s", s.ShortBreak, "set short break minutes")
	SetCmd.Flags().UintVarP(&s.LongBreak, "long", "l", s.LongBreak, "set long break minutes")
	SetCmd.Flags().UintVarP(&s.Times, "times", "t", s.Times, "set pomodoro times")
}

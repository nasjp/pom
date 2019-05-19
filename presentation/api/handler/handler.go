package handler

import (
	"fmt"

	"github.com/YukihiroTaniguchi/pom/application/usecase"
	"github.com/YukihiroTaniguchi/pom/domain/model"
	"github.com/spf13/cobra"
)

var (
	ds = &model.TimerSet{
		Work:       25,
		ShortBreak: 10,
		LongBreak:  20,
		Times:      10,
	}
)

// TimerSetHandler ...
type TimerSetHandler interface {
	ExecRoot() (*cobra.Command, error)
	ExecStart() (*cobra.Command, error)
	ExecSet() (*cobra.Command, error)
	ExecLoop() (*cobra.Command, error)
}

type timerSetHandler struct {
	TimerSetUseCase usecase.TimerSetUsecase
}

// NewTimerSetHandler ...
func NewTimerSetHandler(u usecase.TimerSetUsecase) TimerSetHandler {
	return &timerSetHandler{TimerSetUseCase: u}
}

func (h *timerSetHandler) ExecRoot() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
		Use:   "pom",
		Short: "pomodoro timer on shell",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command")
		},
	}
	return
}

func (h *timerSetHandler) ExecStart() (cmd *cobra.Command, err error) {
	o := &model.StartOption{}
	s, err := h.TimerSetUseCase.Get(model.GetConfigFile())
	if err != nil {
		return
	}

	cmd = &cobra.Command{
		Use:   "start",
		Short: "start pomodoro timer",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			model.Work(o.Mins)
		},
	}
	cmd.Flags().UintVarP(&o.Mins, "set", "s", s.Work, "set the timer")
	return
}

func (h *timerSetHandler) ExecSet() (cmd *cobra.Command, err error) {
	s, err := h.TimerSetUseCase.Get(model.GetConfigFile())
	if err != nil {
		return
	}
	d := false
	cmd = &cobra.Command{
		Use:   "set",
		Short: "set pomodoro timer",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if d {
				h.TimerSetUseCase.Update(ds, model.GetConfigFile())
				fmt.Printf("set default config!!\n")
				return
			}
			h.TimerSetUseCase.Update(s, model.GetConfigFile())
			fmt.Printf("config is updated!!\n")
		},
	}

	cmd.Flags().BoolVarP(&d, "default", "d", false, "set default config")
	cmd.Flags().UintVarP(&s.Work, "work", "w", s.Work, "set work minutes")
	cmd.Flags().UintVarP(&s.ShortBreak, "short", "s", s.ShortBreak, "set short break minutes")
	cmd.Flags().UintVarP(&s.LongBreak, "long", "l", s.LongBreak, "set long break minutes")
	cmd.Flags().UintVarP(&s.Times, "times", "t", s.Times, "set pomodoro times")
	return
}

func (h *timerSetHandler) ExecLoop() (cmd *cobra.Command, err error) {
	s, err := h.TimerSetUseCase.Get(model.GetConfigFile())
	if err != nil {
		return
	}
	cmd = &cobra.Command{
		Use:   "loop",
		Short: "loop pomodoro timer",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			for i := 1; i <= int(s.Times); i++ {
				fmt.Printf("Start %d / %d loops!!\n", i, s.Times)
				model.Work(s.Work)
				model.ShortBreak(s.ShortBreak)
				if i%4 == 0 {
					model.LongBreak(s.LongBreak)
				}
			}
		},
	}
	return
}

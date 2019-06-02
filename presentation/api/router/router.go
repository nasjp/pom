package router

import (
	"github.com/YukihiroTaniguchi/pom/presentation/api/handler"
	"github.com/spf13/cobra"
)

// NewRouter ...
func NewRouter(h handler.TimerSetHandler) (root *cobra.Command, err error) {
	cobra.OnInitialize()
	root, err = h.ExecRoot()
	if err != nil {
		return
	}

	start, err := h.ExecStart()
	if err != nil {
		return
	}
	root.AddCommand(start)

	set, err := h.ExecSet()
	if err != nil {
		return
	}
	root.AddCommand(set)

	loop, err := h.ExecLoop()
	if err != nil {
		return
	}
	root.AddCommand(loop)

	mob, err := h.ExecMob()
	if err != nil {
		return
	}
	root.AddCommand(mob)
	sound, err := h.ExecSetSound()
	if err != nil {
		return
	}
	root.AddCommand(sound)

	return
}

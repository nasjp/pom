package progressbar

import (
	"time"

	pb "gopkg.in/cheggaaa/pb.v2"
)

// Bar ...
type Bar string

const (
	// WORK ...
	WORK Bar = `🍅  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | red}} `
	// SHORTBREAK ...
	SHORTBREAK Bar = `☕  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | green}} `
)

// Run ...
func Run(m uint) {
	WORK.outputBar(m)
}

// ShortBreak ...
func ShortBreak(m uint) {
	SHORTBREAK.outputBar(m)
}

func (b Bar) outputBar(m uint) {
	secs := 60 * m
	bar := pb.ProgressBarTemplate(b).Start(int(secs))
	bar.Set("minutes", m)
	defer bar.Finish()
	for i := 0; i < int(secs); i++ {
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

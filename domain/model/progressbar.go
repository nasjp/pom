package model

import (
	"fmt"
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
	// LONGBREAK ...
	LONGBREAK Bar = `💤  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | magenta}} `
)

// Work ...
func Work(m uint) {
	if 0 == m {
		return
	}
	fmt.Printf("try to stay focus in %d minutes!!\n", m)
	WORK.outputBar(m)
}

// ShortBreak ...
func ShortBreak(m uint) {
	if 0 == m {
		return
	}
	fmt.Printf("Please take a break in %d minutes.\n", m)
	SHORTBREAK.outputBar(m)
}

// LongBreak ...
func LongBreak(m uint) {
	if 0 == m {
		return
	}
	fmt.Printf("You can take a longer break in %d minutes\n", m)
	LONGBREAK.outputBar(m)
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
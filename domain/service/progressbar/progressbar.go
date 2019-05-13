package progressbar

import (
	"fmt"
	"time"

	pb "gopkg.in/cheggaaa/pb.v2"
)

// Run ...
func Run(m uint) {
	fmt.Printf("start pomodoro in %d minutes!!\n", m)
	bar := `🍅  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | red}} `
	// Start parallel processing
	outputBar(bar, m)
}

func outputBar(tmpl string, m uint) {
	secs := 60 * m
	bar := pb.ProgressBarTemplate(tmpl).Start(int(secs))
	bar.Set("minutes", m)
	defer bar.Finish()
	for i := 0; i < int(secs); i++ {
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

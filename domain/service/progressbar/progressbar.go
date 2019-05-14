package progressbar

import (
	"fmt"
	"time"

	pb "gopkg.in/cheggaaa/pb.v2"
)

const (
	// BAR ...
	BAR = `🍅  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | red}} `
)

// Run ...
func Run(m uint) {
	fmt.Printf("try to stay focus in %d minutes!!\n", m)
	// Start parallel processing
	outputBar(BAR, m)
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

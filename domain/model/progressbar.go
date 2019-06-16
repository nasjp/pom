package model

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"gopkg.in/cheggaaa/pb.v2"
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

	// MOBWORK ...
	MOBWORK Bar = `⏰  : {{etime .}} / {{string . "minutes"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | yellow}} `
	// MOBINTERVAL ...
	MOBINTERVAL Bar = `🔄  : {{etime .}} / {{string . "seconds"}}m ( {{percent . }} ) {{bar . "|" ">" ">" "-" "|" | cyan}} `
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

func MobWorkBar(m uint) {
	if 0 == m {
		return
	}
	fmt.Printf("Start mob programming in %d minutes!!\n", m)
	MOBWORK.outputBar(m)
}

func MobIntervalBar(s uint) {
	if 0 == s {
		return
	}
	fmt.Printf("Change from person to person in %d seconds\n", s)
	MOBINTERVAL.outputBarSecs(s)
}

func (b Bar) outputBar(m uint) {
	secs := 60 * m
	bar := pb.ProgressBarTemplate(b).Start(int(secs))
	bar.Set("minutes", m)
	defer bar.Finish()
	for i := 0; i < int(secs); i++ {
		//if int(secs)-i == 5 {
		//	if err := sound(); err != nil {
		//		log.Fatal(err)
		//	}
		//}
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

func (b Bar) outputBarSecs(secs uint) {
	bar := pb.ProgressBarTemplate(b).Start(int(secs))
	bar.Set("seconds", secs)
	defer bar.Finish()
	for i := 0; i < int(secs); i++ {
		//if int(secs)-i == 5 {
		//	if err := sound(); err != nil {
		//		log.Fatal(err)
		//	}
		//}
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

func sound() error {
	f, err := os.Open("count.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

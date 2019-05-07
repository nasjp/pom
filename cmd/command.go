package cmd

import (
  "flag"
  "fmt"
  "os"
  "time"
)

// Run ...
func Run() {
  defer os.Exit(0)
  flag.Parse()
  args := flag.Args()
  Router(args)
}

// Router ..
func Router(args []string) {
  if len(args) == 0 {
    fmt.Fprintf(os.Stderr, "show help\n")
    os.Exit(0)
  }
  switch args[0] {
  case "start":
    Start()
  case "status":
    Status()
  case "stop":
    Stop()
  default:
    fmt.Fprintf(os.Stderr, "\"pom %s\" is not found!!\n", args[0])
  }
}

// Start ..
func Start() {
  fmt.Fprintf(os.Stdout, "Let's start!!\n")
}

// Status ..
func Status() {
  fmt.Fprintf(os.Stdout, "pom's status is ...\n")
}

// Stop ..
func Stop() {
  fmt.Fprintf(os.Stdout, "pom stop timer\n")
}

func timers() {
  // 変数timerに指定の時間を入れる（今回は2分）
  timer := time.NewTimer(time.Minute * 25)
  // チャネルに通知を受ける
  <-timer.C
  // 通知があればPrintする（ここにイベント）
  println("お時間です！")
}

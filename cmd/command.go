package cmd

import (
  "flag"
  "fmt"
)

// Run ...
func Run() {
  flag.Parse()
  args := flag.Args()
  fmt.Println(args)
}

package environment

import (
	"fmt"
	"os"
)

// ENV ...
type ENV struct {
	Name string
	Dir  string
}

// GetEnvVar ...
func GetEnvVar(n string) (ev ENV, err error) {
	err = nil
	d := os.Getenv(n)
	if d == "" {
		err = fmt.Errorf("\"%s\" is not defined", n)
	}
	ev = ENV{
		Name: n,
		Dir:  d,
	}
	return ev, err
}

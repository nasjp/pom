package environment

import (
	"fmt"
	"os"
)

const (
	// GOPATH ...
	GOPATH = "GOPATH"
)

// ENV ...
type ENV struct {
	Name string
	Dir  string
}

// GetEnvVar ...
func GetEnvVar() (ev ENV, err error) {
	err = nil
	d := os.Getenv(GOPATH)
	if d == "" {
		err = fmt.Errorf("\"%s\" is not defined", GOPATH)
	}
	ev = ENV{
		Name: GOPATH,
		Dir:  d,
	}
	return ev, err
}

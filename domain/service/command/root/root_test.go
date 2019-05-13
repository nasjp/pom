package root

import (
	"testing"
)

// RootCmd is root command
func testCmd(t *testing.T) {
	if err := Cmd.Execute(); err != nil {
		t.Error("can't use 'pom'")
	}
}

package command

import (
	"testing"
)

// RootCmd is root command
func testRootCmd(t *testing.T) {
	if err := RootCmd.Execute(); err != nil {
		t.Error("can't use 'pom'")
	}
}

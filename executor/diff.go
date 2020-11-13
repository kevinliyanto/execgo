package main

import (
	// "os/exec"
	"strings"
)

// Diff between two array of string
// Probably put diff algorithm here
func Diff(left string, right string) (bool, string) {
	return strings.Compare(left, right) == 0, ""
}

package ysok

import (
	"fmt"
	"io"
	"os"
)

var (
	// OutWriter is Writer for standard output.
	OutWriter io.Writer = os.Stdout
	// ErrWriter is Writer for starndard error.
	ErrWriter io.Writer = os.Stderr
)

func errf(format string, a ...interface{}) {
	fmt.Fprintln(ErrWriter, fmt.Sprintf(format, a...))
}

func outf(format string, a ...interface{}) {
	fmt.Fprintln(OutWriter, fmt.Sprintf(format, a...))
}

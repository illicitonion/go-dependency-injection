package cmdlib

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type IO interface {
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
}

func NewIO() IO {
	return &realIO{}
}

func NewTestIO(input string) *TestIO {
	stdin := strings.NewReader(input)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	return &TestIO{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
	}
}

type realIO struct{}

func (i *realIO) Stdin() io.Reader {
	return os.Stdin
}

func (i *realIO) Stdout() io.Writer {
	return os.Stdout
}

func (i *realIO) Stderr() io.Writer {
	return os.Stderr
}

type TestIO struct {
	stdin  io.Reader
	stdout *bytes.Buffer
	stderr *bytes.Buffer
}

func (i *TestIO) Stdin() io.Reader {
	return i.stdin
}

func (i *TestIO) Stdout() io.Writer {
	return i.stdout
}

func (i *TestIO) Stderr() io.Writer {
	return i.stderr
}

func (i *TestIO) StdoutString() string {
	return i.stdout.String()
}

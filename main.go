package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type CliToolIO struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

type Flags struct {
	Uppercase bool
}

type CliToolConfig struct {
	Flags Flags
	Args  []string
}

type CliTool struct {
	io     *CliToolIO
	config *CliToolConfig
}

func NewCliTool(io *CliToolIO, config *CliToolConfig) *CliTool {
	return &CliTool{
		io:     io,
		config: config,
	}
}

func (c *CliTool) Run() {
	fmt.Fprintf(c.io.stdout, "Enter your name: ")

	reader := bufio.NewReader(c.io.stdin)

	// reads from the standard input (c.stdin) using the bufio.Reader until it encounters a newline character ('\n').
	// it blocks execution and waits for the user to enter some text and press the Enter key.
	inputText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(c.io.stderr, "(stderr) failed to read the input text")
	}

	containsNumbers := false

	for _, char := range inputText {
		if char >= '0' && char <= '9' {
			containsNumbers = true
			break
		}
	}

	if c.config.Flags.Uppercase {
		inputText = strings.ToUpper(inputText)
	}

	if len(c.config.Args) > 0 {
		inputText = strings.Join(c.config.Args, " ") + " " + inputText
	}

	if containsNumbers {
		fmt.Fprintf(c.io.stderr, "(stderr): BAD (contains numbers): %s", inputText)
	} else {
		fmt.Fprintf(c.io.stdout, "(stdout): GOOD: %s", inputText)
	}
}

func main() {
	config := &CliToolConfig{}

	flag.BoolVar(&config.Flags.Uppercase, "u", false, "convert input text to uppercase")
	flag.Parse()

	config.Args = flag.Args()

	io := &CliToolIO{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}

	cliTool := NewCliTool(io, config)

	cliTool.Run()
}

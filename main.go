package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/bazmurphy/go-dependency-injection/cmdlib"
	"github.com/jessevdk/go-flags"
)

type Flags struct {
	Uppercase bool `short:"u" long:"uppercase" description:"convert input text to uppercase"`
}

type CliToolConfig struct {
	Flags Flags
	Args  []string
}

type CliTool struct {
	io     cmdlib.IO
	config *CliToolConfig
}

func NewCliTool(io cmdlib.IO, config *CliToolConfig) *CliTool {
	return &CliTool{
		io:     io,
		config: config,
	}
}

func (c *CliTool) Run() {
	fmt.Fprintf(c.io.Stdout(), "Enter your name: ")

	reader := bufio.NewReader(c.io.Stdin())

	// reads from the standard input (c.stdin) using the bufio.Reader until it encounters a newline character ('\n').
	// it blocks execution and waits for the user to enter some text and press the Enter key.
	inputText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(c.io.Stderr(), "(stderr) failed to read the input text")
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
		fmt.Fprintf(c.io.Stderr(), "(stderr): BAD (contains numbers): %s", inputText)
	} else {
		fmt.Fprintf(c.io.Stdout(), "(stdout): GOOD: %s", inputText)
	}
}

func main() {
	config := &CliToolConfig{}

	var err error
	config.Args, err = flags.Parse(&config.Flags)
	if err != nil {
		log.Fatalf("Failed to parse flags: %v", err)
	}

	io := cmdlib.NewIO()

	cliTool := NewCliTool(io, config)

	cliTool.Run()
}

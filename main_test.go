package main

import (
	"bytes"
	"testing"

	"github.com/bazmurphy/go-dependency-injection/cmdlib"
)

func TestCliTool(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		args           []string
		uppercase      bool
		expectedStdout string
		expectedStderr string
	}{
		{
			name:           "Valid input",
			input:          "Baz\n",
			args:           []string{},
			uppercase:      false,
			expectedStdout: "Enter your name: (stdout): GOOD: Baz\n",
			expectedStderr: "",
		},
		{
			name:           "valid input with uppercase flag",
			input:          "Baz\n",
			args:           []string{},
			uppercase:      true,
			expectedStdout: "Enter your name: (stdout): GOOD: BAZ\n",
			expectedStderr: "",
		},
		{
			name:           "valid input with additional argument",
			input:          "Baz\n",
			args:           []string{"Hello"},
			uppercase:      false,
			expectedStdout: "Enter your name: (stdout): GOOD: Hello Baz\n",
			expectedStderr: "",
		},
		{
			name:           "valid input with uppercase flag and additional argument",
			input:          "Baz\n",
			args:           []string{"Hello"},
			uppercase:      true,
			expectedStdout: "Enter your name: (stdout): GOOD: Hello BAZ\n",
			expectedStderr: "",
		},
		{
			name:           "invalid input with numbers",
			input:          "Baz 123\n",
			args:           []string{},
			uppercase:      false,
			expectedStdout: "Enter your name: ",
			expectedStderr: "(stderr): BAD (contains numbers): Baz 123\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			io := cmdlib.NewTestIO(tc.input)

			config := &CliToolConfig{
				Flags: Flags{
					Uppercase: tc.uppercase,
				},
				Args: tc.args,
			}

			cliTool := NewCliTool(io, config)

			cliTool.Run()

			if io.StdoutString() != tc.expectedStdout {
				t.Errorf("Unexpected stdout\nExpected: %q\nGot: %q", tc.expectedStdout, io.StdoutString())
			}

			buf := io.Stderr.(*bytes.Buffer)
			if buf.String() != tc.expectedStderr {
				t.Errorf("Unexpected stderr\nExpected: %q\nGot: %q", tc.expectedStderr, io.Stderr.(*bytes.Buffer).String())
			}
		})
	}
}

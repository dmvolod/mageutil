package shellcmd

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestCommandOutput(t *testing.T) {
	tests := map[string]struct {
		command  string
		expected string
	}{
		"success command": {
			command:  "go version",
			expected: "",
		},
		"unexpected command": {
			command:  "go1",
			expected: "executable file not found in $PATH",
		},
		"incorrect command": {
			command:  "go dd",
			expected: "go dd: unknown command",
		},
	}

	for cmd, tt := range tests {
		t.Run(cmd, func(t *testing.T) {
			is := is.New(t)
			_, err := Command(tt.command).Output()
			if tt.expected == "" {
				is.NoErr(err)
			} else {
				println(err.Error())
				is.True(strings.Contains(err.Error(), tt.expected))
			}
		})
	}
}

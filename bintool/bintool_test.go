package bintool

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
)

func TestGoBin(t *testing.T) {
	isg := is.New(t)
	wd, err := os.Getwd()
	isg.NoErr(err)

	tests := map[string]struct {
		env      string
		value    string
		expected string
	}{
		"with GOBIN specified": {
			env:      "GOBIN",
			value:    filepath.Join(wd, "gobin"),
			expected: filepath.Join(wd, "gobin"),
		},
		"with GOPATH specified": {
			env:      "GOPATH",
			value:    wd,
			expected: filepath.Join(wd, "bin"),
		},
	}

	for cmd, tt := range tests {
		t.Run(cmd, func(t *testing.T) {
			defer os.Unsetenv(tt.env)

			is := is.New(t)
			is.NoErr(os.Setenv(tt.env, tt.value))
			gobin, err := GoBin()
			is.NoErr(err)
			is.Equal(gobin, tt.expected)
		})
	}

}

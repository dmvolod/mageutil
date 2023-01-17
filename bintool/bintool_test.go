package bintool

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
)

func TestGoBin(t *testing.T) {
	is := is.New(t)
	wd, err := os.Getwd()
	is.NoErr(err)

	tests := map[string]struct {
		env      string
		path     string
		expected string
	}{
		"with GOBIN specified": {
			env:      "GOBIN",
			path:     filepath.Join(wd, "gobin"),
			expected: filepath.Join(wd, "gobin"),
		},
		"with GOPATH specified": {
			env:      "GOPATH",
			path:     wd,
			expected: filepath.Join(wd, "bin"),
		},
	}

	for cmd, tt := range tests {
		t.Run(cmd, func(t *testing.T) {
			defer os.Unsetenv(tt.env)
			is.NoErr(os.Setenv(tt.env, tt.path))
			gobin, err := GoBin()
			is.NoErr(err)
			is.Equal(gobin, tt.expected)
		})
	}

}

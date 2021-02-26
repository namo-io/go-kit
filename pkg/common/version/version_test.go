package version

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	Application = "app"
	Timestamp = "now"
	Commit = "sha"
	Tag = "tag"

	t.Run("parse returns correct Version", func(t *testing.T) {
		v := parse()

		expected := Version{
			Application: "app",
			Timestamp:   "now",
			Commit:      "sha",
			Tag:         "tag",
			Golang:      runtime.Version(),
			OS:          runtime.GOOS,
			Arch:        runtime.GOARCH,
			Modules:     map[string]string{},
		}
		fmt.Println(v)

		require.Equal(t, expected, v)
	})
}

package input

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestFrom(t *testing.T) {
	t.Run("return all lines from reader, ignoring empty lines", func(t *testing.T) {
		lines := From(strings.NewReader(`first

second
third`))

		require.Equal(t, []string{
			"first",
			"second",
			"third",
		}, lines)
	})
}

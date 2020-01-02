package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WarpString(t *testing.T) {
	t.Run("it warps a string", func(t *testing.T) {
		res := WarpString("French Fries3", 16)
		require.Equal(t, res, []string{"French Fries3 "})
	})

	t.Run("it warps a string over the length", func(t *testing.T) {
		res := WarpString("French Fries3", 4)
		require.Equal(t, res, []string{"Fr-", "en-", "ch", "Fr-", "ie-", "s3 "})
	})
}

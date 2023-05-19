package search_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch_FindOneMatchTest(t *testing.T) {
	t.Run("success match found", func(t *testing.T) {
		searcher := Searcher{}

		res := searcher.search("item")

		require.Equal(t, "Item 1", res)
	})
}

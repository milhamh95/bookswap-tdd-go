package search_test

import (
	"bookswap-tdd-go/search"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch_FindOneMatchTest(t *testing.T) {
	t.Run("success match found", func(t *testing.T) {
		searcher := search.Searcher{}

		searcher.Search("item")

		require.Equal(t, "Item 1", searcher.GetResult())
	})

	t.Run("another match found", func(t *testing.T) {
		t.Skip("refactoring")
		searcher := search.Searcher{}

		searcher.Search("another")

		require.Equal(t, "Another Item", searcher.GetResult())
	})
}

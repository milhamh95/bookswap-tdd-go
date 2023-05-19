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

		require.Equal(t, search.SearchState{Match: "Item 1"}, searcher.ResultState())
	})

	t.Run("another match found", func(t *testing.T) {
		searcher := search.Searcher{}

		searcher.Search("another")

		require.Equal(t, search.SearchState{Match: "Another Item"}, searcher.ResultState())
	})

	t.Run("no match found", func(t *testing.T) {
		searcher := search.Searcher{}

		searcher.Search("coffee")

		require.Equal(t, "No match found for coffee", searcher.GetResult())
		require.Equal(t, search.SearchState{NoMatchFor: "coffee"}, searcher.ResultState())
	})

	t.Run("empty query", func(t *testing.T) {
		searcher := search.Searcher{}

		searcher.Search("")

		require.Equal(t, "Error: bad query", searcher.GetResult())
	})

	t.Run("short query", func(t *testing.T) {
		searcher := search.Searcher{}

		searcher.Search("abc")

		require.Equal(t, "Error: bad query", searcher.GetResult())
	})
}

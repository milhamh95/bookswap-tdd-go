package search_test

import (
	"bookswap-tdd-go/search"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch_FindOneMatchTest(t *testing.T) {
	searcher := search.Searcher{
		Validator: search.QueryValidator{},
	}

	t.Run("success match found", func(t *testing.T) {
		searcher.Search("item")

		require.Equal(t, search.SearchState{Result: "Item 1"}, searcher.ResultState())
	})

	t.Run("another match found", func(t *testing.T) {
		searcher.Search("another")

		require.Equal(t, search.SearchState{Result: "Another Item"}, searcher.ResultState())
	})

	t.Run("no match found", func(t *testing.T) {
		searcher.Search("coffee")

		require.Equal(t, search.SearchState{Result: "No match found for coffee"}, searcher.ResultState())
	})

	t.Run("empty query", func(t *testing.T) {
		searcher.Search("")

		require.Equal(t, "Error: bad query", searcher.GetResult())
		require.Equal(t, search.SearchState{Result: "Error: bad query"}, searcher.ResultState())
	})

	t.Run("short query", func(t *testing.T) {

		searcher.Search("abc")

		require.Equal(t, "Error: bad query", searcher.GetResult())
		require.Equal(t, search.SearchState{Result: "Error: bad query"}, searcher.ResultState())
	})

	t.Run("another short query", func(t *testing.T) {
		searcher.Search("     ab")

		require.Equal(t, search.SearchState{Result: "Error: bad query"}, searcher.ResultState())
	})
}

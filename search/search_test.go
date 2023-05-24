package search_test

import (
	"bookswap-tdd-go/search"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch_FindOneMatchTest(t *testing.T) {
	minQueryLength := 3
	availableQueries := []string{"Item 1", "Another Value"}
	searcher := search.Searcher{
		Validator: search.QueryValidator{
			MinQueryLength: minQueryLength,
		},
		Repository: search.NewRepository(availableQueries),
	}

	t.Run("success match found", func(t *testing.T) {
		searcher.Search("item")

		expectedResult := search.SearchState{
			Result: []string{"Item 1"},
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})

	t.Run("another match found", func(t *testing.T) {
		searcher.Search("another")

		expectedResult := search.SearchState{
			Result: []string{"Another Value"},
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})

	t.Run("no match found", func(t *testing.T) {
		searcher.Search("coffee")

		expectedResult := search.SearchState{
			Result: []string{},
			Error:  errors.New("no match found for coffee"),
		}
		require.Equal(t, expectedResult, searcher.ResultState())
	})

	t.Run("empty query", func(t *testing.T) {
		searcher.Search("")

		expectedResult := search.SearchState{
			Result: []string{},
			Error:  errors.New("error: bad query"),
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})

	t.Run("short query", func(t *testing.T) {
		searcher.Search("abc")

		expectedResult := search.SearchState{
			Result: []string{},
			Error:  errors.New("error: bad query"),
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})

	t.Run("another short query", func(t *testing.T) {
		searcher.Search("     ab")

		expectedResult := search.SearchState{
			Result: []string{},
			Error:  errors.New("error: bad query"),
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})
}

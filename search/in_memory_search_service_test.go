package search_test

import (
	"bookswap-tdd-go/search"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InMemorySearchService(t *testing.T) {
	t.Run("return empty result", func(t *testing.T) {
		availableQueries := []string{"Item1", "item 2"}
		searchService := search.NewInMemorySearchService(availableQueries)
		result, _ := searchService.FindMatches("block")

		require.Equal(t, []string{}, result)
	})

	t.Run("return matches", func(t *testing.T) {
		availableQueries := []string{"one", "item 1", "two", "Item 2", "else", "ITEM 3"}
		searchService := search.NewInMemorySearchService(availableQueries)
		result, _ := searchService.FindMatches("item")

		expectedResult := []string{"item 1", "Item 2", "ITEM 3"}
		require.Equal(t, expectedResult, result)
	})

	t.Run("return bad search", func(t *testing.T) {
		searchService := search.NewInMemorySearchService([]string{"item 1", "item 2"})

		_, err := searchService.FindMatches("")
		expectedError := errors.New("bad search")
		require.EqualError(t, expectedError, err.Error())
	})
}

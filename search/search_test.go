package search_test

import (
	"bookswap-tdd-go/search"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch_FindOneMatch(t *testing.T) {
	minQueryLength := 3
	availableQueries := []string{"Item 1", "Another Value"}
	inMemorySearchService := search.NewInMemorySearchService(availableQueries)
	searcher := search.Searcher{
		Validator: search.QueryValidator{
			MinQueryLength: minQueryLength,
		},
		Repository: search.NewRepository(inMemorySearchService),
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

}

func TestSearch_FindMultipleMatches(t *testing.T) {
	t.Run("multiple matches found", func(t *testing.T) {
		minQueryLength := 2
		availableQueries := []string{"item one", "else", "one", "one value", "other"}
		inMemorySearchService := search.NewInMemorySearchService(availableQueries)
		searcher := search.Searcher{
			Validator: search.QueryValidator{
				MinQueryLength: minQueryLength,
			},
			Repository: search.NewRepository(inMemorySearchService),
		}

		searcher.Search("one")
		expectedResult := search.SearchState{
			Result: []string{"item one", "one", "one value"},
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})
}

func TestSearch_FindNoMatches(t *testing.T) {
	minQueryLength := 3
	availableQueries := []string{"Item 1", "Another Value"}
	inMemorySearchService := search.NewInMemorySearchService(availableQueries)
	searcher := search.Searcher{
		Validator: search.QueryValidator{
			MinQueryLength: minQueryLength,
		},
		Repository: search.NewRepository(inMemorySearchService),
	}

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

type OfflineSearchService struct {
}

func (u OfflineSearchService) FindMatches(query string) ([]string, error) {
	return []string{}, errors.New("offline")
}

type UnavailableSearchService struct {
}

func (u UnavailableSearchService) FindMatches(query string) ([]string, error) {
	return []string{}, errors.New("bad search")
}

func TestSearch_ErrorSearch(t *testing.T) {
	minQueryLength := 3
	t.Run("bad search", func(t *testing.T) {

		unavailableSearchService := UnavailableSearchService{}
		searcher := search.Searcher{
			Validator: search.QueryValidator{
				MinQueryLength: minQueryLength,
			},
			Repository: search.NewRepository(unavailableSearchService),
		}

		searcher.Search("irrelevant")

		expectedResult := search.SearchState{
			Result: []string{},
			Error:  errors.New("bad search"),
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})

	t.Run("offline", func(t *testing.T) {
		offlineSearchService := OfflineSearchService{}
		searcher := search.Searcher{
			Validator: search.QueryValidator{
				MinQueryLength: minQueryLength,
			},
			Repository: search.NewRepository(offlineSearchService),
		}

		searcher.Search("irrelevant")

		expectedResult := search.SearchState{
			Result: []string{},
			Error:  errors.New("offline"),
		}

		require.Equal(t, expectedResult, searcher.ResultState())
	})
}

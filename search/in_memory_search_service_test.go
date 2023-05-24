package search_test

import (
	"bookswap-tdd-go/search"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InMemorySearchService(t *testing.T) {
	availableQueries := []string{"Item1", "item 2"}
	searchService := search.NewInMemorySearchService(availableQueries)
	result, _ := searchService.FindMatches("block")

	require.Equal(t, []string{}, result)
}

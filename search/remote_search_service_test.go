package search_test

import (
	"bookswap-tdd-go/search"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

type OfflineSearchApi struct{}

func (o OfflineSearchApi) FindMatches(query string) ([]string, error) {
	return []string{}, errors.New("search api is offline")
}

func Test_RemoteSearchService(t *testing.T) {
	t.Run("offline", func(t *testing.T) {
		offlineSearchApi := OfflineSearchApi{}
		remoteSearchService := search.NewRemoteSearchService(offlineSearchApi)

		_, err := remoteSearchService.FindMatches("block")
		expectedError := errors.New("search api is offline")

		require.EqualError(t, expectedError, err.Error())

	})

	//t.Run("return empty result", func(t *testing.T) {
	//	offlineSearchApi := search.
	//	remoteSearchService := search.NewRemoteSearchService(availableQueries)
	//	result, _ := remoteSearchService.FindMatches("block")
	//
	//	require.Equal(t, []string{}, result)
	//})
	//
	//t.Run("return matches", func(t *testing.T) {
	//	availableQueries := []string{"one", "item 1", "two", "Item 2", "else", "ITEM 3"}
	//	searchService := search.NewRemoteSearchService(availableQueries)
	//	result, _ := searchService.FindMatches("item")
	//
	//	expectedResult := []string{"item 1", "Item 2", "ITEM 3"}
	//	require.Equal(t, expectedResult, result)
	//})
	//
	//t.Run("return bad search", func(t *testing.T) {
	//	searchService := search.NewRemoteSearchService([]string{"item 1", "item 2"})
	//
	//	_, err := searchService.FindMatches("")
	//	expectedError := errors.New("bad search")
	//	require.EqualError(t, expectedError, err.Error())
	//})
}

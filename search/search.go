package search

import "fmt"

type Searcher struct {
	Validator QueryValidator
}

var searchStateLiveData = ""

func (s Searcher) Search(query string) {
	if !s.Validator.Validate(query) {
		searchStateLiveData = "Error: bad query"
		return
	}

	availableQueries := map[string]string{
		"item":    "Item 1",
		"another": "Another Item",
	}

	item, ok := availableQueries[query]
	if !ok {
		searchStateLiveData = fmt.Sprintf("No match found for %s", query)
		return
	}

	searchStateLiveData = item
}

func (s Searcher) ResultState() SearchState {
	return SearchState{
		Result: searchStateLiveData,
	}
}

package search

import (
	"fmt"
)

type Searcher struct {
	validator QueryValidator
}

var result = "Item 1"
var searchStateLiveData = ""

func (s Searcher) Search(query string) {
	if s.validator.Validate(query) {
		result = "Error: bad query"
		searchStateLiveData = "Error: bad query"
		return
	}

	if query == "item" {
		result = "Item 1"
		searchStateLiveData = "Item 1"
	} else if query == "another" {
		result = "Another Item"
		searchStateLiveData = "Another Item"
	} else {
		result = fmt.Sprintf("No match found for %s", query)
		searchStateLiveData = fmt.Sprintf("No match found for %s", query)
	}

}

func (s Searcher) GetResult() string {
	return result
}

func (s Searcher) ResultState() SearchState {
	return SearchState{
		Result: searchStateLiveData,
	}
}

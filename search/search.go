package search

import (
	"fmt"
	"strings"
)

type Searcher struct {
}

var result = "Item 1"
var searchStateLiveData = ""

func (s Searcher) Search(query string) {
	isValid := len(strings.TrimSpace(query)) > 3
	if !isValid {
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

package search

import "fmt"

type Searcher struct {
}

var result = "Item 1"

func (s Searcher) Search(query string) {
	if query == "" || query == "abc" {
		result = "Error: bad query"
		return
	}

	if query == "item" {
		result = "Item 1"
	} else if query == "another" {
		result = "Another Item"
	} else {
		result = fmt.Sprintf("No match found for %s", query)
	}

}

func (s Searcher) GetResult() string {
	return result
}
